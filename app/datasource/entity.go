package datasource

import "time"

type User struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type CreateUserReq struct {
	Name string `json:"name"`
}
