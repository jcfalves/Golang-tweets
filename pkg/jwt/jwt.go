package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID int64, username string, secretKey string) (string, error) {
	// Implementation for creating a JWT token
	// You can use a JWT library to generate the token with the user ID and username as claims
	// Return the generated token and any error encounter

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       userID,
			"username": username,
			"exp":      time.Now().Add(60 * time.Minute).Unix(), // Token expiration time
		},
	)

	key := []byte(secretKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secretKey string, withClaimValidation bool) (int64, string, error) {

	var (
		key    = []byte(secretKey)
		claims = jwt.MapClaims{}
		token  *jwt.Token
		err    error
	)

	if withClaimValidation {
		token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	} else {
		token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithoutClaimsValidation())
	}

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	userID := int64(claims["id"].(float64))
	username := claims["username"].(string)

	return userID, username, nil
}
