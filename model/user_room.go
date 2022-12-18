package model

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
