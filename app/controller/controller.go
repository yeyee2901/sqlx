// TODO ------
// - CreateUser
// - DeleteUserById
// - UpdateUserById
package controller

import (
	"net/http"

	"github.com/yeyee2901/sqlx/app/config"
	"github.com/yeyee2901/sqlx/app/datasource"
	"github.com/yeyee2901/sqlx/app/user"

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
// @Success 200 {object} user.GetUserResp
func (T *Controller) GetUser(ctx *gin.Context) {
	ds := datasource.NewDatasource(T.Config, T.DB)
	userService := user.NewUserService(ds)

	// cek query string
	id := ctx.Query("id")

	users, errResp := userService.GetUser(ctx, id)

	if errResp.HttpStatus != http.StatusOK {
		ctx.AbortWithStatusJSON(errResp.HttpStatus, errResp.Details)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Tags User
// @Summary Membuat user baru
// @Router /user [post]
// @Param request body user.CreateUserReq true "request body JSON"
// @Success 200 {object} entity.Response
func (T *Controller) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "in progress",
	})
}

// DeleteUserById godoc
// @Tags User
// @Summary Menghapus data user by ID
// @Router /user/{id} [delete]
// @Param id path int true "User ID (angka positif)"
// @Success 200 {object} entity.Response
func (T *Controller) DeleteUserById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "in progress",
	})
}

// UpdateUserById godoc
// @Tags User
// @Summary mengubah data user berdasarkan ID
// @Router /user [put]
// @Param request body user.UpdateUserByIdReq true "User ID (angka positif)"
// @Success 200 {object} entity.Response
func (T *Controller) UpdateUserById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "in progress",
	})
}
