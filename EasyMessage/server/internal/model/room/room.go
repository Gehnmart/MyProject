package room

type Room struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	FirstUserId  int    `json:"first_user_id"`
	SecondUserId int    `json:"second_user_id"`
}

type CreateRoomReq struct {
	Name         string `json:"name"`
	FirstUserId  int    `json:"first_user_id"`
	SecondUserId int    `json:"second_user_id"`
}

type CreateRoomRes struct {
	Id int `json:"id"`
}

type GetRoomByUserIdReq struct {
	UserId string `json:"user_id"`
}

type GetRoomByUserIdRes struct {
	Rooms []Room `json:"rooms"`
}
