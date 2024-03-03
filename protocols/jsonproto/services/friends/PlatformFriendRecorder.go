/*
I have no idea if this works for wii. Perhaps we should check that sometime.
Works with PS3, RPCS3 (Xbox Live would work if GoCentral was supported)
*/
package friends

import (
	"log"
	"reflect"
	"fmt"
	
	"rb3server/protocols/jsonproto/marshaler"

	"github.com/ihatecompvir/nex-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FriendRecordRequest struct {
    Region      string `json:"region"`
    SystemMS    int    `json:"system_ms"`
    MachineID   string `json:"machine_id"`
    SessionGUID string `json:"session_guid"`
	PID         int    `json:"pid"`
    Name000     string `json:"name000"`
    Guid000     string `json:"guid000"`
    Name001     string `json:"name001"`
    Guid001     string `json:"guid001"`
    Name002     string `json:"name002"`
    Guid002     string `json:"guid002"`
    Name003     string `json:"name003"`
    Guid003     string `json:"guid003"`
    Name004     string `json:"name004"`
    Guid004     string `json:"guid004"`
    Name005     string `json:"name005"`
    Guid005     string `json:"guid005"`
    Name006     string `json:"name006"`
    Guid006     string `json:"guid006"`
    Name007     string `json:"name007"`
    Guid007     string `json:"guid007"`
    Name008     string `json:"name008"`
    Guid008     string `json:"guid008"`
    Name009     string `json:"name009"`
    Guid009     string `json:"guid009"`
    Name010     string `json:"name010"`
    Guid010     string `json:"guid010"`
    Name011     string `json:"name011"`
    Guid011     string `json:"guid011"`
    Name012     string `json:"name012"`
    Guid012     string `json:"guid012"`
    Name013     string `json:"name013"`
    Guid013     string `json:"guid013"`
    Name014     string `json:"name014"`
    Guid014     string `json:"guid014"`
    Name015     string `json:"name015"`
    Guid015     string `json:"guid015"`
    Name016     string `json:"name016"`
    Guid016     string `json:"guid016"`
    Name017     string `json:"name017"`
    Guid017     string `json:"guid017"`
    Name018     string `json:"name018"`
    Guid018     string `json:"guid018"`
    Name019     string `json:"name019"`
    Guid019     string `json:"guid019"`
    Name020     string `json:"name020"`
    Guid020     string `json:"guid020"`
    Name021     string `json:"name021"`
    Guid021     string `json:"guid021"`
    Name022     string `json:"name022"`
    Guid022     string `json:"guid022"`
    Name023     string `json:"name023"`
    Guid023     string `json:"guid023"`
    Name024     string `json:"name024"`
    Guid024     string `json:"guid024"`
    Name025     string `json:"name025"`
    Guid025     string `json:"guid025"`
    Name026     string `json:"name026"`
    Guid026     string `json:"guid026"`
    Name027     string `json:"name027"`
    Guid027     string `json:"guid027"`
    Name028     string `json:"name028"`
    Guid028     string `json:"guid028"`
    Name029     string `json:"name029"`
    Guid029     string `json:"guid029"`
    Name030     string `json:"name030"`
    Guid030     string `json:"guid030"`
    Name031     string `json:"name031"`
    Guid031     string `json:"guid031"`
    Name032     string `json:"name032"`
    Guid032     string `json:"guid032"`
    Name033     string `json:"name033"`
    Guid033     string `json:"guid033"`
    Name034     string `json:"name034"`
    Guid034     string `json:"guid034"`
    Name035     string `json:"name035"`
    Guid035     string `json:"guid035"`
    Name036     string `json:"name036"`
    Guid036     string `json:"guid036"`
    Name037     string `json:"name037"`
    Guid037     string `json:"guid037"`
    Name038     string `json:"name038"`
    Guid038     string `json:"guid038"`
    Name039     string `json:"name039"`
    Guid039     string `json:"guid039"`
    Name040     string `json:"name040"`
    Guid040     string `json:"guid040"`
    Name041     string `json:"name041"`
    Guid041     string `json:"guid041"`
    Name042     string `json:"name042"`
    Guid042     string `json:"guid042"`
    Name043     string `json:"name043"`
    Guid043     string `json:"guid043"`
    Name044     string `json:"name044"`
    Guid044     string `json:"guid044"`
    Name045     string `json:"name045"`
    Guid045     string `json:"guid045"`
    Name046     string `json:"name046"`
    Guid046     string `json:"guid046"`
    Name047     string `json:"name047"`
    Guid047     string `json:"guid047"`
    Name048     string `json:"name048"`
    Guid048     string `json:"guid048"`
    Name049     string `json:"name049"`
    Guid049     string `json:"guid049"`
    Name050     string `json:"name050"`
    Guid050     string `json:"guid050"`
    Name051     string `json:"name051"`
    Guid051     string `json:"guid051"`
    Name052     string `json:"name052"`
    Guid052     string `json:"guid052"`
    Name053     string `json:"name053"`
    Guid053     string `json:"guid053"`
    Name054     string `json:"name054"`
    Guid054     string `json:"guid054"`
    Name055     string `json:"name055"`
    Guid055     string `json:"guid055"`
    Name056     string `json:"name056"`
    Guid056     string `json:"guid056"`
    Name057     string `json:"name057"`
    Guid057     string `json:"guid057"`
    Name058     string `json:"name058"`
    Guid058     string `json:"guid058"`
    Name059     string `json:"name059"`
    Guid059     string `json:"guid059"`
    Name060     string `json:"name060"`
    Guid060     string `json:"guid060"`
    Name061     string `json:"name061"`
    Guid061     string `json:"guid061"`
    Name062     string `json:"name062"`
    Guid062     string `json:"guid062"`
    Name063     string `json:"name063"`
    Guid063     string `json:"guid063"`
    Name064     string `json:"name064"`
    Guid064     string `json:"guid064"`
    Name065     string `json:"name065"`
    Guid065     string `json:"guid065"`
    Name066     string `json:"name066"`
    Guid066     string `json:"guid066"`
    Name067     string `json:"name067"`
    Guid067     string `json:"guid067"`
    Name068     string `json:"name068"`
    Guid068     string `json:"guid068"`
    Name069     string `json:"name069"`
    Guid069     string `json:"guid069"`
    Name070     string `json:"name070"`
    Guid070     string `json:"guid070"`
    Name071     string `json:"name071"`
    Guid071     string `json:"guid071"`
    Name072     string `json:"name072"`
    Guid072     string `json:"guid072"`
    Name073     string `json:"name073"`
    Guid073     string `json:"guid073"`
    Name074     string `json:"name074"`
    Guid074     string `json:"guid074"`
    Name075     string `json:"name075"`
    Guid075     string `json:"guid075"`
    Name076     string `json:"name076"`
    Guid076     string `json:"guid076"`
    Name077     string `json:"name077"`
    Guid077     string `json:"guid077"`
    Name078     string `json:"name078"`
    Guid078     string `json:"guid078"`
    Name079     string `json:"name079"`
    Guid079     string `json:"guid079"`
    Name080     string `json:"name080"`
    Guid080     string `json:"guid080"`
    Name081     string `json:"name081"`
    Guid081     string `json:"guid081"`
    Name082     string `json:"name082"`
    Guid082     string `json:"guid082"`
    Name083     string `json:"name083"`
    Guid083     string `json:"guid083"`
    Name084     string `json:"name084"`
    Guid084     string `json:"guid084"`
    Name085     string `json:"name085"`
    Guid085     string `json:"guid085"`
    Name086     string `json:"name086"`
    Guid086     string `json:"guid086"`
    Name087     string `json:"name087"`
    Guid087     string `json:"guid087"`
    Name088     string `json:"name088"`
    Guid088     string `json:"guid088"`
    Name089     string `json:"name089"`
    Guid089     string `json:"guid089"`
    Name090     string `json:"name090"`
    Guid090     string `json:"guid090"`
    Name091     string `json:"name091"`
    Guid091     string `json:"guid091"`
    Name092     string `json:"name092"`
    Guid092     string `json:"guid092"`
    Name093     string `json:"name093"`
    Guid093     string `json:"guid093"`
    Name094     string `json:"name094"`
    Guid094     string `json:"guid094"`
    Name095     string `json:"name095"`
    Guid095     string `json:"guid095"`
    Name096     string `json:"name096"`
    Guid096     string `json:"guid096"`
    Name097     string `json:"name097"`
    Guid097     string `json:"guid097"`
    Name098     string `json:"name098"`
    Guid098     string `json:"guid098"`
    Name099     string `json:"name099"`
    Guid099     string `json:"guid099"`
}

