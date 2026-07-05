package middleware

import "net/http"

func PreventCORF(next http.Handler) http.Handler {
	cop := http.NewCrossOriginProtection()

	cop.SetDenyHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "CSRF check failed", http.StatusBadRequest)
	}))

	return cop.Handler(next)
}
