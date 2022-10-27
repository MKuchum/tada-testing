package web

import (
	"github.com/MKuchum/tada-testing/internal/tribonacci"
	"net/http"
)

type Server struct {
	*http.ServeMux
	t *tribonacci.Tribonacci
}

func NewServer() *Server {
	s := &Server{
		ServeMux: http.NewServeMux(),
		t:        tribonacci.NewTribonacci(),
	}
	s.HandleFunc(getTribonacciPath, s.GetTribonacciHandler)
	return s
}
