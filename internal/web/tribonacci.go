package web

import (
	"encoding/json"
	"fmt"
	"github.com/MKuchum/tada-testing/models"
	"io"
	"net/http"
)

const getTribonacciPath = "/tri"

func (s *Server) GetTribonacciHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != getTribonacciPath {
		http.Error(w, fmt.Sprintf("unknown path %s", req.URL.Path), http.StatusNotFound)
		return
	}
	if req.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("can use only GET method, not %s", req.Method), http.StatusNotFound)
		return
	}

	reqBytes, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("can not read body: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	input := &models.TribonacciInput{}
	if err := json.Unmarshal(reqBytes, input); err != nil {
		http.Error(w, fmt.Sprintf("can not unmarshal request: %v", err.Error()), http.StatusBadRequest)
		return
	}
	if input.Signature == nil { //default signature
		input.Signature = []float32{1, 1, 1}
	}
	if err := input.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("can not validate input: %v", err.Error()), http.StatusBadRequest)
		return
	}

	output, err := s.t.Generate(input)
	if err != nil {
		http.Error(w, fmt.Sprintf("internal error: %v", err.Error()), http.StatusInternalServerError)
		return
	}
	respBytes, err := json.Marshal(output)
	if err != nil {
		http.Error(w, fmt.Sprintf("can not marshal response: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(respBytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
