package gochi

import (
	"net/http"
)

func (r *Controller) authorized(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		authorized := true

		if !authorized {
			http.Error(w, "Not Authorized", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}