package server

import (
	"net/http"
	"fmt"
)

func (s *Server) indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from minicube app - updated 05/03/2018")
	}
}
