package sqlite

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library

	"github.com/codingtroop/ubl-store/pkg/entities"
	"github.com/codingtroop/ubl-store/pkg/repositories/interfaces"
	"github.com/google/uuid"
)

type sqliteAttachmentRepository struct {
	db *sql.DB
}

func NewSqliteAttanchmentRepository(db *sql.DB) interfaces.AttachmentRepository {
	return &sqliteAttachmentRepository{db: db}
}

func (r *sqliteAttachmentRepository) Get(cntxt context.Context, id uuid.UUID) (*entities.AttachmentEntity, error) {
	q := "SELECT * FROM attachment where ID = ?"

	ps, err := r.db.Prepare(q)

	if err != nil {
		return nil, err
	}

	qs := ps.QueryRow(id)

	if err != nil {
		return nil, err
	}

	entity := entities.AttachmentEntity{}

	if err := qs.Scan(&entity.ID, &entity.Created, &entity.UblID, &entity.Hash); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &entity, nil
}

func (r *sqliteAttachmentRepository) Insert(cntxt context.Context, e entities.AttachmentEntity) error {
	q := "INSERT INTO attachment(ID, Created, UblID, Hash) VALUES(?, ?, ?, ?)"

	ps, err := r.db.Prepare(q)

	if err != nil {
		return err
	}

	if _, err := ps.Exec(e.ID, e.Created, e.UblID, e.Hash); err != nil {
		return err
	}

	return nil
}

func (r *sqliteAttachmentRepository) Delete(cntxt context.Context, id uuid.UUID) error {
	q := "DELETE FROM attachment WHERE ID = ?"

	ps, err := r.db.Prepare(q)

	if err != nil {
		return err
	}

	if _, err := ps.Exec(id); err != nil {
		return err
	}

	return nil
}