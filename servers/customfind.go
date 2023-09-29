package servers

import (
	"log"
	"context"
	"rb3server/database"
	"rb3server/models"
	"time"

	"github.com/knvtva/nex-go"
	nexproto "github.com/knvtva/nex-protocols-go"
	"go.mongodb.org/mongo-driver/bson"
)

func CustomFind(err error, client *nex.Client, callID uint32, data []byte) {

	if client.PlayerID() == 0 {
		log.Println("Client is attempting to check for gatherings without a valid server-assigned PID, rejecting call")
		SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
		return
	}

	if client.Username == "Master User" {
		log.Printf("Ignoring CheckForGatherings for unauthenticated Wii master user with friend code %s\n", client.WiiFC)
		SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
		return
	}
	log.Printf("Checking for available gatherings for %s...\n", client.Username)

	gatheringCollection := database.RockcentralDatabase.Collection("gatherings")
	usersCollection := database.RockcentralDatabase.Collection("users")

	cur, err := gatheringCollection.Aggregate(nil, []bson.M{
		bson.M{"$match": bson.D{
			{
				Key:   "creator",
				Value: bson.D{{Key: "$eq", Value: client.Username}},
			},
		}},
		bson.M{"$project": bson.D{
			{
				Key: "hosting",
				Value: bson.D{
					{
						Key: "$eq",
						Value: bson.A{
							bson.M{"$ifNull": []interface{}{"$host", 0}},
							1,
						},
					},
				},
			},
			{
				Key: "room_code",
				Value: "$room_code",
			},
		}},
	})


	
	if err != nil {
		log.Printf("Error: %+v\n", err)
		SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
		return
	}


	var gatheringResult bson.M
	for cur.Next(context.TODO()) {
		if err := cur.Decode(&gatheringResult); err != nil {
			log.Printf("Error decoding gathering: %+v\n", err)
			SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
			return
		}

		hosting := gatheringResult["hosting"].(bool)
		roomCode := gatheringResult["room_code"].(string)

		if hosting {
			rmcResponseStream := nex.NewStream()
			rmcResponseStream.Grow(50)
			log.Println("%v is hosting a private lobby.", client.Username)
	
			// Update the gathering to set it as public and in state 0
			_, err := gatheringCollection.UpdateOne(
				context.TODO(),
				bson.M{"creator": client.Username},
				bson.M{
					"$set": bson.M{
						"public": 1,
						"state": 0,
					},
				},
			)
			if err != nil {
				log.Printf("Error updating gathering: %v", err)
			}
	
			// Find gatherings with host 0 and a matching room code
			filter := bson.M{
				"host":      0,
				"room_code": roomCode,
			}
			ctx := context.TODO()
			cur, err := gatheringCollection.Find(ctx, filter)
			if err != nil {
				log.Printf("Error finding gatherings: %v", err)
			}
			defer cur.Close(ctx)
	
			var gatheringsWithHost0 []models.Gathering
			for cur.Next(ctx) {
				var gathering models.Gathering
				if err := cur.Decode(&gathering); err != nil {
					log.Printf("Error decoding gathering: %v", err)
				}
				gatheringsWithHost0 = append(gatheringsWithHost0, gathering)
			}
	
			// Check if there are any gatherings found
			if len(gatheringsWithHost0) > 0 {
				log.Println("Found a client.", client.Username)
			} else {
				log.Println("No active gatherings are trying to join.", client.Username)
				rmcResponseStream.WriteU32LENext([]uint32{0})
			}

			rmcResponseBody := rmcResponseStream.Bytes()

			rmcResponse := nex.NewRMCResponse(nexproto.CustomMatchmakingProtocolID, callID)
			rmcResponse.SetSuccess(nexproto.RegisterGathering, rmcResponseBody)
		
			rmcResponseBytes := rmcResponse.Bytes()
		
			responsePacket, _ := nex.NewPacketV0(client, nil)
		
			responsePacket.SetVersion(0)
			responsePacket.SetSource(0x31)
			responsePacket.SetDestination(0x3F)
			responsePacket.SetType(nex.DataPacket)
		
			responsePacket.SetPayload(rmcResponseBytes)
		
			responsePacket.AddFlag(nex.FlagNeedsAck)
			responsePacket.AddFlag(nex.FlagReliable)
		
			SecureServer.Send(responsePacket)
		}else if roomCode != "" {
			// Check for another gathering with the same room code and host equal to 1
			cur, err := gatheringCollection.Aggregate(nil, []bson.M{
				bson.M{"$match": bson.D{
					{"room_code", roomCode},
					{"host", 1}, // Ensure host is equal to 1
					{"creator", bson.D{{"$ne", client.Username}}}, // Don't find our own gathering
				}},
			})
				
			if err != nil {
				log.Printf("Could not find the requested gathering: %+v\n", err)
				SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
				return
			}
		
			var matchingGatherings []models.Gathering // Store matching gatherings in a slice
		
			for cur.Next(context.TODO()) {
				var matchingGathering models.Gathering
				if err := cur.Decode(&matchingGathering); err != nil {
					log.Printf("Error decoding gathering: %+v\n", err)
					SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
					return
				}
				matchingGatherings = append(matchingGatherings, matchingGathering)
			}
		
			rmcResponseStream := nex.NewStream()
			rmcResponseStream.Grow(50)
		
			if len(matchingGatherings) == 0 {
				log.Println("There are no active gatherings with the same room code and host 1.")
			} else {
				log.Printf("Found gatherings - telling client to attempt joining (%d)", len(matchingGatherings))
				rmcResponseStream.WriteU32LENext([]uint32{uint32(len(matchingGatherings))})
				for _, gathering := range matchingGatherings {
					var user models.User
		
					if err = usersCollection.FindOne(nil, bson.M{"username": gathering.Creator}).Decode(&user); err != nil {
						log.Printf("Could not find creator %s of gathering: %+v\n", gathering.Creator, err)
						SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
						return
					}
					rmcResponseStream.WriteBufferString("HarmonixGathering")
					rmcResponseStream.WriteU32LENext([]uint32{uint32(len(gathering.Contents) + 4)})
					rmcResponseStream.WriteU32LENext([]uint32{uint32(len(gathering.Contents))})
					rmcResponseStream.Grow(int64(len(gathering.Contents)))
					rmcResponseStream.WriteBytesNext(gathering.Contents[0:4])
					rmcResponseStream.WriteU32LENext([]uint32{user.PID})
					rmcResponseStream.WriteU32LENext([]uint32{user.PID})
					rmcResponseStream.WriteBytesNext(gathering.Contents[12:])
				}
			}
			rmcResponseBody := rmcResponseStream.Bytes()

			rmcResponse := nex.NewRMCResponse(nexproto.CustomMatchmakingProtocolID, callID)
			rmcResponse.SetSuccess(nexproto.RegisterGathering, rmcResponseBody)
		
			rmcResponseBytes := rmcResponse.Bytes()
		
			responsePacket, _ := nex.NewPacketV0(client, nil)
		
			responsePacket.SetVersion(0)
			responsePacket.SetSource(0x31)
			responsePacket.SetDestination(0x3F)
			responsePacket.SetType(nex.DataPacket)
		
			responsePacket.SetPayload(rmcResponseBytes)
		
			responsePacket.AddFlag(nex.FlagNeedsAck)
			responsePacket.AddFlag(nex.FlagReliable)
		
			SecureServer.Send(responsePacket)
		}else {
			cur, err := gatheringCollection.Aggregate(nil, []bson.M{
				bson.M{"$match": bson.D{
					// don't find our own gathering
					{
						Key:   "creator",
						Value: bson.D{{Key: "$ne", Value: client.Username}},
					},
					// only look for gatherings updated in the last 15 minutes
					{
						Key:   "last_updated",
						Value: bson.D{{Key: "$gt", Value: (time.Now().Unix()) - (15 * 60)}},
					},
					// don't look for gatherings in the "in song" state
					{
						Key:   "state",
						Value: bson.D{{Key: "$ne", Value: 2}},
					},
					// don't look for gatherings in the "on song select" state
					{
						Key:   "state",
						Value: bson.D{{Key: "$ne", Value: 6}},
					},
					// only look for public gatherings
					{
						Key:   "public",
						Value: bson.D{{Key: "$eq", Value: 1}},
					},
					// filter out gatherings with non-empty room codes
					{
						Key:   "room_code",
						Value: bson.D{{Key: "$eq", Value: ""}},
					},
				}},
				bson.M{"$sample": bson.M{"size": 10}},
			})
			
			if err != nil {
				log.Printf("Could not get a random gathering: %s\n", err)
				SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
				return
			}
			var gatherings = make([]models.Gathering, 0)
			for cur.Next(nil) {
				var g models.Gathering
				err = cur.Decode(&g)
				if err != nil {
					log.Printf("Error decoding gathering: %+v\n", err)
					SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
					return
				}
				gatherings = append(gatherings, g)
			}
		
			rmcResponseStream := nex.NewStream()
			rmcResponseStream.Grow(50)
		
			// if there are no availble gatherings, tell the client to check again.
			// otherwise, pass the available gathering to the client
			if len(gatherings) == 0 {
				log.Println("There are no active gatherings. Tell client to keep checking")
				rmcResponseStream.WriteU32LENext([]uint32{0})
			} else {
				log.Printf("Found %d gatherings - telling client to attempt joining", len(gatherings))
				rmcResponseStream.WriteU32LENext([]uint32{uint32(len(gatherings))})
				for _, gathering := range gatherings {
					var user models.User
		
					if err = usersCollection.FindOne(nil, bson.M{"username": gathering.Creator}).Decode(&user); err != nil {
						log.Printf("Could not find creator %s of gathering: %+v\n", gathering.Creator, err)
						SendErrorCode(SecureServer, client, nexproto.CustomMatchmakingProtocolID, callID, 0x00010001)
						return
					}
					log.Println(user.PID)
					rmcResponseStream.WriteBufferString("HarmonixGathering")
					rmcResponseStream.WriteU32LENext([]uint32{uint32(len(gathering.Contents) + 4)})
					rmcResponseStream.WriteU32LENext([]uint32{uint32(len(gathering.Contents))})
					rmcResponseStream.Grow(int64(len(gathering.Contents)))
					rmcResponseStream.WriteBytesNext(gathering.Contents[0:4])
					rmcResponseStream.WriteU32LENext([]uint32{user.PID})
					rmcResponseStream.WriteU32LENext([]uint32{user.PID})
					rmcResponseStream.WriteBytesNext(gathering.Contents[12:])
				}
			}
			rmcResponseBody := rmcResponseStream.Bytes()

			rmcResponse := nex.NewRMCResponse(nexproto.CustomMatchmakingProtocolID, callID)
			rmcResponse.SetSuccess(nexproto.RegisterGathering, rmcResponseBody)
		
			rmcResponseBytes := rmcResponse.Bytes()
		
			responsePacket, _ := nex.NewPacketV0(client, nil)
		
			responsePacket.SetVersion(0)
			responsePacket.SetSource(0x31)
			responsePacket.SetDestination(0x3F)
			responsePacket.SetType(nex.DataPacket)
		
			responsePacket.SetPayload(rmcResponseBytes)
		
			responsePacket.AddFlag(nex.FlagNeedsAck)
			responsePacket.AddFlag(nex.FlagReliable)
		
			SecureServer.Send(responsePacket)
		}
	}

}
