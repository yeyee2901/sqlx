package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeyee2901/sqlx/app/datasource"
	"github.com/yeyee2901/sqlx/app/entity"
)

type UserService struct {
	DataSource *datasource.DataSource
}

func NewUserService(ds *datasource.DataSource) *UserService {
	return &UserService{
		ds,
	}
}

func (us *UserService) GetUser(ctx *gin.Context, id string) (users []User, errResp entity.ResponseWithHTTPStatus) {
	// case user tidak masukin id
	// maka get all user
	if len(id) == 0 {
		err := us.DataSource.GetAllUsers(&users)

		if err != nil {
			// masalah pasti ada di DB / query nya
			errResp.HttpStatus = http.StatusInternalServerError
			errResp.Details.Code = API_CODE_USER
			errResp.Details.Msg = fmt.Sprintf("Internal Server Error - %s", err.Error())
		} else {
			// kalau tidak ada masalah berarti lempar ok
			errResp.HttpStatus = http.StatusOK
		}

		return
	}

	// kasus lain ketika user memasukkan id,
	// maka ambil yang id nya sesuai saja
	var singleUser User
	err := us.DataSource.GetUserById(&singleUser, id)

	// kalau error != nil berarti besar kemungkinan user dengan id tersebut tidak ada
	if err != nil {
		errResp.HttpStatus = http.StatusNotFound
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Content Not Found - %s", err.Error())
		return
	}

	errResp.HttpStatus = http.StatusOK
	users = append(users, singleUser)

	return
}

func (us *UserService) CreateUser(ctx *gin.Context) (newUser RespCreateUser, errResp entity.ResponseWithHTTPStatus) {
	var req ReqCreateUser
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		errResp.HttpStatus = http.StatusBadRequest
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Bad Request - %s", err.Error())
		return
	}

    // kalau user name kosong, maka kasih bad request
    if len(req.Name) < 1 {
		errResp.HttpStatus = http.StatusBadRequest
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = "Bad Request. User name must not be an empty string"
		return
    }

    // simpan di database
    userId, err := us.DataSource.CreateUser(req)
    if err != nil {
 		errResp.HttpStatus = http.StatusInternalServerError
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Internal Server Error - %s", err.Error()) 
		return
    }

    // success
    newUser.Id = int(userId)
    newUser.Name = req.Name
    errResp.HttpStatus = http.StatusOK

    return
}
