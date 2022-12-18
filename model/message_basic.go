package model

import (
	// "context"

	// "go.mongodb.org/mongo-driver/bson"
)


type MessageBasic struct {
	UserIdentity  string `bson:"user_identity"`
	RoomIdentity  string `bson:"room_identity"`
	Data          string `bson:"data"`
	CreatedAt     int64  `bson:"created_at"`
	UpdatedAt     int64  `bson:"updated_at"`
}

func (MessageBasic) CollectionName() string {
	return "message_basic"
}


// func GetUserBasicByUsernamePassword(account, password string) (*UserBasic, error) {
// 	ub := new(UserBasic)

// 	// 文档记录 => 结构体
// 	err := MongoClient.Collection(UserBasic{}.CollectionName()).
// 		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
// 		Decode(ub)
// 	return ub, err
// }

