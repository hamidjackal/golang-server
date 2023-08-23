package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			unauthorizedError(w)
			return
		} else {
			splittedToken := strings.Split(token, "Bearer ")
			if splittedToken[1] != "" {
				id, err := decodeToken(splittedToken[1])
				if err != nil {
					unauthorizedError(w)
					return
				}
				r.Header.Set("Userid", id)
			} else {
				unauthorizedError(w)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func unauthorizedError(w http.ResponseWriter) {
	w.WriteHeader(401)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fiber.Map{
		"success": false,
		"result":  "Unauthorized",
	})
}

func decodeToken(token string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	var id string

	for key, val := range claims {
		if key == "id" {
			id = fmt.Sprint(val)
		}
	}

	return id, nil
}
