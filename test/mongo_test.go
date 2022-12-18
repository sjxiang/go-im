package test

import (
	"fmt"
	"time"

	"github.com/sjxiang/go-im/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"

	"testing"
)

func TestFindOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "admin",
	}).ApplyURI("mongodb://172.17.0.2:27017"))
	if err != nil {
		t.Fatal(err)
	}

	db := client.Database("im")
	
	ub := new(model.UserBasic)
	err = db.Collection("user_basic").FindOne(context.Background(), bson.D{}).Decode(ub)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("=> %+v\n", ub)
}


func TestFind(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "admin",
	}).ApplyURI("mongodb://172.17.0.2:27017"))
	if err != nil {
		t.Fatal(err)
	}

	db := client.Database("im")
	
	cursor, err := db.Collection("user_room").Find(context.Background(), bson.D{})
	if err != nil {
		t.Fatal(err)
	}

	urs := make([]*model.UserRoom, 0)

	// 批量查询
	for cursor.Next(context.Background()) {
		ur := new(model.UserRoom)
		err = cursor.Decode(ur)
		if err != nil {
			t.Fatal(err)
		}
		urs = append(urs, ur)
	}

	for _, v := range urs {
		fmt.Println("user_room => ", v)
	}

}