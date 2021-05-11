package gosayhello

import (
	"encoding/json"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/dimasadiandrea/go-say-hello/model"
)

var SECRET_KEY = []byte("asdf")

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}

func DecodeToken(encodedToken string) (decodedResult model.DecodedStructure, errData error) {
	tokenString := encodedToken
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return decodedResult, err
	}
	if !token.Valid {
		return decodedResult, errors.New("invalid token")
	}

	// // do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

	jsonbody, err := json.Marshal(claims)
	if err != nil {
		return decodedResult, err
	}

	var obj model.DecodedStructure
	if err := json.Unmarshal(jsonbody, &obj); err != nil {
		return decodedResult, err
	}

	return obj, nil
}
