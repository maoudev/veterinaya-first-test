package authentication

import (
	"crypto/rsa"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/maoudev/veterinaya-go/internal/models"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func init() {
	privateBytes, err := os.ReadFile("./private.pem")
	if err != nil {
		log.Fatal(err.Error())
	}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	publicBytes, err := os.ReadFile("./public.pem")
	if err != nil {
		log.Fatal(err.Error())
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GenerateToken(user models.User) (string, error) {
	claim := models.Claims{
		User: models.User{
			Id:       user.Id,
			UserName: user.UserName,
			Roles:    user.Roles,
		},

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    user.Id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenString, err := token.SignedString(PrivateKey)

	return tokenString, err

}
