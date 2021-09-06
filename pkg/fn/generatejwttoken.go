package fn

import (
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"log"
	"time"
)

func GenerateJWTToken(role string, isKeyFile bool, secret string) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["Role"] = role

	if isKeyFile {
		content, err := ioutil.ReadFile(secret)
		if err != nil {
			log.Fatal(err)
		}
		key, err := jwt.ParseRSAPrivateKeyFromPEM(content)
		if err != nil {
			log.Fatal(err)
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

		if err != nil {
			log.Fatal(err)
			return "", err
		} else {
			return token, nil
		}

	} else {

		jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := jwt.SignedString([]byte(secret))
		if err != nil {
			log.Fatal(err)
			return "", err
		} else {
			return token, nil
		}
	}
}
