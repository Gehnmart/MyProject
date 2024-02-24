package message

type message struct {
	Id        int    `json:"id"`
	RoomId    int    `json:"room_id"`
	UserId    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CreateMessageReq struct {
	UserId  int    `json:"user_id"`
	RoomId  int    `json:"room_id"`
	Content string `json:"content"`
}

type CreateMessageRes struct {
	Id int `json:"id"`
}
