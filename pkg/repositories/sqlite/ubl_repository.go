package sqlite

import (
	"context"
	"database/sql"

	"github.com/codingtroop/ubl-store/pkg/entities"
	"github.com/codingtroop/ubl-store/pkg/repositories/interfaces"
	"github.com/google/uuid"
)

type sqliteUblRepository struct {
	db *sql.DB
}

func NewSqliteUblRepository(db *sql.DB) interfaces.UblRepository {
	return &sqliteUblRepository{db: db}
}

func (r *sqliteUblRepository) Get(cntxt context.Context, id uuid.UUID) (*entities.UblEntity, error) {
	q := "SELECT * FROM ubl where ID = ?"

	ps, err := r.db.Prepare(q)

	if err != nil {
		return nil, err
	}

	qs := ps.QueryRow(id)

	if err != nil {
		return nil, err
	}

	entity := entities.UblEntity{}

	if err := qs.Scan(&entity.ID, &entity.Created); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &entity, nil
}

func (r *sqliteUblRepository) Insert(cntxt context.Context, e entities.UblEntity) error {
	q := "INSERT INTO ubl(ID, Created) VALUES(?, ?)"

	ps, err := r.db.Prepare(q)

	if err != nil {
		return err
	}

	if _, err := ps.Exec(e.ID, e.Created); err != nil {
		return err
	}

	return nil
}

func (r *sqliteUblRepository) Delete(cntxt context.Context, id uuid.UUID) error {
	q := "DELETE FROM ubl WHERE ID = ?"

	ps, err := r.db.Prepare(q)

	if err != nil {
		return err
	}

	if _, err := ps.Exec(id); err != nil {
		return err
	}

	return nil
}
