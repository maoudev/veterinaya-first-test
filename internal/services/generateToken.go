package services

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
		return
	}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	publicBytes, err := os.ReadFile("./public.pem")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}

func GenerateToken(user models.User) string {
	claims := models.Claims{
		User: models.User{
			Id:       user.Id,
			Rut:      user.Rut,
			Email:    user.Email,
			UserName: user.UserName,
			Roles:    user.Roles,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			Issuer:    user.Id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
	result, err := token.SignedString(PrivateKey)
	if err != nil {
		log.Fatal("Error al registrar token")
	}

	return result
}
