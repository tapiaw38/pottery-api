package database

import (
	"context"

	"github.com/tapiaw38/pottery-api/models"
)

// Get all users from database
func GetUsers(ctx context.Context) ([]models.User, error) {

	q := `
	SELECT id, first_name, last_name, username, email, picture, password, created_at, updated_at
		FROM users;
	`

	rows, err := data.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var u models.User

		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username, &u.Email, &u.Picture, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil

}

// Get user by id from database
func GetUserById(ctx context.Context, id string) (*models.User, error) {

	q := `
	SELECT id, first_name, last_name, username, email, picture, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	row := data.QueryRowContext(
		ctx, q, id,
	)

	u := models.User{}

	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username, &u.Email, &u.Picture, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil

}
