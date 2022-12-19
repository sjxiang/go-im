package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
)


type MessageBasic struct {
	Identity      string `bson:"identity"`
	UserIdentity  string `bson:"user_identity"`
	RoomIdentity  string `bson:"room_identity"`
	Data          string `bson:"data"`

	CreatedAt     int64  `bson:"created_at"`
	UpdatedAt     int64  `bson:"updated_at"`
}


func (MessageBasic) CollectionName() string {
	return "message_basic"
}



func InsertOneMessageBasic(mb *MessageBasic) error {
	_, err := MongoClient.Collection(MessageBasic{}.CollectionName()).InsertOne(context.Background(), mb)
	return err
}


func GetMessageListByRoomIdentity(roomIdentity string, limit, skip int64) ([]*MessageBasic, error) {
	data := make([]*MessageBasic, 0)
	cursor, err := MongoClient.Collection(MessageBasic{}.CollectionName()).
		Find(context.Background(), bson.M{"room_identity": roomIdentity }, &options.FindOptions{
			Limit: &limit,
			Skip: &skip,
			Sort: bson.D{{"created_at", -1 }},  // 降序
		})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		mb := new(MessageBasic)
		err = cursor.Decode(mb)
		if err != nil {
			return nil, err
		}
		data = append(data, mb)		
	}

	return data, nil
}

