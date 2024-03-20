package server

import (
	"net/http"
)

func checkValidId(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := idFromPathValue(r)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		if id < 0 {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
