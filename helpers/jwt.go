package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mhdiiilham/gournal/helpers"
	log "github.com/sirupsen/logrus"
)

// TokenMetaData ...
type TokenMetaData struct {
	ID    uint64
	Email string
}

// CreateToken ...
func CreateToken(uid uint64, email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["uid"] = uid
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(helpers.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func extractToken(ha string) string {
	return ha
}

func verifyToken(authHeader string) (*jwt.Token, error) {
	tokenString := extractToken(authHeader)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(helpers.Getenv("JWT_SECRET")), nil
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
		uid, err := strconv.ParseUint(fmt.Sprintf("%v", claims["uid"]), 10, 64)
		if err != nil {
			log.Info(80, err)
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
