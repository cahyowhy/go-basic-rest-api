package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(payloads jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloads)

	return token.SignedString([]byte(os.Getenv("SECRETKEY")))
}

func DecodedToken(req *http.Request) []byte {
	if !ValidTokenHeader(req) {
		return generateJson(nil)
	}

	return nil
}

func ValidTokenHeader(req *http.Request) bool {
	tokenString := req.Header.Get("Authorization")

	return ValidToken(tokenString)
}

func ValidToken(tokenString string) bool {
	if !(len(tokenString) > 0) {
		return false
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		return []byte(os.Getenv("SECRETKEY")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		expired, okExp := claims["expired"].(string)
		i, err := strconv.ParseInt(expired, 10, 64)

		return i > time.Now().Unix() && err == nil && okExp
	}

	return false
}

func GetTokenParsed(req *http.Request) (mapClaim jwt.MapClaims, ok bool) {
	tokenString := req.Header.Get("Authorization")

	if len(tokenString) == 0 {
		return nil, false
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		return []byte(os.Getenv("SECRETKEY")), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)

	if !token.Valid {
		return nil, false
	}

	return claims, ok
}

func generateJson(payload interface{}) []byte {
	response, _ := json.Marshal(payload)
	if payload == nil {
		response, _ = json.Marshal(map[string]string{"data": "UNAUTHORIZED", "status": TOKEN_NOT_VALID})
	}

	return response
}
