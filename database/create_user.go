package database

import (
	"context"
	"time"

	"github.com/tapiaw38/pottery-api/models"
)

func CreateUser(ctx context.Context, u *models.User) error {

	q := `
    INSERT INTO users (first_name, last_name, username, email, picture, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id;
    `

	if err := u.HashPassword(); err != nil {
		return err
	}

	row := data.QueryRowContext(
		ctx, q, u.FirstName, u.LastName, u.Username, u.Email,
		u.Picture, u.PasswordHash, time.Now(), time.Now(),
	)

	err := row.Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil

}
