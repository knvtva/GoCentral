package battles

import (
	"log"
	"rb3server/models"
	"rb3server/protocols/jsonproto/marshaler"

	"github.com/ihatecompvir/nex-go"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetBattlesRequest struct {
	Region      string `json:"region"`
	Locale      string `json:"locale"`
	SystemMS    int    `json:"system_ms"`
	SongID      int    `json:"song_id"`
	MachineID   string `json:"machine_id"`
	SessionGUID string `json:"session_guid"`
	PID000      int    `json:"pid000"`
}

type GetBattlesResponse struct {
	Instrument   int    `json:"id"`
	PID          int    `json:"pid"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Type         int    `json:"type"`
	Owner        string `json:"owner"`
	OwnerGUID    string `json:"owner_guid"`
	GUID         string `json:"guid"`
	ArtURL       string `json:"battle_art"`
	TimeEndVal   int    `json:"time_left"`
	SongID000    int    `json:"s_id000"`
	SongName000  string `json:"s_name000"`
}

type GetBattlesService struct {
}

func (service GetBattlesService) Path() string {
	return "battles/closed/get"
}

func (service GetBattlesService) Handle(data string, database *mongo.Database, client *nex.Client) (string, error) {
	var req GetBattlesRequest

	err := marshaler.UnmarshalRequest(data, &req)
	if err != nil {
		return "", err
	}

	if req.PID000 != int(client.PlayerID()) {
		log.Println("Client-supplied PID did not match server-assigned PID, rejecting getting battles")
		return "", err
	}

	battleCollection := database.Collection("battles")

	battleCursor, err := battleCollection.Find(nil, bson.D{})

	log.Println(battleCursor)

	if err != nil {
		log.Printf("Error getting setlist for battle: %s", err)
	}

	res := []GetBattlesResponse{}

	for battleCursor.Next(nil) {
		var battle GetBattlesResponse
		var battleToCopy models.Battles

		battleCursor.Decode(&battleToCopy)

		copier.Copy(&battle, &battleToCopy)

		res = append(res, battle)
	}

	if len(res) == 0 {
		return marshaler.MarshalResponse(service.Path(), []GetBattlesResponse{{}})
	} else {
		return marshaler.MarshalResponse(service.Path(), res)
	}
}
