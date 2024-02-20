package message

import "database/sql"

type Message struct {
	Id        uint64   `json:"id"`
	Text      string   `json:"text"`
	Sender_id []uint64 `json:"sender_id"`
}

type MessageModel struct {
	DB *sql.DB
}
