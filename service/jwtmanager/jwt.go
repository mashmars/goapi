package jwtmanager

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"errors"
)


func CreateJwtToken(claims map[string]interface{}) (tokenString string, err error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": claims["username"],
	})

	tokenString, err = token.SignedString([]byte("jwt security key string"))
	return
}


func ParseJwtToken(tokenString string) (claims map[string]interface{}, err error) {
	if tokens := strings.Split(tokenString, "."); len(tokens) != 3 {
		return map[string]interface{}{}, errors.New("token格式不正确")
	}
	
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