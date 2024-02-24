package message

import (
	"context"
	"database/sql"
	"time"
)

type MessageModel struct {
	DB *sql.DB
}

const timeout time.Duration = 2 * time.Second

func (m *MessageModel) CreateMessage(c context.Context, req CreateMessageReq) (*CreateMessageRes, error) {
	var res CreateMessageRes

	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	query := `INSERT INTO messages (text, room_id, user_id) VALUES ($1, $2, $3) RETURNING id`
	err := m.DB.QueryRowContext(ctx, query,
		req.Content,
		req.RoomId,
		req.UserId,
	).Scan(&res.Id)
	if err != nil {
		return &CreateMessageRes{}, err
	}

	return &res, nil
}
