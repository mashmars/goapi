package jwtmanager

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)


func CreateJwtToken(claims map[string]interface{}) (tokenString string, err error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": claims["username"],
	})

	tokenString, err = token.SignedString([]byte("jwt security key string"))
	return
}


func ParseJwtTOken(tokenString string) (claims map[string]interface{}, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("jwt security key string"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
	
}