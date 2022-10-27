package web

import (
	"github.com/MKuchum/tada-testing/internal/tribonacci"
	"net/http"
)

type Server struct {
	*http.ServeMux
	t *tribonacci.Tribonacci
}

func NewServer(redisAddr string, redisPassword string, redisDB int) (*Server, error) {
	t, err := tribonacci.NewTribonacciRedis(redisAddr, redisPassword, redisDB)
	if err != nil {
		return nil, err
	}
	s := &Server{
		ServeMux: http.NewServeMux(),
		t:        t,
	}
	s.HandleFunc(getTribonacciPath, s.GetTribonacciHandler)
	return s, nil
}
