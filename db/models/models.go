// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"database/sql"
)

type Session struct {
	ID        string
	UserID    string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
