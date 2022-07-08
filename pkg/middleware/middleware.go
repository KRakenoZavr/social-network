package middleware

import (
	"fmt"
	"net/http"
)

func Check(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("asd")

		next(w, r)
	})
}
