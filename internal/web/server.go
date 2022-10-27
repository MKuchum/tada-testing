package web

import (
	"github.com/MKuchum/tada-testing/internal/tribonacci"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	*http.ServeMux
	t      *tribonacci.Tribonacci
	logger *zap.Logger
}

func NewServer(redisAddr string, redisPassword string, redisDB int, logger *zap.Logger) (*Server, error) {
	t, err := tribonacci.NewTribonacciRedis(redisAddr, redisPassword, redisDB, logger)
	if err != nil {
		return nil, err
	}
	s := &Server{
		ServeMux: http.NewServeMux(),
		t:        t,
		logger:   logger,
	}
	s.HandleFunc(getTribonacciPath, s.GetTribonacciHandler)
	return s, nil
}
