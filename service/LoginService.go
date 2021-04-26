package service

import (
	"authentication/common"
	"authentication/config"
	"authentication/model"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

type authCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	claims := &authCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return t
}

func  ValidateToken(encodedToken string) (*jwt.Token, error) {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(secret), nil
	})

}

func Login(dto model.UserLoginDto) (bool,string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel();
	var  user model.User

	if err := config.UserCollection.FindOne(ctx, bson.M{"username" : dto.Username}).Decode(&user); err != nil{
		msg := fmt.Println("Username or password incorrect!");
		return false, msg
	}
	
	if(!common.CheckPasswordHash(dto.Password, user.Password)){
		msg := fmt.Println("Username or password incorrect!");
		return false, msg
	}
	defer cancel();
	return true, GenerateToken(user.Username)
}


