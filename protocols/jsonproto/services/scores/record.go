package scores

import (
	"context"
	"log"
	"rb3server/models"
	"rb3server/protocols/jsonproto/marshaler"
	"strings"
	"sort"

	"github.com/knvtva/nex-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instrumentMap = map[int]int{
	0: 1,
	1: 2,
	2: 4,
	3: 8,
	4: 16,
	5: 32,
	6: 64,
	7: 128,
	8: 256,
	9: 512,
}

type ScoreRecordRequestOnePlayer struct {
	Region           string `json:"region"`
	Locale           string `json:"locale"`
	SystemMS         int    `json:"system_ms"`
	SongID           int    `json:"song_id"`
	MachineID        string `json:"machine_id"`
	SessionGUID      string `json:"session_guid"`
	PID000           int    `json:"pid000"`
	BoiID            int    `json:"boi_id"`
	BandMask         int    `json:"band_mask"`
	ProvideInstaRank int    `json:"provide_insta_rank"`
	
	// band stuff
	RoleID000  int `json:"role_id000"`
	Score000   int `json:"score000"`
	Stars000   int `json:"stars000"`
	Slot000    int `json:"slot000"`
	DiffID000  int `json:"diff_id000"`
	CScore000  int `json:"c_score000"`
	CCScore000 int `json:"cc_score000"`
	Percent000 int `json:"percent000"`

	// individual contributors
	RoleID001  int `json:"role_id001"`
	Score001   int `json:"score001"`
	Stars001   int `json:"stars001"`
	PID001     int `json:"pid001"`
	Slot001    int `json:"slot001"`
	DiffID001  int `json:"diff_id001"`
	CScore001  int `json:"c_score001"`
	CCScore001 int `json:"cc_score001"`
	Percent001 int `json:"percent001"`
}

type ScoreRecordRequestTwoPlayer struct {
	Region           string `json:"region"`
	Locale           string `json:"locale"`
	SystemMS         int    `json:"system_ms"`
	SongID           int    `json:"song_id"`
	MachineID        string `json:"machine_id"`
	SessionGUID      string `json:"session_guid"`
	PID000           int    `json:"pid000"`
	BoiID            int    `json:"boi_id"`
	BandMask         int    `json:"band_mask"`
	ProvideInstaRank int    `json:"provide_insta_rank"`
	OvershellSlot 	 int	`json:"slot"`

	// band stuff
	RoleID000  int `json:"role_id000"`
	Score000   int `json:"score000"`
	Stars000   int `json:"stars000"`
	Slot000    int `json:"slot000"`
	DiffID000  int `json:"diff_id000"`
	CScore000  int `json:"c_score000"`
	CCScore000 int `json:"cc_score000"`
	Percent000 int `json:"percent000"`

	// individual contributors
	RoleID001  int `json:"role_id001"`
	Score001   int `json:"score001"`
	Stars001   int `json:"stars001"`
	PID001     int `json:"pid001"`
	Slot001    int `json:"slot001"`
	DiffID001  int `json:"diff_id001"`
	CScore001  int `json:"c_score001"`
	CCScore001 int `json:"cc_score001"`
	Percent001 int `json:"percent001"`

	RoleID002  int `json:"role_id002"`
	Score002   int `json:"score002"`
	Stars002   int `json:"stars002"`
	PID002     int `json:"pid002"`
	Slot002    int `json:"slot002"`
	DiffID002  int `json:"diff_id002"`
	CScore002  int `json:"c_score002"`
	CCScore002 int `json:"cc_score002"`
	Percent002 int `json:"percent002"`

	RoleID003  int `json:"role_id003"`
	Score003   int `json:"score003"`
	Stars003   int `json:"stars003"`
	PID003     int `json:"pid003"`
	Slot003    int `json:"slot003"`
	DiffID003  int `json:"diff_id003"`
	CScore003  int `json:"c_score003"`
	CCScore003 int `json:"cc_score003"`
	Percent003 int `json:"percent003"`
}

