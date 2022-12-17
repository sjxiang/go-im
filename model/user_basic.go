package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserBasic struct {
	Identity  string `bson:"_id"`
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

	err := MongoClient.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}


func GetUserBasicByIdentity(identity primitive.ObjectID) (*UserBasic, error) {
	ub := new(UserBasic)

	err := MongoClient.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub)
	return ub, err
}