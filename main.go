package main

import (
	"fmt"
	"log"
	"strings"

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

	var configuration config.Configuration

	v := LoadConfig()

	err := v.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	sqliteConnector := sqlite.NewSqliteConnector(configuration.Db.Sqlite.Path)
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

	us := helpers.NewIOStorer(configuration.Storage.Filesystem.UblPath)
	as := helpers.NewIOStorer(configuration.Storage.Filesystem.AttachmentPath)
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

	e.Logger.Fatal(e.Start(":" + configuration.Port))
}

func LoadConfig() *viper.Viper {
	conf := viper.New()

	conf.AutomaticEnv()
	conf.SetEnvPrefix("ublstore")
	conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	conf.SetConfigName("config")
	conf.SetConfigType("yml")
	conf.AddConfigPath(".")
	err := conf.ReadInConfig()

	if err != nil {
		switch err.(type) {
		default:
			panic(fmt.Errorf("fatal error loading config file: %s", err))
		case viper.ConfigFileNotFoundError:
			fmt.Errorf("No config file found. Using defaults and environment variables")
		}
	}

	// workaround because viper does not treat env vars the same as other config
	for _, key := range conf.AllKeys() {
		val := conf.Get(key)
		conf.Set(key, val)
	}

	return conf
}
