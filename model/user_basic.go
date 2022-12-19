package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)


type UserBasic struct {
	Identity  string `bson:"identity"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}


func GetUserBasicByUsernamePassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)

	// 文档记录 => 结构体
	err := MongoClient.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}


func GetUserBasicByIdentity(identity string) (*UserBasic, error) {  
	ub := new(UserBasic)

	err := MongoClient.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"identity", identity}}).  // primitive.ObjectID
		Decode(ub)
	return ub, err
}


func GetUserBasicByEmail(email string) (int64, error) {
	return MongoClient.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{ "email", email }})
}

func GetUserBasicCountByAccount(account string) (int64, error) {
	return MongoClient.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{ "account", account }})
}

func InsertOneUserBasic(ub *UserBasic) error {
	_, err := MongoClient.Collection(UserBasic{}.CollectionName()).InsertOne(context.Background(), ub)
	return err
}