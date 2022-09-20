package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/yeyee2901/sqlx/app/config"
	"github.com/yeyee2901/sqlx/app/datasource"
	"github.com/yeyee2901/sqlx/app/entity"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Controller struct {
	Router *gin.Engine
	Config *config.Config
	DB     *sqlx.DB
}

func New(r *gin.Engine, c *config.Config, db *sqlx.DB) *Controller {
	return &Controller{
		Router: r,
		Config: c,
		DB:     db,
	}
}

func (T *Controller) InitRouting() {
	// Endpoint GET sebisa mungkin singular (bukan plural)
	// yang membedakan get all & get sesuai kriteria dari params nya aja
	T.Router.GET("/user", T.GetUser)
	T.Router.POST("/user", T.CreateUser)
	T.Router.PUT("/user", T.UpdateUserById)
	T.Router.DELETE("/user/:id", T.DeleteUserById)
}

// GetUser godoc
// @Tags User
// @Summary mengambil data-data user
// @Router /user [get]
// @Param id query int false "Jika tidak memasukkan user ID maka akan get semua"
// @Success 200 {object} entity.GetUsersResp
// @Failure 500 {object} entity.Response
func (T *Controller) GetUser(ctx *gin.Context) {
	ds := datasource.NewDatasource(T.Config, T.DB)

	// cek query string
	id := ctx.Query("id")

	// case user tidak masukin id
	// maka get all user
	if len(id) == 0 {
		users, err := ds.GetAllUsers()
		if err != nil {
			resp := entity.Response{
				Msg: err.Error(),
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			return
		}
		resp := entity.GetUsersResp{
			Total: len(users),
			Users: users,
		}

		// send to client
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// kasus lain ketika user memasukkan id,
	// maka ambil yang id nya sesuai saja
	user, err := ds.GetUserById(id)
	if err != nil {
		resp := entity.GetUsersResp{
			Total: 0,
			Users: []datasource.User{},
		}
		ctx.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	// untuk menyamakan respon saja, biar simetris
	// tapi len nya pasti 1 kalau memasukkan id
	var users []datasource.User
	users = append(users, user)
	resp := entity.GetUsersResp{
		Total: len(users),
		Users: users,
	}

	// send to client
	ctx.JSON(http.StatusOK, resp)
	return
}

// CreateUser godoc
// @Tags User
// @Summary Membuat user baru
// @Router /user [post]
// @Param request body datasource.CreateUserReq true "request body JSON"
// @Success 200 {object} entity.Response
func (T *Controller) CreateUser(ctx *gin.Context) {
	ds := datasource.NewDatasource(T.Config, T.DB)
	// binding ke model
	var req datasource.CreateUserReq
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		resp := entity.Response{
			Msg: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
	}

	userId, err := ds.CreateUser(&req)
	if err != nil {
		resp := entity.Response{
			Msg: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
	}

	resp := entity.CreateUserResp{
		Id:   userId,
		Name: req.Name,
	}
	ctx.JSON(http.StatusOK, resp)
}

// DeleteUserById godoc
// @Tags User
// @Summary Menghapus data user by ID
// @Router /user/{id} [delete]
// @Param id path int true "User ID (angka positif)"
// @Success 200 {object} entity.Response
func (T *Controller) DeleteUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		resp := entity.Response{
			Msg: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp := entity.Response{
		Msg: fmt.Sprintf("ID: %d", idInt),
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateUserById godoc
// @Tags User
// @Summary mengubah data user berdasarkan ID
// @Router /user [put]
// @Param request body datasource.UpdateUserByIdReq true "User ID (angka positif)"
// @Success 200 {object} entity.Response
func (T *Controller) UpdateUserById(ctx *gin.Context) {
	ds := datasource.NewDatasource(T.Config, T.DB)
	var req datasource.UpdateUserByIdReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		resp := entity.Response{
			Msg: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
	}

	err = ds.UpdateUserById(&req)
	if err != nil {
		resp := entity.Response{
			Msg: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
	}

    resp := entity.Response{
        Msg: "Sukses",
    }
    ctx.JSON(http.StatusOK, resp)
}
