package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeyee2901/sqlx/app/datasource"
	"github.com/yeyee2901/sqlx/app/entity"
)

type UserService struct {
	DataSource *datasource.DataSource
	GinContext *gin.Context
}

func NewUserService(ds *datasource.DataSource, ctx *gin.Context) *UserService {
	return &UserService{
		ds,
		ctx,
	}
}

func (us *UserService) GetUser() (users []User, errResp entity.ResponseWithHTTPStatus) {
	// cek query string
	var reqQuery ReqGetUser
	err := us.GinContext.ShouldBind(&reqQuery)

	// lets just handle the error :)
	if err != nil {
		errResp.HttpStatus = http.StatusBadRequest
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Bad Request - %s", err.Error())
		return
	}

	// case user tidak masukin id
	// maka get all user
	if len(reqQuery.Id) == 0 {
		// cek jika user memasukkan nilai floating
		_, err = strconv.ParseInt(reqQuery.Id, 10, 0)
		if err != nil {
			errResp.HttpStatus = http.StatusBadRequest
			errResp.Details.Code = API_CODE_USER
			errResp.Details.Msg = fmt.Sprintf("Bad Request - User 'id' must be an integer")
		}

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
	err = us.DataSource.GetUserById(&singleUser, reqQuery.Id)

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

func (us *UserService) CreateUser() (newUser RespCreateUser, errResp entity.ResponseWithHTTPStatus) {
	var req ReqCreateUser
	err := us.GinContext.ShouldBindJSON(&req)
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
		errResp.Details.Msg = "Bad Request. User 'name' must not be an empty string"
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

func (us *UserService) DeleteUserById() (errResp entity.ResponseWithHTTPStatus) {
	id := us.GinContext.Param("id")

	// handle error in case terjadi keanehan dengan path params :)
	if len(id) < 1 {
		errResp.HttpStatus = http.StatusBadRequest
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Bad Request - Invalid mandatory field 'id'")
		return
	}

	// handle jika dimasukkan parameter nilai floating
	idInteger, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		errResp.HttpStatus = http.StatusBadRequest
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Bad Request - user 'id' must be an integer")
		return
	}

	// have to typecast here
	rowsAffected, err := us.DataSource.DeleteUserById(int(idInteger))
	if err != nil {
		errResp.HttpStatus = http.StatusInternalServerError
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Internal Server Error - %s", err.Error())
		return
	}

	// some custom message
	if rowsAffected == 0 {
		errResp.HttpStatus = http.StatusNotFound
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Content not found - no such user with id = %d", idInteger)
		return
	}

	errResp.HttpStatus = http.StatusOK
	errResp.Details.Code = API_CODE_USER
	errResp.Details.Msg = fmt.Sprintf("Success - %d rows affected", rowsAffected)

	return
}

func (us *UserService) UpdateUserById() (errResp entity.ResponseWithHTTPStatus) {
	// cek request body
	var req ReqUpdateUserById
	err := us.GinContext.ShouldBindJSON(&req)
	if err != nil {
		errResp.HttpStatus = http.StatusBadRequest
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Bad Request - %s", err.Error())
		return
	}

	// cek db dulu, ada user nya ato ngga
	var user User
	err = us.DataSource.GetUserById(&user, fmt.Sprintf("%d", req.Id))
	if err != nil {
		errResp.HttpStatus = http.StatusNotFound
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Content Not Found - No such user with id = %d", req.Id)
		return
	}

	// pass by value aja karena kita don't care apa yang
	// terjadi sama object 'req'
	err = us.DataSource.UpdateUserById(req)
	if err != nil {
		errResp.HttpStatus = http.StatusInternalServerError
		errResp.Details.Code = API_CODE_USER
		errResp.Details.Msg = fmt.Sprintf("Internal Server Error - %s", err.Error())
		return
	}

	// success
	errResp.HttpStatus = http.StatusOK
	errResp.Details.Code = API_CODE_USER
	errResp.Details.Msg = "Success - successfully updated user data"
	return
}