type ScoreRecordRequestThreePlayer struct {
	Region           string `json:"region"`
	Locale           string `json:"locale"`
	SystemMS         int    `json:"system_ms"`
	SongID           int    `json:"song_id"`
	MachineID        string `json:"machine_id"`
	SessionGUID      string `json:"session_guid"`
	PID000           int    `json:"pid000"`
	BoiID            int    `json:"boi_id"`
	BandMask         int    `json:"band_mask"`
	ProvideInstaRank int    `json:"provide_insta_rank"`
	OvershellSlot 	 int	`json:"slot"`

	// band stuff
	RoleID000  int `json:"role_id000"`
	Score000   int `json:"score000"`
	Stars000   int `json:"stars000"`
	Slot000    int `json:"slot000"`
	DiffID000  int `json:"diff_id000"`
	CScore000  int `json:"c_score000"`
	CCScore000 int `json:"cc_score000"`
	Percent000 int `json:"percent000"`

	// individual contributors
	RoleID001  int `json:"role_id001"`
	Score001   int `json:"score001"`
	Stars001   int `json:"stars001"`
	PID001     int `json:"pid001"`
	Slot001    int `json:"slot001"`
	DiffID001  int `json:"diff_id001"`
	CScore001  int `json:"c_score001"`
	CCScore001 int `json:"cc_score001"`
	Percent001 int `json:"percent001"`

	RoleID002  int `json:"role_id002"`
	Score002   int `json:"score002"`
	Stars002   int `json:"stars002"`
	PID002     int `json:"pid002"`
	Slot002    int `json:"slot002"`
	DiffID002  int `json:"diff_id002"`
	CScore002  int `json:"c_score002"`
	CCScore002 int `json:"cc_score002"`
	Percent002 int `json:"percent002"`

	RoleID003  int `json:"role_id003"`
	Score003   int `json:"score003"`
	Stars003   int `json:"stars003"`
	PID003     int `json:"pid003"`
	Slot003    int `json:"slot003"`
	DiffID003  int `json:"diff_id003"`
	CScore003  int `json:"c_score003"`
	CCScore003 int `json:"cc_score003"`
	Percent003 int `json:"percent003"`

	RoleID004  int `json:"role_id004"`
	Score004   int `json:"score004"`
	Stars004   int `json:"stars004"`
	PID004     int `json:"pid004"`
	Slot004    int `json:"slot004"`
	DiffID004  int `json:"diff_id004"`
	CScore004  int `json:"c_score004"`
	CCScore004 int `json:"cc_score004"`
	Percent004 int `json:"percent004"`

	RoleID005  int `json:"role_id005"`
	Score005   int `json:"score005"`
	Stars005   int `json:"stars005"`
	PID005     int `json:"pid005"`
	Slot005    int `json:"slot005"`
	DiffID005  int `json:"diff_id005"`
	CScore005  int `json:"c_score005"`
	CCScore005 int `json:"cc_score005"`
	Percent005 int `json:"percent005"`
}

