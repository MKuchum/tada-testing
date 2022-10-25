package web

import "net/http"

type Server struct {
	*http.ServeMux
}

func NewServer() *Server {
	s := &Server{ServeMux: http.NewServeMux()}
	s.HandleFunc(getTribonacciPath, s.GetTribonacciHandler)
	return s
}
