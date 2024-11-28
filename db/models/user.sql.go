// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package models

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO user (id, name, email, password)
VALUES (?, ?, ?, ?)
RETURNING id, name, email
`

type CreateUserParams struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type CreateUserRow struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
	)
	var i CreateUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, created_at, updated_at
FROM user
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, email, password, created_at, updated_at
FROM user
WHERE id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, name, email, password, created_at, updated_at
FROM user
ORDER BY name
`

func (q *Queries) ListUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const login = `-- name: Login :one
SELECT id, email, name
FROM user
WHERE email = ? AND password = ?
`

type LoginParams struct {
	Email    string
	Password string
}

type LoginRow struct {
	ID    string
	Email string
	Name  string
}

func (q *Queries) Login(ctx context.Context, arg LoginParams) (LoginRow, error) {
	row := q.db.QueryRowContext(ctx, login, arg.Email, arg.Password)
	var i LoginRow
	err := row.Scan(&i.ID, &i.Email, &i.Name)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE user
SET name = ?, email = ?, password = ?
WHERE id = ?
RETURNING id, name, email
`

type UpdateUserParams struct {
	Name     string
	Email    string
	Password string
	ID       string
}

type UpdateUserRow struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.ID,
	)
	var i UpdateUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}
