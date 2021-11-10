package sqlite

import (
	"database/sql"
	"errors"
	"log"
	"os"
)

type sqliteConnection struct {
	dbPath string
}

type SqliteConnector interface {
	Connect() (*sql.DB, error)
	Init(db *sql.DB) error
}

func NewSqliteConnector(dbPath string) SqliteConnector {
	return &sqliteConnection{dbPath}
}

func (c *sqliteConnection) Connect() (*sql.DB, error) {
	if _, err := os.Stat(c.dbPath); errors.Is(err, os.ErrNotExist) {

		file, err := os.Create(c.dbPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()

		log.Println("sqlite-database.db created")
	}

	db, _ := sql.Open("sqlite3", c.dbPath)
	return db, nil
}

func (c *sqliteConnection) Init(db *sql.DB) error {
	ublSql := `create table if not exists ubl (
		"ID" GUID,		
		"Created" datetime
	  );`

	us, err := db.Prepare(ublSql)
	if err != nil {
		return os.ErrDeadlineExceeded
	}

	us.Exec()

	aSql := `create table if not exists attachment (
		"ID" GUID,		
		"Created" datetime,
		"UblID" GUID,
		"Hash" TEXT
	  );`

	as, err := db.Prepare(aSql)
	if err != nil {
		return err
	}

	as.Exec()

	return nil
}
