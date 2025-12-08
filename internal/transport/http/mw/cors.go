package mw

import "net/http"

func CORS(new http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")
		new.ServeHTTP(w, r)
	})
}
