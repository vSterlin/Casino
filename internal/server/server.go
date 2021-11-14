package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr int, db *sql.DB) *Server {
	strAddr := strconv.Itoa(addr)
	return &Server{addr: strAddr, db: db}
}

func (s *Server) Init() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/games", func(r chi.Router) {
		r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("test"))
		})
	})

	fmt.Printf("Listening on port %s\n", s.addr)
	http.ListenAndServe(":"+s.addr, r)
}

func (s *Server) Shutdown() {
	s.db.Close()
}
