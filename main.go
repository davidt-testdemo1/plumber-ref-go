// Minimal HTTP service used as the Plumber campaign's container control.
//
// Deliberately dependency-free and stdlib-only: the point of a control is that
// when a deploy fails, the cause is Plumber and not this program. It honours
// $PORT because the generated Dockerfile sets it, and answers 200 at "/"
// because the generated target group health-checks that path.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	started := time.Now()
	mux := http.NewServeMux()

	payload := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"service": "plumber-ref-go",
			"status":  "ok",
			"uptimeSeconds": int(time.Since(started).Seconds()),
			"host":    r.Host,
			"path":    r.URL.Path,
		})
	}

	mux.HandleFunc("/", payload)
	mux.HandleFunc("/healthz", payload)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("plumber-ref-go listening on :%s", port)
	log.Fatal(srv.ListenAndServe())
}
