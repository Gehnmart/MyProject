package data

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"
)

var (
	ErrDuplicateEmail  = errors.New("duplicate email")
	ErrRecordsNotFound = errors.New("records not found")
)

type User struct {
	Id        uint64   `json:"id"`
	Name      string   `json:"name"`
	ShortName string   `json:"short_name"`
	Email     string   `json:"-"`
	Password  string   `json:"-"`
	Rooms     []uint64 `json:"-"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(user *User) error {
	hashedPassword, _ := HashPassword(user.Password)
	if user.ShortName != "" {
		user.ShortName = GenerateShortName(user.Name)
	}
	user.Password = hashedPassword
	user.Email = strings.ToLower(user.Email)

	query := `INSERT INTO users (name, short_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id`

	args := []any{
		user.Name,
		user.ShortName,
		user.Email,
		user.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m *UserModel) GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, name, short_name, email, password, Rooms FROM users WHERE email = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.ShortName,
		&user.Email,
		&user.Password,
		&user.Rooms,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordsNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m *UserModel) Update(user *User) error {
	return nil
}
