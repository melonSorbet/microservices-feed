package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte("6f23c4c24227b0d5039a56b7c29b44372bce67c54bdcdc3d879778ceff01d58d8b9a2f0fc9370812626e1bc70a29bc36e23c0626fb0bcc4146203b6d8d3416c9038bfc20ffaad9d21072b763cfac161d85a295c53830b62a110b841b88efb5f7816426f30623a5bec501c35d34e06acb090f54ea37decf3c3c4946970a4048e3")

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
