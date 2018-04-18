package webserver

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Config struct {
	Port string
}

type requestLogger struct {
	Handle http.Handler
	Logger *log.Logger
}

func Start(cfg Config) {
	log.Println("Initializing web server")
	l := log.New(os.Stdout, "[iskandar] ", 0)

	port := ":" + cfg.Port
	if len(cfg.Port) < 1 {
		port = ":" + os.Getenv("PORT")
	}

	r := httprouter.New()
	loadRouter(r)

	http.ListenAndServe(port, requestLogger{Handle: r, Logger: l})
}
func (rl requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// === for development only ===
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	// === end for development only ===
	rl.Handle.ServeHTTP(w, r)
	log.Printf("[iskandar] %s %s in %v", r.Method, r.URL.Path, time.Since(start))
}
