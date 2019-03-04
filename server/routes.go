package server

func (s *Server) Routes() {
	s.Router.HandleFunc("/", s.indexHandler())
}
