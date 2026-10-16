package server

import (
	"database/sql"
	"net/http"
)

type Server struct {
	db *sql.DB
}

func New(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// mux.Handle()

	return mux
}

func (s *Server) healthCheck(w http.ResponseWriter, r http.Request){
	

}
