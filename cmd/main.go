package main

import (
	"github.com/MKuchum/tada-testing/internal/web"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger, %v", err.Error())
	}
	logger.Info("logger created")
	defer logger.Sync()

	var redisAddr string
	var redisPassword string
	var redisDB int
	var serverAddr string
	if redisAddr = os.Getenv("REDIS_ADDR"); redisAddr == "" {
		logger.Info("$REDIS_ADDR is null, use default value 'localhost:6379'")
		redisAddr = "localhost:6379"
	}
	if redisPassword = os.Getenv("REDIS_PASSWORD"); redisPassword == "" {
		logger.Info("$REDIS_PASSWORD is null, use default empty value")
		redisPassword = ""
	}
	if redisDBStr := os.Getenv("REDIS_DB"); redisDBStr == "" {
		logger.Info("$REDIS_DB is null, use default zero value")
		redisDB = 0
	} else {
		var err error
		redisDB, err = strconv.Atoi(redisDBStr)
		if err != nil {
			logger.Fatal("$REDIS_DB is not int", zap.String("redisDB", redisDBStr))
		}
	}
	if serverAddr = os.Getenv("SERVER_ADDR"); serverAddr == "" {
		logger.Info("$SERVER_ADDR is null, use default value 'localhost:8080'")
		serverAddr = "localhost:8080"
	}
	logger.Info("Environments",
		zap.String("redisAddr", redisAddr),
		zap.String("redisPassword", redisPassword),
		zap.Int("redisDB", redisDB),
		zap.String("serverAddr", serverAddr),
	)

	s, err := web.NewServer(redisAddr, redisPassword, redisDB, logger)
	if err != nil {
		logger.Fatal("can not create server", zap.Error(err))
	}
	if err := http.ListenAndServe(serverAddr, s); err != nil {
		logger.Fatal("internal error while serve", zap.Error(err))
	}
	logger.Info("App successfully finished!!!")
}
