package security

import (
	"fmt"
	"go-api/model"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	JwtSecretKey = "JWT_SECRET_KEY"
)

func GenerateToken(user *model.User) (string, error) {
	secret := os.Getenv(JwtSecretKey)
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimPrefix("Bearer", token)
	}
	return token
}

func VerifyToken(token string) (user *model.User, err error) {
	secret := os.Getenv(JwtSecretKey)

	verifiedToken, err := jwt.Parse(RemoveBearerPrefix(token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := verifiedToken.Claims.(jwt.MapClaims)
	if !ok || !verifiedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return &model.User{
		ID: claims["id"].(string),
	}, nil
}
