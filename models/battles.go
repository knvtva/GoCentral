package models

type Battles struct {
	ID           string `bson:"_id"`
	Instrument   int    `bson:"id"`
	PID          int    `bson:"pid"`
	Title        string `bson:"title"`
	Desc         string `bson:"desc"`
	Type         int    `bson:"type"`
	Owner        string `bson:"owner"`
	OwnerGUID    string `bson:"owner_guid"`
	GUID         string `bson:"guid"`
	ArtURL       string `bson:"battle_url"`
	TimeEndVal   int    `bson:"time_left"`
	SongID000    int    `bson:"s_id000"`
	SongName000  string `bson:"s_name000"`
}