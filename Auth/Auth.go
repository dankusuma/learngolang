package Auth

import (
	"encoding/json"
	"fmt"
	"learngolang/Constants"
	"learngolang/Models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(Constants.PrivateKey + Constants.PublicKey), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Models.Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(Models.Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Models.Exception{Message: "An authorization header is required"})
		}
	})
}

func GenerateToken(w http.ResponseWriter, req *http.Request) {
	var user Models.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})
	tokenString, error := token.SignedString([]byte(Constants.PrivateKey + Constants.PublicKey))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(Models.JwtToken{Token: tokenString})
}
