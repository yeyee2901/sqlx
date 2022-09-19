package controller

import (
	"net/http"
	"sqlx/config"
	"sqlx/entity"

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
	T.Router.GET("/user", T.GetUsers)
	T.Router.GET("/user/:id", T.GetUserById)
	T.Router.POST("/user", T.CreateUser)
	T.Router.PUT("/user/:id", T.UpdateUserById)
	T.Router.DELETE("/user/:id", T.DeleteUserById)
}

// GetUsers godoc
// @Tags User
// @Summary mengambil data-data user
// @Router /user [get]
// @Success 200 {object} entity.Response
func (T *Controller) GetUsers(ctx *gin.Context) {
	resp := entity.Response{
		Msg: "Testing",
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetUserById godoc
// @Tags User
// @Summary mengambil data-data user berdasarkan ID
// @Router /user/:id [get]
// @Success 200 {object} entity.Response
func (T *Controller) GetUserById(ctx *gin.Context) {
	resp := entity.Response{
		Msg: "Testing",
	}
	ctx.JSON(http.StatusOK, resp)
}

// CreateUser godoc
// @Tags User
// @Summary Membuat user baru
// @Router /user [post]
// @Success 200 {object} entity.Response
func (T *Controller) CreateUser(ctx *gin.Context) {
	resp := entity.Response{
		Msg: "Testing",
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteUserById godoc
// @Tags User
// @Summary Menghapus data user by ID
// @Router /user/:id [delete]
// @Success 200 {object} entity.Response
func (T *Controller) DeleteUserById(ctx *gin.Context) {
	resp := entity.Response{
		Msg: "Testing",
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateUserById godoc
// @Tags User
// @Summary mengubah data user berdasarkan ID
// @Router /user/:id [put]
// @Success 200 {object} entity.Response
func (T *Controller) UpdateUserById(ctx *gin.Context) {
	resp := entity.Response{
		Msg: "Testing",
	}

	ctx.JSON(http.StatusOK, resp)
}
