package datasource

import "time"

type User struct {
	Id        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type CreateUserReq struct {
	Name string `json:"name"`
}

type UpdateUserByIdReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
