package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)
//by the way it doesnt matter i did this because it is a small project in a real project i would have never put the key in here lol
var secretKey = []byte("YOUR_KEY")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
