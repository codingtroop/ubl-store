package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	_ "github.com/codingtroop/ubl-store/docs"
	api "github.com/codingtroop/ubl-store/pkg/handlers"
	"github.com/codingtroop/ubl-store/pkg/repositories/sqlite"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	echoSwagger "github.com/swaggo/echo-swagger"
)

func createTable(db *sql.DB) {
	ublSql := `create table if not exists ubl (
		"ID" GUID,		
		"Created" TEXT
	  );`

	us, err := db.Prepare(ublSql)
	if err != nil {
		log.Fatal(err.Error())
	}
	us.Exec()

	aSql := `create table if not exists attachment (
		"ID" GUID,		
		"Created" TEXT,
		"UblID" GUID,
		"Hash" TEXT
	  );`

	as, err := db.Prepare(aSql)
	if err != nil {
		log.Fatal(err.Error())
	}
	as.Exec()
}

func main() {
	e := echo.New()

	hc := api.NewHealthCheckHandler()

	dbPath := "sqlite-database.db"

	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {

		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()

		log.Println("sqlite-database.db created")
	}

	db, _ := sql.Open("sqlite3", "./"+dbPath)

	createTable(db)

	defer db.Close()

	ur := sqlite.NewSqliteUblRepository(db)
	ar := sqlite.NewSqliteAttanchmentRepository(db)

	uh := api.NewUblStoreHandler(ur, ar)

	e.GET("/health", hc.Live)
	e.GET("/health/live", hc.Live)
	e.GET("/health/ready", hc.Ready)

	ug := e.Group("/api/v1/ubl")

	ug.GET("", uh.Get)
	ug.POST("", uh.Post)
	ug.DELETE("", uh.Delete)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World haha oldu!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
