package main

import (
	"flag"
	"fmt"
	"github.com/MKuchum/tada-testing/internal/web"
	"net/http"
)

func main() {
	var redisAddr string
	var redisPassword string
	var redisDB int
	var serverAddr string
	flag.StringVar(&redisAddr, "redis_addr", "localhost:6379", "Redis address in format 'host:port', by default 'localhost:6739'")
	flag.StringVar(&redisPassword, "redis_password", "", "Redis password, by default empty")
	flag.IntVar(&redisDB, "redis_db", 0, "Redis database, by default 0 (redis default DB)")
	flag.StringVar(&serverAddr, "addr", "localhost:8080", "Server address in format 'host:port', by default 'localhost:8080'")
	flag.Parse()

	s, err := web.NewServer(redisAddr, redisPassword, redisDB)
	if err != nil {
		panic(fmt.Errorf("can not create server: %v", err))
	}
	if err := http.ListenAndServe(serverAddr, s); err != nil {
		panic(fmt.Errorf("internal error while serve: %v", err))
	}
}