type ScoreRecordRequestFourPlayer struct {
	Region           string `json:"region"`
	Locale           string `json:"locale"`
	SystemMS         int    `json:"system_ms"`
	SongID           int    `json:"song_id"`
	MachineID        string `json:"machine_id"`
	SessionGUID      string `json:"session_guid"`
	PID000           int    `json:"pid000"`
	BoiID            int    `json:"boi_id"`
	BandMask         int    `json:"band_mask"`
	ProvideInstaRank int    `json:"provide_insta_rank"`
	OvershellSlot 	 int	`json:"slot"`

	// band stuff
	RoleID000  int `json:"role_id000"`
	Score000   int `json:"score000"`
	Stars000   int `json:"stars000"`
	Slot000    int `json:"slot000"`
	DiffID000  int `json:"diff_id000"`
	CScore000  int `json:"c_score000"`
	CCScore000 int `json:"cc_score000"`
	Percent000 int `json:"percent000"`

	// individual contributors
	RoleID001  int `json:"role_id001"`
	Score001   int `json:"score001"`
	Stars001   int `json:"stars001"`
	PID001     int `json:"pid001"`
	Slot001    int `json:"slot001"`
	DiffID001  int `json:"diff_id001"`
	CScore001  int `json:"c_score001"`
	CCScore001 int `json:"cc_score001"`
	Percent001 int `json:"percent001"`

	RoleID002  int `json:"role_id002"`
	Score002   int `json:"score002"`
	Stars002   int `json:"stars002"`
	PID002     int `json:"pid002"`
	Slot002    int `json:"slot002"`
	DiffID002  int `json:"diff_id002"`
	CScore002  int `json:"c_score002"`
	CCScore002 int `json:"cc_score002"`
	Percent002 int `json:"percent002"`

	RoleID003  int `json:"role_id003"`
	Score003   int `json:"score003"`
	Stars003   int `json:"stars003"`
	PID003     int `json:"pid003"`
	Slot003    int `json:"slot003"`
	DiffID003  int `json:"diff_id003"`
	CScore003  int `json:"c_score003"`
	CCScore003 int `json:"cc_score003"`
	Percent003 int `json:"percent003"`

	RoleID004  int `json:"role_id004"`
	Score004   int `json:"score004"`
	Stars004   int `json:"stars004"`
	PID004     int `json:"pid004"`
	Slot004    int `json:"slot004"`
	DiffID004  int `json:"diff_id004"`
	CScore004  int `json:"c_score004"`
	CCScore004 int `json:"cc_score004"`
	Percent004 int `json:"percent004"`

	RoleID005  int `json:"role_id005"`
	Score005   int `json:"score005"`
	Stars005   int `json:"stars005"`
	PID005     int `json:"pid005"`
	Slot005    int `json:"slot005"`
	DiffID005  int `json:"diff_id005"`
	CScore005  int `json:"c_score005"`
	CCScore005 int `json:"cc_score005"`
	Percent005 int `json:"percent005"`

	RoleID006  int `json:"role_id006"`
	Score006   int `json:"score006"`
	Stars006   int `json:"stars006"`
	PID006     int `json:"pid006"`
	Slot006    int `json:"slot006"`
	DiffID006  int `json:"diff_id006"`
	CScore006  int `json:"c_score006"`
	CCScore006 int `json:"cc_score006"`
	Percent006 int `json:"percent006"`

	RoleID007  int `json:"role_id007"`
	Score007   int `json:"score007"`
	Stars007   int `json:"stars007"`
	PID007     int `json:"pid007"`
	Slot007    int `json:"slot007"`
	DiffID007  int `json:"diff_id007"`
	CScore007  int `json:"c_score007"`
	CCScore007 int `json:"cc_score007"`
	Percent007 int `json:"percent007"`
}

