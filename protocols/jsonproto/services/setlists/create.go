package setlists

import (
	"log"
	"rb3server/protocols/jsonproto/marshaler"

	"github.com/ihatecompvir/nex-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type SetlistCreateRequest struct {
	Type	    int    `json:"type"`
	Name  		string `json:"name"`	
	Region 		string `json:"region"`
	Description string `json:"description"`
	Flags		int	   `json:"flags"`
	SystemMS 	int	   `json:"system_ms"`
	MachineID   string `json:"machine_id"`
	SessionGUID string `json:"session_guid"`
	PID         int    `json:"pid"`
	Shared		string `json:"shared"`
	ListGuid	string `json:"list_guid"`
	SongID000   int    `json:"song_id000"`
	SongID001   int    `json:"song_id001"`
	SongID002   int    `json:"song_id002"`
}

type SetlistCreateResponse struct {
	PID     int `json:"pid"`
	Creator int `json:"creator"`
}

type SetlistCreateService struct {
}

func (service SetlistCreateService) Path() string {
	return "setlists/update"
}

func (service SetlistCreateService) Handle(data string, database *mongo.Database, client *nex.Client) (string, error) {
	var req SetlistCreateRequest
	err := marshaler.UnmarshalRequest(data, &req)
	if err != nil {
		return "", err
	}

	log.Println("[GoCentral] Setlist Creation Triggered")


	res := []SetlistCreateResponse{{
		req.PID,
		0,
	}}

	return marshaler.MarshalResponse(service.Path(), res)
}