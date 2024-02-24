package user

type User struct {
	Id        int
	Name      string
	ShortName string
	Email     string
	Password  string
}

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Id    int    `json:"id"`
}
