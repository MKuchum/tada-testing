package tests

import (
	"bytes"
	"encoding/json"
	"github.com/MKuchum/tada-testing/internal/web"
	"github.com/MKuchum/tada-testing/models"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

var ts = []*Test{
	{Signature: []float32{0, 0, 0}, N: 10, Sequence: []float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	{Signature: []float32{0, 0, 1}, N: 1, Sequence: []float32{0}},
	{Signature: []float32{0, 0, 1}, N: 2, Sequence: []float32{0, 0}},
	{Signature: []float32{0, 0, 1}, N: 3, Sequence: []float32{0, 0, 1}},
	{Signature: []float32{0, 0, 1}, N: 10, Sequence: []float32{0, 0, 1, 1, 2, 4, 7, 13, 24, 44}},
}

func createServer() {
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
		serverAddr = "localhost:8082"
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
	go func() {
		if err := http.ListenAndServe(serverAddr, s); err != nil {
			logger.Fatal("internal error while serve", zap.Error(err))
		}
	}()
}

func TestSuccessfully(t *testing.T) {
	createServer()
	for i, test := range ts {
		input := &models.TribonacciInput{Signature: test.Signature, N: test.N}
		b, err := json.Marshal(input)
		assertError(t, err)
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8082/tri", bytes.NewReader(b))
		assertError(t, err)
		resp, err := http.DefaultClient.Do(req)
		assertError(t, err)
		respBody, err := io.ReadAll(resp.Body)
		assertError(t, err)
		output := &models.TribonacciOutput{}
		err = json.Unmarshal(respBody, output)
		assertError(t, err)
		assertEqualsTribonacci(t, test.Sequence, output.Sequence, "test "+strconv.Itoa(i))
	}
}
