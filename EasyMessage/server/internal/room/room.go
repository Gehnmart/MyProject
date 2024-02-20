package room

import (
	"context"
	"database/sql"
	"time"
)

type Room struct {
	Id          uint64   `json:"id"`
	Name        string   `json:"name"`
	Users_id    []uint64 `json:"user_id"`
	Messages_id []uint64 `json:"messages_id"`
}

type RoomModel struct {
	DB *sql.DB
}

func (m *RoomModel) Insert(room *Room) error {
	query := `INSERT INTO rooms (name, users_id, messages_id) VALUES ($1, $2, $3) RETURNING id`

	args := []any{
		room.Name,
		room.Users_id,
		room.Messages_id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&room.Id)

	if err != nil {
		return err
	}
	return nil
}

func (m *RoomModel) GetRoomsByUserId(userId uint64) ([]Room, error) {
	query := `SELECT id, name, user_id, messages_id FROM rooms WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []Room
	for rows.Next() {
		var room Room
		err := rows.Scan(&room.Id, &room.Name, &room.Users_id, &room.Messages_id)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
