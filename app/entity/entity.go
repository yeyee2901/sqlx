package entity

import "github.com/yeyee2901/sqlx/app/datasource"

type Response struct {
	Msg string `json:"msg"`
}

type GetUsersResp struct {
	Total int               `json:"total"`
	Users []datasource.User `json:"users"`
}


type CreateUserResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
