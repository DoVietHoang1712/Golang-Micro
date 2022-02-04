package utils

import (
	"Golang-Microservice/src/microservice-movie/common"
	"errors"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SdtClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt_lib.StandardClaims
}

type Util struct {
}

func (u *Util) GenerateJWT(name, role string) (string, error) {
	claims := SdtClaims{
		name,
		role,
		jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Issuer:    common.Config.Issuer,
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(common.Config.JwtSecretPassword))
	return tokenString, err
}

func (u *Util) ValidateObjectID(id string) error {
	if !primitive.IsValidObjectID(id) {
		return errors.New(common.ErrNotObjectIDHex)
	}

	return nil
}
