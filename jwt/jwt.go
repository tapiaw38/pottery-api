package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tapiaw38/pottery-api/models"
)

// GenerateToken generates a JWT token
func GenerateJWT(user models.User) (string, error) {

	myKey := []byte("bXlTZWNyZXRQYXNzd29yZG15U2VjcmV0UGFzc3dvcmQK")

	payload := jwt.MapClaims{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"username":   user.Username,
		"picture":    user.Picture,
		"is_active":  user.IsActive,
		"is_admin":   user.IsAdmin,
		"exp":        user.CreatedAt.Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
