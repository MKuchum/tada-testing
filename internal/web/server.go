package web

import (
	"github.com/MKuchum/tada-testing/internal/tribonacci"
	"net/http"
)

type Server struct {
	*http.ServeMux
	t *tribonacci.Tribonacci
}

func NewServer(redisAddr string, redisPassword string, redisDB int) *Server {
	s := &Server{
		ServeMux: http.NewServeMux(),
		t:        tribonacci.NewTribonacciRedis(redisAddr, redisPassword, redisDB),
	}
	s.HandleFunc(getTribonacciPath, s.GetTribonacciHandler)
	return s
}
