package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yeyee2901/sqlx/app/config"
	"github.com/yeyee2901/sqlx/app/controller"
	"github.com/yeyee2901/sqlx/app/middlewares"
	"github.com/yeyee2901/sqlx/docs"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	Config *config.Config
	DB     *sqlx.DB
	Router *gin.Engine
}

func (T *App) loadConfig() {
	config := config.LoadConfig("setting/setting.yaml")
	T.Config = config
}

func (T *App) initRouting() {
	c := controller.New(T.Router, T.Config, T.DB)
	c.InitRouting()
}

func (T *App) initSwagger() {
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "SQLX di Golang"
	docs.SwaggerInfo.Description = "Percobaan mapping object dari database menggunakan sqlx"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = T.Config.App.Listener
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func (T *App) initDatabase() {
	dsName := mysql.Config{
		User:                 T.Config.DBConfig.Username,
		Passwd:               T.Config.DBConfig.Password,
		Net:                  "tcp",
		Addr:                 T.Config.DBConfig.Host,
		DBName:               T.Config.DBConfig.DB,
		ParseTime:            T.Config.DBConfig.ParseTime,
		AllowNativePasswords: true,
	}

	db, err := sqlx.Connect("mysql", dsName.FormatDSN())
	if err != nil {
		log.Fatal("[DB INIT] ", err.Error())
	}

    if db == nil {
		log.Fatal("[DB INIT] Database is dead. Make sure the host is alive!")
    }

	T.DB = db
	T.DB.SetMaxIdleConns(1)
	T.DB.SetMaxOpenConns(10)
}

func main() {
	// init aplikasi
	app := App{}

	// LOAD: config
	app.loadConfig()

	// INIT: DB
	app.initDatabase()

	// INIT: gin
	if app.Config.App.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// middlewares
	router.Use(middlewares.CORS())

	app.Router = router
	app.initRouting()

	// INIT: swagger
	app.initSwagger()

	// populate swagger docs hanya untuk mode development
	if app.Config.App.Mode != "production" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	listener := &http.Server{
		Addr:         app.Config.App.Listener,
		Handler:      app.Router,
		ReadTimeout:  65 * time.Second,
		WriteTimeout: 65 * time.Second,
	}

	// berdasarkan docs gin:
	go func() {
		// catch error nya in case nge crash
		err := listener.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Server Closed...", err)
		}
	}()

	// EXIT:
	// close menunggu signal dari OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shutting down server")

	// clear semua proses yg msh menggantung
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := listener.Shutdown(ctx); err != nil {
		fmt.Println("Server Forced Shutdown ")
	}

	fmt.Println("Server exiting...")
}
