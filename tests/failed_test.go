package tests

import (
	"github.com/MKuchum/tada-testing/models"
	"net/http"
	"testing"
)

func TestFailed(t *testing.T) {
	_, code, _ := doReq(t, nil)
	assertEqualsInt(t, http.StatusBadRequest, code, "nil body")
	_, code, _ = doReq(t, &models.TribonacciInput{})
	assertEqualsInt(t, http.StatusBadRequest, code, "empty body")
	_, code, _ = doReq(t, &models.TribonacciInput{N: 0})
	assertEqualsInt(t, http.StatusBadRequest, code, "n = 0")
	_, code, _ = doReq(t, &models.TribonacciInput{Signature: []float32{1, 1, 1}})
	assertEqualsInt(t, http.StatusBadRequest, code, "without n")
	_, code, _ = doReq(t, &models.TribonacciInput{Signature: []float32{1}, N: 10})
	assertEqualsInt(t, http.StatusBadRequest, code, "invalid signature")
}
