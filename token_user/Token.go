package token_user

import (
	"fmt"
	"log"
	"time"

	"github.com/ariefsam/go-chat/entity"
	"github.com/dgrijalva/jwt-go"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Token struct {
	Secret []byte
}
type MyCustomClaims struct {
	User entity.User
	jwt.StandardClaims
}

func (t *Token) Create(user entity.User) (tokenString string) {

	// Create the Claims
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(),
			Issuer:    "go-chat",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(t.Secret)
	fmt.Printf("%v %v", tokenString, err)
	return
}
func (t *Token) Parse(tokenString string) (isValid bool, user entity.User) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return t.Secret, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		isValid = true
		user = claims.User
	} else {
		log.Println(err)
	}
	return

}
