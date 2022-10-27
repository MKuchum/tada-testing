package main

import (
	"fmt"
	"github.com/MKuchum/tada-testing/internal/web"
	"net/http"
)

func main() {
	s, err := web.NewServer("localhost:6379", "", 0)
	if err != nil {
		panic(fmt.Errorf("can not create server: %v", err))
	}
	if err := http.ListenAndServe("localhost:8080", s); err != nil {
		panic(fmt.Errorf("internal error while serve: %v", err))
	}
}
