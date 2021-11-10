package main

import (
	"log"
	"net/http"

	_ "github.com/codingtroop/ubl-store/docs"
	api "github.com/codingtroop/ubl-store/pkg/handlers"
	"github.com/codingtroop/ubl-store/pkg/helpers"
	"github.com/codingtroop/ubl-store/pkg/repositories/sqlite"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	hc := api.NewHealthCheckHandler()

	dbPath := "./sqlite-database.db"
	sqliteConnector := sqlite.NewSqliteConnector(dbPath)
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

	uh := api.NewUblStoreHandler(ur, ar, as, us, c)

	e.GET("/health", hc.Live)
	e.GET("/health/live", hc.Live)
	e.GET("/health/ready", hc.Ready)

	ug := e.Group("/api/v1/ubl")

	ug.GET("/:id", uh.Get)
	ug.POST("", uh.Post)
	ug.DELETE("/:id", uh.Delete)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World haha oldu!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
