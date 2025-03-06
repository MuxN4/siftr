package main

import (
	"database/sql"

	"github.com/MuxN4/siftr/internal/db"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func databaseUserToUser(dbUser db.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}
