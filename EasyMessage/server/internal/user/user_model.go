package user

import (
	"MyProject/EasyMessage/server/util"
	"context"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (m *UserModel) Insert(c context.Context, u *CreateUserRequest) (*User, error) {
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()

	user := &User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	if err := ValidateUser(*user); err != nil {
		return &User{}, err
	}
	user.ShortName = util.GenerateShortName(user.Name)
	user.Password, _ = util.HashPassword(user.Password)
	user.Email = strings.ToLower(user.Email)

	query := `INSERT INTO users (name, short_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id`

	args := []any{
		user.Name,
		user.ShortName,
		user.Email,
		user.Password,
	}

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id)
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (m *UserModel) GetUserByEmail(c context.Context, email string) (*User, error) {
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()

	query := `SELECT id, name, short_name, email, password FROM users WHERE email = $1`

	var user User

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.ShortName,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserModel) Update(c context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()

	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $3`

	hashedPassword, _ := util.HashPassword(user.Password)
	user.Password = hashedPassword
	user.Email = strings.ToLower(user.Email)

	args := []any{
		user.Name,
		user.Email,
		user.Password,
	}

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.Id,
		&user.Name,
		&user.ShortName,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

type MyJWTClaims struct {
	Id       uint64 `json:"id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

func (m *UserModel) Login(c context.Context, u *LoginUserRequest) (*LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()

	user, err := m.GetUserByEmail(ctx, u.Email)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	err = util.HashPasswordCheck(u.Password, user.Password)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		Id:       user.Id,
		UserName: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	secret := []byte(os.Getenv("JWT_SECRET"))
	ss, err := token.SignedString(secret)

	if err != nil {
		return &LoginUserResponse{}, err
	}

	return &LoginUserResponse{ss, user.Name, user.Id}, nil
}
