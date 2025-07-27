package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Adapt(h httprouter.Handle) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		h(w, r, params)
	})
}
