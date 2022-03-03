package user

import (
	"context"
	"time"

	"github.com/tapiaw38/pottery-api/database"
	"github.com/tapiaw38/pottery-api/models"
)

// UpdateUser updates a user in the database
func UpdateUser(ctx context.Context, id string, u models.User) (models.User, error) {

	var user models.User

	q := `
	UPDATE users
		SET first_name = $1, last_name = $2, username = $3, email = $4, picture = $5, updated_at = $6
		WHERE id = $7
		RETURNING id, first_name, last_name, username, email, picture, created_at, updated_at;
	`

	rows := database.Data().QueryRowContext(
		ctx, q, u.FirstName, u.LastName, u.Username, u.Email,
		u.Picture, time.Now(), id,
	)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Picture,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil

}
