package model

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	Id             uuid.UUID      `db:"id" json:"id"`
	FullName       string         `db:"full_name" json:"full_name"`
	Email          string         `db:"email" json:"email"`
	Password       string         `db:"password" json:"password"`
	ForgetPassword sql.NullString `db:"forget_password" json:"forget_password"`
	RoleId         uuid.UUID      `db:"role_id" json:"role_id"`
	Role           Role           `json:"role"`
	CreatedAt      time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at" json:"updated_at"`
}

type Role struct {
	Id          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
}
