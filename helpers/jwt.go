package helpers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenMetaData ...
type TokenMetaData struct {
	ID    string
	Email string
}

// CreateToken ...
func CreateToken(uid, email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["uid"] = uid
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func extractToken(ha string) string {
	token := strings.Split(ha, " ")
	if len(token) == 2 {
		return token[1]
	}
	return ""
}

func verifyToken(authHeader string) (*jwt.Token, error) {
	tokenString := extractToken(authHeader)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValidation ...
func TokenValidation(authHeader string) error {
	token, err := verifyToken(authHeader)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

// ExtractTokenMetaData ...
func ExtractTokenMetaData(authHeader string) (*TokenMetaData, error) {
	token, err := verifyToken(authHeader)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, ok := claims["uid"].(string)
		if !ok {
			return nil, err
		}
		email, ok := claims["email"].(string)
		if !ok {
			return nil, err
		}
		return &TokenMetaData{
			ID:    uid,
			Email: email,
		}, nil
	}
	return nil, err
}
