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

// 通过 identity、room_identity 查找 user_room 文档记录
func GetUserRoomByUserIdentityRoomIdentity(userIdentity, roomIdentity string) (*UserRoom, error) {
	ur := new(UserRoom)

	err := MongoClient.Collection(UserRoom{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"user_identity", userIdentity}, {"room_identity", roomIdentity}}).
		Decode(ur)
	return ur, err
}


func GetUserRoomByRoomIdentity(roomIdentity string) ([]*UserRoom, error) {
	cursor, err := MongoClient.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{ {"room_identity" , roomIdentity }})
	if err != nil {
		return nil, err
	}

	urs := make([]*UserRoom, 0)

	// 批量查询
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err = cursor.Decode(ur)
		if err != nil {
			return nil, err
		}
		urs = append(urs, ur)
	}

	return urs, nil

}