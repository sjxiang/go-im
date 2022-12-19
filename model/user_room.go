package model


import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)


// 对应 QQ 群里有哪些人
type UserRoom struct {
	UserIdentity  string `bson:"user_identity"`
	RoomIdentity  string `bson:"room_identity"`
	CreatedAt     int64  `bson:"created_at"`
	UpdatedAt     int64  `bson:"updated_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}


func GetUserRoomByUserIdentityRoomIdentity(userIdentity, roomIdentity string) (*UserRoom, error) {
	ur := new(UserRoom)

	err := MongoClient.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"user_identity", userIdentity}, {"room_identity", roomIdentity}}).
		Decode(ur)
	return ur, err
}
