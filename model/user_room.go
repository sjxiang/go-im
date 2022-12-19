package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// 对应 QQ 群里有哪些人
type UserRoom struct {
	Identity      string `bson:"identity"`
	UserIdentity  string `bson:"user_identity"`
	RoomIdentity  string `bson:"room_identity"`
	RoomType      int    `bson:"room_type"`  // 房间类型 1-私聊【好友】 2-群聊【陌生】
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


func JudgeUserIsFriend(userIdentity1, userIdentity2 string) (bool, error) {
	// 查询 userIdentity1 单聊房间列表
	cursor, err := MongoClient.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{"user_identity", userIdentity1}, {"room_type", 1} ,})
	
	roomIdentities := make([]string, 0)
	
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		return false, err
	}

	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Printf("[Decode ERROR]: %v\n", err)
			return false, err
		}

		roomIdentities = append(roomIdentities, ur.RoomIdentity)
	}
	
	// 获取关联 userIdentity2 单聊房间
	cnt, err := MongoClient.Collection(UserRoom{}.CollectionName()).CountDocuments(context.Background(), 
		bson.M{"user_identity": userIdentity2, "room_identity": bson.M{"$in":roomIdentities}})
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		return false, err
	}
	
	if cnt > 0 {
		return true, nil 
	}

	return false, nil
}



func InsertOneUserRoom(ur *UserRoom) error {
	_, err := MongoClient.Collection(UserRoom{}.CollectionName()).InsertOne(context.Background(), ur)
	return err
}