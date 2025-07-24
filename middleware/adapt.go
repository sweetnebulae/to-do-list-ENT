package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Adapt(fn httprouter.Handle) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		fn(w, r, params)
	})
}
