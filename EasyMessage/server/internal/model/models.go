package model

import (
	"MyProject/EasyMessage/server/internal/message"
	"MyProject/EasyMessage/server/internal/room"
	"MyProject/EasyMessage/server/internal/user"
	"database/sql"
)

type Models struct {
	Users    user.UserModel
	Rooms    room.RoomModel
	Messages message.MessageModel
}

func NewModel(db *sql.DB) *Models {
	return &Models{
		Users:    user.UserModel{DB: db},
		Rooms:    room.RoomModel{DB: db},
		Messages: message.MessageModel{DB: db},
	}
}
