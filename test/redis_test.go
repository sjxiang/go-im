package test


import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)


var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "172.20.0.1:6379",
	Password: "",  // no password set
	DB: 0,         // use default DB
})

var expiredTime = time.Hour * 1024

func TestSetValue(t *testing.T) {
	
	err := rdb.Set(ctx, "kpop", "jisoo", expiredTime).Err()
	if err != nil {
		t.Error(err)
	}
}


func TestGetValue(t *testing.T) {

	val, err := rdb.Get(ctx, "kpop").Result()
	if err != nil {
		t.Error(err)
	}

	t.Log(val)
}