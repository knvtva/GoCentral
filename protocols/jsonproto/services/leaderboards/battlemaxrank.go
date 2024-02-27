package leaderboard

import (
	//"context"
	"log"
	"rb3server/protocols/jsonproto/marshaler"

	"github.com/ihatecompvir/nex-go"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BattleMaxrankGetRequest struct {
	Region      string `json:"region"`
	SystemMS    int    `json:"system_ms"`
	BattleID    int    `json:"battle_id"`
	MachineID   string `json:"machine_id"`
	SessionGUID string `json:"session_guid"`
	PID000      int    `json:"pid000"`
}

type BattleMaxrankGetResponse struct {
	MaxRank int `json:"max_rank"`
}

type BattleMaxrankGetService struct {
}

func (service BattleMaxrankGetService) Path() string {
	return "leaderboards/battle_maxrank/get"
}

func (service BattleMaxrankGetService) Handle(data string, database *mongo.Database, client *nex.Client) (string, error) {
	var req BattleMaxrankGetRequest

	err := marshaler.UnmarshalRequest(data, &req)
	if err != nil {
		return "", err
	}

	if req.PID000 != int(client.PlayerID()) {
		log.Println("Client-supplied PID did not match server-assigned PID, rejecting request for Battleomplishment leaderboards")
		return "", err
	}

	res := []BattleMaxrankGetResponse{{
		int(1),
	}}

	return marshaler.MarshalResponse(service.Path(), res)
}
