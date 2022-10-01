package user

import "time"

const (
    API_CODE_USER = "01"
)

type User struct {
	Id        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type GetUserResp struct {
	Total int    `json:"total"`
	Users []User `json:"users"`
}

type CreateUserReq struct {
	Name string `json:"name"`
}

type CreateUserResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UpdateUserByIdReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
