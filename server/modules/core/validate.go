package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body T
		err := decoder.Decode(&body)
		if err != nil {
			invalidInputError(w, []string{err.Error()})
			return
		}

		validate := validator.New()
		err = validate.Struct(body)
		if err != nil {
			errMsgs := getValidationErrors(err)
			invalidInputError(w, errMsgs)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getValidationErrors(err error) []string {
	errMsgs := make([]string, 0)

	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
	}

	for _, err := range err.(validator.ValidationErrors) {

		errMsgs = append(errMsgs, fmt.Sprintf(
			"%s : Needs to implement '%s'",
			err.Field(),
			err.Tag(),
		))
	}

	return errMsgs
}

func invalidInputError(w http.ResponseWriter, errMsgs []string) {
	w.WriteHeader(400)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fiber.Map{
		"success": false,
		"result":  errMsgs,
	})
}
