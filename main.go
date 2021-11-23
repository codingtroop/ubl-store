package main

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/codingtroop/ubl-store/docs"
	"github.com/codingtroop/ubl-store/pkg/config"
	api "github.com/codingtroop/ubl-store/pkg/handlers"
	"github.com/codingtroop/ubl-store/pkg/helpers"
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

	us := helpers.NewIOStorer(configuration.Storage.Filesystem.UblPath)
	as := helpers.NewIOStorer(configuration.Storage.Filesystem.AttachmentPath)
	c := helpers.NewGZip()
	u := helpers.NewUblExtension()

	uh := api.NewUblStoreHandler(us, as, c, u)

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
			panic("No config file found. Using defaults and environment variables")
		}
	}

	// workaround because viper does not treat env vars the same as other config
	for _, key := range conf.AllKeys() {
		val := conf.Get(key)
		conf.Set(key, val)
	}

	return conf
}
