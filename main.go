package main

import (
	"log"

	_ "github.com/codingtroop/ubl-store/docs"
	"github.com/codingtroop/ubl-store/pkg/config"
	api "github.com/codingtroop/ubl-store/pkg/handlers"
	"github.com/codingtroop/ubl-store/pkg/helpers"
	"github.com/codingtroop/ubl-store/pkg/repositories/sqlite"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	hc := api.NewHealthCheckHandler()

	viper.AddConfigPath(".")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    //

	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	sqliteConnector := sqlite.NewSqliteConnector(configuration.Db.Path)
	db, err := sqliteConnector.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	if err := sqliteConnector.Init(db); err != nil {
		log.Fatal(err.Error())
	}

	ur := sqlite.NewSqliteUblRepository(db)
	ar := sqlite.NewSqliteAttanchmentRepository(db)

	us := helpers.NewIOStorer("ubls")
	as := helpers.NewIOStorer("attachments")
	c := helpers.NewGZip()
	u := helpers.NewUblExtension()

	uh := api.NewUblStoreHandler(ur, ar, us, as, c, u)

	e.GET("/health", hc.Live)
	e.GET("/health/live", hc.Live)
	e.GET("/health/ready", hc.Ready)

	ug := e.Group("/api/v1/ubl")

	ug.GET("/:id", uh.Get)
	ug.POST("", uh.Post)
	ug.DELETE("/:id", uh.Delete)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