type ScoreRecordRequestFivePlayer struct {
	Region           string `json:"region"`
	Locale           string `json:"locale"`
	SystemMS         int    `json:"system_ms"`
	SongID           int    `json:"song_id"`
	MachineID        string `json:"machine_id"`
	SessionGUID      string `json:"session_guid"`
	PID000           int    `json:"pid000"`
	BoiID            int    `json:"boi_id"`
	BandMask         int    `json:"band_mask"`
	ProvideInstaRank int    `json:"provide_insta_rank"`
	OvershellSlot 	 int	`json:"slot"`

	// band stuff
	RoleID000  int `json:"role_id000"`
	Score000   int `json:"score000"`
	Stars000   int `json:"stars000"`
	Slot000    int `json:"slot000"`
	DiffID000  int `json:"diff_id000"`
	CScore000  int `json:"c_score000"`
	CCScore000 int `json:"cc_score000"`
	Percent000 int `json:"percent000"`

	// individual contributors
	RoleID001  int `json:"role_id001"`
	Score001   int `json:"score001"`
	Stars001   int `json:"stars001"`
	PID001     int `json:"pid001"`
	Slot001    int `json:"slot001"`
	DiffID001  int `json:"diff_id001"`
	CScore001  int `json:"c_score001"`
	CCScore001 int `json:"cc_score001"`
	Percent001 int `json:"percent001"`

	RoleID002  int `json:"role_id002"`
	Score002   int `json:"score002"`
	Stars002   int `json:"stars002"`
	PID002     int `json:"pid002"`
	Slot002    int `json:"slot002"`
	DiffID002  int `json:"diff_id002"`
	CScore002  int `json:"c_score002"`
	CCScore002 int `json:"cc_score002"`
	Percent002 int `json:"percent002"`

	RoleID003  int `json:"role_id003"`
	Score003   int `json:"score003"`
	Stars003   int `json:"stars003"`
	PID003     int `json:"pid003"`
	Slot003    int `json:"slot003"`
	DiffID003  int `json:"diff_id003"`
	CScore003  int `json:"c_score003"`
	CCScore003 int `json:"cc_score003"`
	Percent003 int `json:"percent003"`

	RoleID004  int `json:"role_id004"`
	Score004   int `json:"score004"`
	Stars004   int `json:"stars004"`
	PID004     int `json:"pid004"`
	Slot004    int `json:"slot004"`
	DiffID004  int `json:"diff_id004"`
	CScore004  int `json:"c_score004"`
	CCScore004 int `json:"cc_score004"`
	Percent004 int `json:"percent004"`

	RoleID005  int `json:"role_id005"`
	Score005   int `json:"score005"`
	Stars005   int `json:"stars005"`
	PID005     int `json:"pid005"`
	Slot005    int `json:"slot005"`
	DiffID005  int `json:"diff_id005"`
	CScore005  int `json:"c_score005"`
	CCScore005 int `json:"cc_score005"`
	Percent005 int `json:"percent005"`

	RoleID006  int `json:"role_id006"`
	Score006   int `json:"score006"`
	Stars006   int `json:"stars006"`
	PID006     int `json:"pid006"`
	Slot006    int `json:"slot006"`
	DiffID006  int `json:"diff_id006"`
	CScore006  int `json:"c_score006"`
	CCScore006 int `json:"cc_score006"`
	Percent006 int `json:"percent006"`

	RoleID007  int `json:"role_id007"`
	Score007   int `json:"score007"`
	Stars007   int `json:"stars007"`
	PID007     int `json:"pid007"`
	Slot007    int `json:"slot007"`
	DiffID007  int `json:"diff_id007"`
	CScore007  int `json:"c_score007"`
	CCScore007 int `json:"cc_score007"`
	Percent007 int `json:"percent007"`

	RoleID008  int `json:"role_id008"`
	Score008   int `json:"score008"`
	Stars008   int `json:"stars008"`
	PID008     int `json:"pid008"`
	Slot008    int `json:"slot008"`
	DiffID008  int `json:"diff_id008"`
	CScore008  int `json:"c_score008"`
	CCScore008 int `json:"cc_score008"`
	Percent008 int `json:"percent008"`

	RoleID009  int `json:"role_id009"`
	Score009   int `json:"score009"`
	Stars009   int `json:"stars009"`
	PID009     int `json:"pid009"`
	Slot009    int `json:"slot009"`
	DiffID009  int `json:"diff_id009"`
	CScore009  int `json:"c_score009"`
	CCScore009 int `json:"cc_score009"`
	Percent009 int `json:"percent009"`
}

type ScoreRecordResponse struct {
	ID           int    `json:"id"`
	IsBOI        int    `json:"is_boi"`
	InstaRank    int    `json:"insta_rank"`
	IsPercentile int    `json:"is_percentile"`
	Part1        string `json:"part_1"`
	Part2        string `json:"part_2"`
	Slot         int    `json:"slot"`
}

type ScoreRecordService struct {
}

func (service ScoreRecordService) Path() string {
	return "scores/record"
}

