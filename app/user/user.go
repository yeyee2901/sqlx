package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeyee2901/sqlx/app/datasource"
	"github.com/yeyee2901/sqlx/app/entity"
)

type UserService struct {
	ds *datasource.DataSource
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
        err := us.ds.GetAllUsers(&users)

        if err != nil {
            // masalah pasti ada di DB / query nya
            errResp.HttpStatus = http.StatusInternalServerError
            errResp.Details.Code = API_CODE_USER
            errResp.Details.Msg = err.Error()
        } else {
            // kalau tidak ada masalah berarti lempar ok
            errResp.HttpStatus = http.StatusOK
        }

		return
	}

	// kasus lain ketika user memasukkan id,
	// maka ambil yang id nya sesuai saja
	var singleUser User
    err := us.ds.GetUserById(&singleUser, id)

	// kalau error != nil berarti besar kemungkinan user dengan id tersebut tidak ada
	if err != nil {
        errResp.HttpStatus = http.StatusNotFound
        errResp.Details.Code = API_CODE_USER
        errResp.Details.Msg = err.Error()
        return
	}

    errResp.HttpStatus = http.StatusOK
	users = append(users, singleUser)

	return
}
