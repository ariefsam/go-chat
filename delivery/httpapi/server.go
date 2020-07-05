package httpapi

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/api/login", LoginHandler).Methods("POST", "GET")
	r.HandleFunc("/api/login/verify", VerifyLoginHandler).Methods("POST", "GET")
	r.HandleFunc("/api/channel/create", CreateChannelHandler).Methods("POST", "GET")
	r.HandleFunc("/api/channel/search", SearchChannelHandler).Methods("POST", "GET")
	r.HandleFunc("/api/channel/detail", DetailChannelHandler).Methods("POST", "GET")
	r.HandleFunc("/api/channel/chat/create", ChannelCreateChatHandler).Methods("POST", "GET")

	r.PathPrefix("/").Handler(quasarHandler())

	s := &http.Server{
		Addr:           ":8889",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func quasarHandler() http.Handler {
	var dir string
	dir = "./quasar/dist/spa"
	return http.StripPrefix("/", http.FileServer(http.Dir(dir)))
}
