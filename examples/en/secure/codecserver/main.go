package main

import (
	"log"
	"net/http"

	"go.temporal.io/sdk/converter"

	"tcourse/secure"
)

func main() {
	key, err := secure.KeyFromEnv("TEMPORAL_ENC_KEY")
	if err != nil {
		log.Fatalln(err)
	}
	handler := converter.NewPayloadCodecHTTPHandler(secure.NewCodecs("key-2026-11", key)...)

	http.Handle("/", withCORS("http://localhost:8233", handler))
	log.Println("codec server listening on :8081")
	log.Fatalln(http.ListenAndServe(":8081", nil))
}

// withCORS lets the Web UI call the codec server directly from the browser.
func withCORS(origin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Namespace, Authorization")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
