package scores

import (
	"log"
	"context"
	"rb3server/protocols/jsonproto/marshaler"

	"github.com/ihatecompvir/nex-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BattleRecordRequest struct {
	Region      string `json:"region"`
	SystemMS    int    `json:"system_ms"`
	MachineID   string `json:"machine_id"`
	SessionGUID string `json:"session_guid"`
	Score 		int    `json:"score"`
	PID 		int    `json:"pid000"`
	BattleID    int    `json:"battle_id"` 
	Slot000     int    `json:"slot000"`
}

type BattleRecordResponse struct {
	ID           int    `json:"id"`
	IsBOI        int    `json:"is_boi"`
	InstaRank    int    `json:"insta_rank"`
	IsPercentile int    `json:"is_percentile"`
	Part1        string `json:"part_1"`
	Part2        string `json:"part_2"`
	Slot         int    `json:"slot"`
}

type BattleRecordService struct {
}

func (service BattleRecordService) Path() string {
	return "battles/record"
}

func (service BattleRecordService) Handle(data string, database *mongo.Database, client *nex.Client) (string, error) {
	var req BattleRecordRequest

	err := marshaler.UnmarshalRequest(data, &req)
	if err != nil {
		return "", err
	}

	battlesCollection := database.Collection("battle-scores") 
	if err != nil {
		log.Println("Error:", err)
		return "", err
	}

	playerData := bson.D{
		{Key: "battle_id", Value: req.BattleID},
		{Key: "pid", Value: req.PID},
		{Key: "score", Value: req.Score},
	}
	
	_, err = battlesCollection.InsertOne(context.TODO(), playerData)
	if err != nil {
		log.Println("Error:", err)
		return "", err
	}

  // This is mock data. TODO: Calculate the Instarank properly with the battle leaderboards.
	return marshaler.MarshalResponse(service.Path(), []ScoreRecordResponse{{2, 0, 8, 1, "d|2466|9", "j", 2}, {2, 1, 41, 1, "d|427|42", "j", 2}})
}