type FriendRecordResponse struct {
	Success int `json:"success"`
}

type FriendRecordService struct {
}

func (service FriendRecordService) Path() string {
	return "leaderboards/friends/update"
}

func (service FriendRecordService) Handle(data string, database *mongo.Database, client *nex.Client) (string, error) {
	var req FriendRecordRequest
	err := marshaler.UnmarshalRequest(data, &req)
	if err != nil {
		return "", err
	}
	
	if req.PID != int(client.PlayerID()) {
		log.Println("Client-supplied PID did not match server-assigned PID, rejecting PlatformFriendRecorder upload.")
	}

	friendsCollection := database.Collection("platform-friends") 

	var friendDocument bson.M
	err = friendsCollection.FindOne(nil, bson.M{"pid": req.PID}).Decode(&friendDocument)

	if err != nil {
		friendDocument = bson.M{
			"pid": req.PID,
		}
	}

	reqValue := reflect.ValueOf(&req).Elem()
	for i := 0; i <= 99; i++ {
		fieldName := fmt.Sprintf("name%03d", i)
		FriendID := reqValue.FieldByName(fmt.Sprintf("Name%03d", i)).Interface().(string)
		if FriendID != "" {
			friendDocument[fieldName] = FriendID
		}
	}

	_, err = friendsCollection.ReplaceOne(nil, bson.M{"pid": req.PID}, friendDocument, options.Replace().SetUpsert(true))

	if err != nil {
		log.Println("Error: ", err)
		return "", err
	}

	res := []FriendRecordResponse{{1}}

	return marshaler.MarshalResponse(service.Path(), res)
}
