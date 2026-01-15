package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"web-song/assets"
)

var (
	//go:embed index.html
	indexHtml  string
	serverPort string
	mediaDir   string
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	mediaDir = "/app/media"
	serverPort = os.Getenv("SERVE_PORT")
	if serverPort == "" {
		serverPort = "7880"
	}

	addrFlag := flag.String("addr", ":"+serverPort, "Listen adress")
	flag.Parse()

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(mediaDir))

	fileServer := http.FileServer(assets.FS)
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fileServer))

	mux.Handle("/media/", http.StripPrefix("/media/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "audio/mpeg")
		fs.ServeHTTP(w, r)
	})))

	mux.Handle("/", http.HandlerFunc(indexHandler))

	srv := &http.Server{
		Addr:              *addrFlag,
		Handler:           securityHeaders(mux),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("Serving on %s", *addrFlag)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe error: %v", err)
	}

}

func securityHeaders(next http.Handler) http.Handler {
	const csp = "default-src 'self'; script-src 'self' 'unsafe-inline'; img-src 'self' data:; style-src 'self' 'unsafe-inline'"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Content-Security-Policy", csp)
		next.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.New("index").Parse(indexHtml)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	path := r.URL.Path
	files, err := loadFiles(mediaDir)
	if err != nil {
		http.Error(w, "load files fail", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, struct {
		Files template.JS
		Path  string
	}{Path: path, Files: template.JS(mustJson(files))})

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func loadFiles(filesDir string) ([]string, error) {
	entries, err := os.ReadDir(filesDir)
	if err != nil {
		return nil, errors.New("erro ao carregar as musicas")
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, "/media/"+e.Name())
		}
	}
	return files, nil
}

func mustJson(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
