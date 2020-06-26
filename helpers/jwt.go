package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken ...
func CreateToken(uid, email string) (string, error) {
	atClaims          := jwt.MapClaims{}
	atClaims["uid"]   = uid
	atClaims["email"] = email
	atClaims["exp"]   = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodES256, atClaims)
	return at.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