func (service ScoreRecordService) Handle(data string, database *mongo.Database, client *nex.Client) (string, error) {
	var req interface{}
	var playerData []bson.D
	
	// check for number of players so we can parse the message correctly
	if strings.Contains(data, "slot009") {
		req = ScoreRecordRequestFivePlayer{}
		log.Println(data)
	} else if strings.Contains(data, "slot007") {
		req = ScoreRecordRequestFourPlayer{}
		log.Println(data)
	} else if strings.Contains(data, "slot005") {
		req = ScoreRecordRequestThreePlayer{}
		log.Println(data)
	} else if strings.Contains(data, "slot003") {
		req = ScoreRecordRequestTwoPlayer{}
		log.Println(data)
	} else {
		req = ScoreRecordRequestOnePlayer{}
		log.Println(data)
	}

	var err error

	// TODO: Make this not so horrible
	// this is an unholy abomination
	var songID int
	switch request := req.(type) {
	case ScoreRecordRequestOnePlayer:
		err = marshaler.UnmarshalRequest(data, &request)
		if err != nil {
			return "", err
		}
		songID = request.SongID
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID000},
			{Key: "score", Value: request.Score000},
			{Key: "notespct", Value: request.Percent000},
			{Key: "role_id", Value: request.RoleID000},
			{Key: "diffid", Value: request.DiffID000},
			{Key: "boi", Value: 0},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID001},
			{Key: "score", Value: request.Score001},
			{Key: "notespct", Value: request.Percent001},
			{Key: "role_id", Value: request.RoleID001},
			{Key: "diffid", Value: request.DiffID001},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		log.Println("Slot:", request.Slot001)
	case ScoreRecordRequestTwoPlayer:
		err = marshaler.UnmarshalRequest(data, &request)
		if err != nil {
			return "", err
		}
		if request.PID000 != int(client.PlayerID()) {
			users := database.Collection("users")
			var user models.User
			err = users.FindOne(nil, bson.M{"pid": request.PID000}).Decode(&user)
			log.Println("Client-supplied PID did not match server-assigned PID, rejecting recording score")
			client.SetPlayerID(user.PID)
		}
		songID = request.SongID
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID000},
			{Key: "score", Value: request.Score000},
			{Key: "notespct", Value: request.Percent000},
			{Key: "role_id", Value: request.RoleID000},
			{Key: "diffid", Value: request.DiffID000},
			{Key: "boi", Value: 0},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID001},
			{Key: "score", Value: request.Score001},
			{Key: "notespct", Value: request.Percent001},
			{Key: "role_id", Value: request.RoleID001},
			{Key: "diffid", Value: request.DiffID001},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID002},
			{Key: "score", Value: request.Score002},
			{Key: "notespct", Value: request.Percent002},
			{Key: "role_id", Value: request.RoleID002},
			{Key: "diffid", Value: request.DiffID002},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: instrumentMap[request.RoleID002]},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID003},
			{Key: "score", Value: request.Score003},
			{Key: "notespct", Value: request.Percent003},
			{Key: "role_id", Value: request.RoleID003},
			{Key: "diffid", Value: request.DiffID003},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: instrumentMap[request.RoleID003]},
		})
		log.Println("Slot000:", request.Slot000)
		log.Println("Slot001:", request.Slot001)
		log.Println("Slot002:", request.Slot002)
		log.Println("Slot003:", request.Slot003)
	case ScoreRecordRequestThreePlayer:
		err = marshaler.UnmarshalRequest(data, &request)
		if err != nil {
			return "", err
		}
		if request.PID000 != int(client.PlayerID()) {
			users := database.Collection("users")
			var user models.User
			err = users.FindOne(nil, bson.M{"pid": request.PID000}).Decode(&user)
			log.Println("Client-supplied PID did not match server-assigned PID, rejecting recording score")
			client.SetPlayerID(user.PID)
		}
		songID = request.SongID
		// Band Scores Are Applied Here
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID000},
			{Key: "score", Value: request.Score000},
			{Key: "notespct", Value: request.Percent000},
			{Key: "role_id", Value: request.RoleID000},
			{Key: "diffid", Value: request.DiffID000},
			{Key: "boi", Value: 0},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID001},
			{Key: "score", Value: request.Score001},
			{Key: "notespct", Value: request.Percent001},
			{Key: "role_id", Value: request.RoleID001},
			{Key: "diffid", Value: request.DiffID001},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID002},
			{Key: "score", Value: request.Score002},
			{Key: "notespct", Value: request.Percent002},
			{Key: "role_id", Value: request.RoleID002},
			{Key: "diffid", Value: request.DiffID002},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		// Individual Scores Are Applied Here
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID003},
			{Key: "score", Value: request.Score003},
			{Key: "notespct", Value: request.Percent003},
			{Key: "role_id", Value: request.RoleID003},
			{Key: "diffid", Value: request.DiffID003},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: instrumentMap[request.RoleID003]},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID004},
			{Key: "score", Value: request.Score004},
			{Key: "notespct", Value: request.Percent004},
			{Key: "role_id", Value: request.RoleID004},
			{Key: "diffid", Value: request.DiffID004},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: instrumentMap[request.RoleID004]},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID005},
			{Key: "score", Value: request.Score005},
			{Key: "notespct", Value: request.Percent005},
			{Key: "role_id", Value: request.RoleID005},
			{Key: "diffid", Value: request.DiffID005},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: instrumentMap[request.RoleID005]},
		})
		log.Println("Slot:", request.Slot001)
	case ScoreRecordRequestFourPlayer:
		err = marshaler.UnmarshalRequest(data, &request)
		if err != nil {
			return "", err
		}
		if request.PID000 != int(client.PlayerID()) {
			users := database.Collection("users")
			var user models.User
			err = users.FindOne(nil, bson.M{"pid": request.PID000}).Decode(&user)
			log.Println("Client-supplied PID did not match server-assigned PID, rejecting recording score")
			client.SetPlayerID(user.PID)
		}
		songID = request.SongID
		// Band Scores Are Applied Here
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID000},
			{Key: "score", Value: request.Score000},
			{Key: "notespct", Value: request.Percent000},
			{Key: "role_id", Value: request.RoleID000},
			{Key: "diffid", Value: request.DiffID000},
			{Key: "boi", Value: 0},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID001},
			{Key: "score", Value: request.Score001},
			{Key: "notespct", Value: request.Percent001},
			{Key: "role_id", Value: request.RoleID001},
			{Key: "diffid", Value: request.DiffID001},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID002},
			{Key: "score", Value: request.Score002},
			{Key: "notespct", Value: request.Percent002},
			{Key: "role_id", Value: request.RoleID002},
			{Key: "diffid", Value: request.DiffID002},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: request.BandMask},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID003},
			{Key: "score", Value: request.Score003},
			{Key: "notespct", Value: request.Percent003},
			{Key: "role_id", Value: request.RoleID003},
			{Key: "diffid", Value: request.DiffID003},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: request.BandMask},
		})

		// Individual Scores Are Applied Here
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID004},
			{Key: "score", Value: request.Score004},
			{Key: "notespct", Value: request.Percent004},
			{Key: "role_id", Value: request.RoleID004},
			{Key: "diffid", Value: request.DiffID004},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: instrumentMap[request.RoleID004]},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID005},
			{Key: "score", Value: request.Score005},
			{Key: "notespct", Value: request.Percent005},
			{Key: "role_id", Value: request.RoleID005},
			{Key: "diffid", Value: request.DiffID005},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: instrumentMap[request.RoleID004]},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID006},
			{Key: "score", Value: request.Score006},
			{Key: "notespct", Value: request.Percent006},
			{Key: "role_id", Value: request.RoleID006},
			{Key: "diffid", Value: request.DiffID006},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: instrumentMap[request.RoleID006]},
		})
		playerData = append(playerData, bson.D{
			{Key: "song_id", Value: request.SongID},
			{Key: "pid", Value: request.PID007},
			{Key: "score", Value: request.Score007},
			{Key: "notespct", Value: request.Percent007},
			{Key: "role_id", Value: request.RoleID007},
			{Key: "diffid", Value: request.DiffID007},
			{Key: "boi", Value: 1},
			{Key: "instrument_mask", Value: instrumentMap[request.RoleID007]},
		})
		log.Println("Slot:", request.Slot001)
	}

	upsert := true
	//var results []ScoreRecordResponse
	var rank int
	var Bandrank int
	var beatenRankUser string
	var beatenBandRankUser string


	for i := 0; i < len(playerData); i++ {
		var playerScore models.Score // Assuming you have a Score model
		pid := playerData[i][1].Value.(int)
		score := playerData[i][2].Value.(int)
		role_id := playerData[i][4].Value.(int)

		playerFilter := bson.M{"song_id": songID, "pid": pid, "role_id": role_id}
		err := database.Collection("scores").FindOne(context.TODO(), playerFilter).Decode(&playerScore)

		if err != nil {
			log.Printf("Error while retrieving player score: %v\n", err)
		}

		if score >= playerScore.Score {
			_, err = database.Collection("scores").ReplaceOne(context.TODO(), bson.M{"song_id": songID, "pid": pid, "role_id": role_id}, playerData[i], &options.ReplaceOptions{Upsert: &upsert})
			if err != nil {
				log.Printf("Could not upsert score for song ID %v: %v\n", songID, err)
				// Handle the error appropriately, e.g., return an error or continue
			}
		}

		playerFilter = bson.M{"song_id": songID, "pid": pid, "role_id": role_id}
		err = database.Collection("scores").FindOne(context.TODO(), playerFilter).Decode(&playerScore)

		if err != nil {
			log.Printf("Error while retrieving player score: %v\n", err)
		}

		newScore := playerScore.Score

		// Use BandPosition if role_id is 10, otherwise use Position (Variables)
		if role_id == 10 {
			var bandScores []int
		
			playerFilter = bson.M{"song_id": songID, "role_id": role_id}
			cursor, err := database.Collection("scores").Find(context.TODO(), playerFilter)
			if err != nil {
				log.Printf("Error while retrieving scores: %v\n", err)
			}
			defer cursor.Close(context.TODO())
		
			for cursor.Next(context.TODO()) {
				var playerScore models.Score
				if err := cursor.Decode(&playerScore); err != nil {
					log.Printf("Error decoding score: %v\n", err)
					return "", err
				}
				bandScores = append(bandScores, playerScore.Score)
			}
		
			sort.Sort(sort.Reverse(sort.IntSlice(bandScores)))
		
			filteredBandScores := []int{}
			for _, score := range bandScores {
				if score != 0 {
					filteredBandScores = append(filteredBandScores, score)
				}
			}
		
			Bandrank = 0
			for i, score := range filteredBandScores {
				if newScore == score {
					Bandrank = i + 1
					break
				}
			}

			lastBandValue := filteredBandScores[len(filteredBandScores)-1]

			if lastBandValue == playerScore.Score {
				beatenRankUser = ""
			} else {
				var BeatenBandScore int

				BeatenBandScore = filteredBandScores[rank+2]

				playerFilter = bson.M{"song_id": songID, "role_id": role_id, "score": BeatenBandScore}
				cursor, err := database.Collection("scores").Find(context.TODO(), playerFilter)
				if err != nil {
					log.Printf("Error while retrieving BeatenBandScore: %v\n", err)
					return "", err
				}
				defer cursor.Close(context.TODO())

				var playerScore models.Score
				if cursor.Next(context.TODO()) {
					if err := cursor.Decode(&playerScore); err != nil {
						log.Printf("Error decoding score: %v\n", err)
						return "", err
					}
				}

				var users models.User
				playerFilter = bson.M{"pid": playerScore.OwnerPID}
				cursor, err = database.Collection("users").Find(context.TODO(), playerFilter)
				if err != nil {
					log.Printf("Error while retrieving user data: %v\n", err)
					return "", err
				}
				defer cursor.Close(context.TODO())

				if cursor.Next(context.TODO()) {
					if err := cursor.Decode(&users); err != nil {
						log.Printf("Error decoding user data: %v\n", err)
						return "", err
					}
				}
				beatenBandRankUser = users.Username
			}

		} else {
			var soloScores []int
		
			playerFilter = bson.M{"song_id": songID, "role_id": role_id}
			cursor, err := database.Collection("scores").Find(context.TODO(), playerFilter)
			if err != nil {
				log.Printf("Error while retrieving scores: %v\n", err)
			}
			defer cursor.Close(context.TODO())
		
			for cursor.Next(context.TODO()) {
				var playerScore models.Score
				if err := cursor.Decode(&playerScore); err != nil {
					log.Printf("Error decoding score: %v\n", err)
					return "", err
				}
				soloScores = append(soloScores, playerScore.Score)
			}
		
			sort.Sort(sort.Reverse(sort.IntSlice(soloScores)))
		
			filteredSoloScores := []int{}
			for _, score := range soloScores {
				if score != 0 {
					filteredSoloScores = append(filteredSoloScores, score)
				}
			}
		
			rank = 0
			for i, score := range filteredSoloScores {
				if newScore == score {
					rank = i + 1
					break
				}
			}

			lastSoloValue := filteredSoloScores[len(filteredSoloScores)-1]

			if lastSoloValue == playerScore.Score {
				beatenRankUser = ""
			} else {
				var BeatenUserScore int

				BeatenUserScore = filteredSoloScores[rank]

				playerFilter = bson.M{"song_id": songID, "role_id": role_id, "score": BeatenUserScore}
				cursor, err := database.Collection("scores").Find(context.TODO(), playerFilter)
				if err != nil {
					log.Printf("Error while retrieving BeatenUserScore: %v\n", err)
					return "", err
				}
				defer cursor.Close(context.TODO())

				var playerScore models.Score
				if cursor.Next(context.TODO()) {
					if err := cursor.Decode(&playerScore); err != nil {
						log.Printf("Error decoding score: %v\n", err)
						return "", err
					}
				}

				var users models.User
				playerFilter = bson.M{"pid": playerScore.OwnerPID}
				cursor, err = database.Collection("users").Find(context.TODO(), playerFilter)
				if err != nil {
					log.Printf("Error while retrieving user data: %v\n", err)
					return "", err
				}
				defer cursor.Close(context.TODO())

				if cursor.Next(context.TODO()) {
					if err := cursor.Decode(&users); err != nil {
						log.Printf("Error decoding user data: %v\n", err)
						return "", err
					}
				}
				beatenRankUser = users.Username
			}
		}
	}

	var BeatenRankUserRes string
	var BeatenBandRankUserRes string

	if beatenRankUser != "" {
		BeatenRankUserRes = "g|" + beatenRankUser + "|" + beatenRankUser
	} else {
		BeatenRankUserRes = "j"
	}

	if beatenBandRankUser != "" {
		BeatenBandRankUserRes = "g|" + beatenBandRankUser + "|" + beatenBandRankUser
	} else {
		BeatenBandRankUserRes = "j"
	}


	res := []ScoreRecordResponse{		
	{501, 0, rank, 1, "b", BeatenRankUserRes, 0},
	{501, 1, Bandrank, 1, "b", BeatenBandRankUserRes, 0},
	{502, 0, rank, 1, "b", BeatenRankUserRes, 1},
	{502, 1, Bandrank, 1, "b", BeatenBandRankUserRes, 1},
	{503, 0, rank, 1, "b", BeatenRankUserRes, 2},
	{503, 1, Bandrank, 1, "b", BeatenBandRankUserRes, 2},
	{503, 0, rank, 1, "b", BeatenRankUserRes, 2},
	{503, 1, Bandrank, 1, "b", BeatenBandRankUserRes, 2},
	{503, 0, rank, 1, "b", BeatenRankUserRes, 3},
	{503, 1, Bandrank, 1, "b", BeatenBandRankUserRes, 3},
	{503, 0, rank, 1, "b", BeatenRankUserRes, 3},
	{503, 1, Bandrank, 1, "b", BeatenBandRankUserRes, 3},
	}

	return marshaler.MarshalResponse(service.Path(), res)
}
