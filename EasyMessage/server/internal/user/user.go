package user

import (
	"database/sql"
)

type User struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
	Id          uint64 `json:"id"`
}

type UserModel struct {
	DB *sql.DB
}
