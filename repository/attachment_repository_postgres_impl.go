package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ryuuzake/todolist-gofiber/config"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type AttachmentRepositoryPostgresImpl struct {
	*sqlx.DB
}

func NewAttachmentRepositoryPostgresImpl() *AttachmentRepositoryPostgresImpl {
	db, err := config.OpenDBConnection()
	if err != nil {
		panic(err.Error())
	}

	return &AttachmentRepositoryPostgresImpl{DB: db}
}

func (repo *AttachmentRepositoryPostgresImpl) FindAllFromTodolistId(id uuid.UUID) ([]model.Attachment, error) {
	attachments := make([]model.Attachment, 0)

	query := `SELECT * FROM attachments WHERE todolist_id=$1`

	err := repo.Select(&attachments, query, id)
	if err != nil {
		return attachments, err
	}

	return attachments, nil
}

func (repo *AttachmentRepositoryPostgresImpl) FindById(id uuid.UUID) (model.Attachment, error) {
	var attachment model.Attachment

	query := `SELECT * FROM attachments WHERE id=$1 LIMIT 1`

	err := repo.Get(&attachment, query, id)
	if err != nil {
		return attachment, err
	}

	return attachment, nil
}

func (repo *AttachmentRepositoryPostgresImpl) Create(attachment model.Attachment) error {
	query := `INSERT INTO attachments (todolist_id, url, caption)
		VALUES(:todolist_id, :url, :caption)`

	_, err := repo.NamedExec(query, attachment)
	if err != nil {
		return err
	}

	return nil
}

func (repo *AttachmentRepositoryPostgresImpl) UpdateById(id uuid.UUID, attachment model.Attachment) error {
	// Set Id to Id from params
	attachment.Id = id

	var query string
	if attachment.Url != "" {
		query = `UPDATE attachments SET todolist_id=:todolist_id, url=:url, caption=:caption WHERE id=:id`
	} else {
		query = `UPDATE attachments SET todolist_id=:todolist_id, caption=:caption WHERE id=:id`
	}

	_, err := repo.NamedExec(query, attachment)
	if err != nil {
		return err
	}

	return nil
}

func (repo *AttachmentRepositoryPostgresImpl) DeleteById(id uuid.UUID) error {
	query := `DELETE FROM attachments WHERE id=$1`

	_, err := repo.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
