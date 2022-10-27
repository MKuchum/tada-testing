package main

import (
	"fmt"
	"github.com/MKuchum/tada-testing/internal/web"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var redisAddr string
	var redisPassword string
	var redisDB int
	var serverAddr string
	if redisAddr = os.Getenv("REDIS_ADDR"); redisAddr == "" {
		fmt.Println("$REDIS_ADDR is null, use default value 'localhost:6379'")
		redisAddr = "localhost:6379"
	}
	if redisPassword = os.Getenv("REDIS_PASSWORD"); redisPassword == "" {
		fmt.Println("$REDIS_PASSWORD is null, use default empty value")
		redisPassword = ""
	}
	if redisDBStr := os.Getenv("REDIS_DB"); redisDBStr == "" {
		fmt.Println("$REDIS_DB is null, use default zero value")
		redisDB = 0
	} else {
		var err error
		redisDB, err = strconv.Atoi(redisDBStr)
		if err != nil {
			panic(fmt.Sprintf("$REDIS_DB is not int"))
		}
	}
	if serverAddr = os.Getenv("SERVER_ADDR"); serverAddr == "" {
		fmt.Println("$SERVER_ADDR is null, use default value 'localhost:8080'")
		serverAddr = "localhost:8080"
	}

	s, err := web.NewServer(redisAddr, redisPassword, redisDB)
	if err != nil {
		panic(fmt.Errorf("can not create server: %v", err))
	}
	if err := http.ListenAndServe(serverAddr, s); err != nil {
		panic(fmt.Errorf("internal error while serve: %v", err))
	}
}
