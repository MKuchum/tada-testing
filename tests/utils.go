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

func init() {
	createServer()
}

func assertError(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Error(err)
}

func assertEqualsTribonacci(t *testing.T, expected []float32, real []float32, comment ...string) {
	c := ""
	if len(comment) > 0 {
		c = comment[0]
	}
	if len(expected) != len(real) {
		t.Errorf("%s, expected = %v, real = %v", c, expected, real)
	}
	for i := range expected {
		if expected[i] != real[i] {
			t.Errorf("%s, expected = %v, real = %v", c, expected, real)
		}
	}
}

func assertEqualsInt(t *testing.T, expected int, real int, comment ...string) {
	c := ""
	if len(comment) > 0 {
		c = comment[0]
	}
	if expected != real {
		t.Errorf("%s, expected = %d, real = %d", c, expected, real)
	}
}

type Test struct {
	Signature []float32
	N         int
	Sequence  []float32
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

func doReq(t *testing.T, input *models.TribonacciInput) (*models.TribonacciOutput, int, error) {
	b, err := json.Marshal(input)
	assertError(t, err)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8082/tri", bytes.NewReader(b))
	assertError(t, err)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	respBody, err := io.ReadAll(resp.Body)
	assertError(t, err)
	output := &models.TribonacciOutput{}
	err = json.Unmarshal(respBody, output)
	return output, resp.StatusCode, nil
}
