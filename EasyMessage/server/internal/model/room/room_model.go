package room

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type RoomModel struct {
	DB *sql.DB
}

const timeout time.Duration = 2 * time.Second

func (m *RoomModel) CreateRoom(c context.Context, req CreateRoomReq) (*CreateRoomRes, error) {
	var res CreateRoomRes

	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	query := `INSERT INTO rooms (name, first_user_id, second_user_id) VALUES ($1, $2, $3) RETURNING id`
	err := m.DB.QueryRowContext(ctx, query,
		req.Name,
		req.FirstUserId,
		req.SecondUserId,
	).Scan(&res.Id)
	if err != nil {
		return &CreateRoomRes{}, err
	}

	return &res, nil
}

func (m *RoomModel) GetRoomById(c context.Context, id int) (*Room, error) {
	var room Room

	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	query := `SELECT id, name, first_user_id, second_user_id FROM rooms WHERE id = $1`
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&room.Id,
		&room.Name,
		&room.FirstUserId,
		&room.SecondUserId,
	)
	if err != nil {
		return &Room{}, err
	}

	return &room, nil
}

func (m *RoomModel) GetRoomByUserId(c context.Context, userId int) (*GetRoomByUserIdRes, error) {
	var res GetRoomByUserIdRes

	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	fmt.Println(userId)

	query := `SELECT id, name, first_user_id, second_user_id FROM rooms WHERE first_user_id = $1 OR second_user_id = $1`
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return &GetRoomByUserIdRes{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var room Room
		err := rows.Scan(
			&room.Id,
			&room.Name,
			&room.FirstUserId,
			&room.SecondUserId,
		)
		if err != nil {
			return &GetRoomByUserIdRes{}, err
		}
		res.Rooms = append(res.Rooms, room)
	}

	return &res, nil
}
