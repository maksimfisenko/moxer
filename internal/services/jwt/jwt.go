package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secret = "70fff2fdf9be90828790a6429c92d5fa67521dfbe031e63f6cd2fca94a8deb5621db6a5d12d941729c0c47162831aa57"
)

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims format")
	}

	userId, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("user id not found")
	}

	return userId, nil
}
