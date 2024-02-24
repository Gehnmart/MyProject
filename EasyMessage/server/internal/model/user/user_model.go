package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"server/internal/util"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserModel struct {
	DB *sql.DB
}

const timeout time.Duration = 2 * time.Second

func (m *UserModel) CreateUser(c context.Context, req CreateUserReq) (*CreateUserRes, error) {
	var res CreateUserRes

	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	shortName := toShortName(req.Name)
	hashPassword, _ := util.HashPassword(req.Password)

	query := `INSERT INTO users (name, short_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id, name`
	err := m.DB.QueryRowContext(ctx, query,
		req.Name,
		shortName,
		req.Email,
		hashPassword).Scan(&res.Id, &res.Name)
	if err != nil {
		fmt.Println(err)
		return &CreateUserRes{}, err
	}

	return &res, nil
}

func (m *UserModel) GetUserByEmail(c context.Context, email string) (*User, error) {
	var user User

	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	query := `SELECT id, name, short_name, email, password FROM users WHERE email = $1`
	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.ShortName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

type MyJWTClaims struct {
	jwt.StandardClaims
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (m *UserModel) LoginUser(c context.Context, req LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	user, err := m.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = util.CompareHashAndPassword(req.Password, user.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyJWTClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Name,
	})

	secret := []byte(os.Getenv("JWT_SECRET"))
	ss, err := token.SignedString(secret)
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{Token: ss, Name: user.Name, Id: user.Id}, nil
}

func (m *UserModel) ParseToken(token_src string) (int, error) {
	token, err := jwt.ParseWithClaims(token_src, &MyJWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Auth type")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*MyJWTClaims)
	if !ok {
		return 0, errors.New("token claims are not type *MyJWTClaims")
	}

	return claims.Id, nil
}
