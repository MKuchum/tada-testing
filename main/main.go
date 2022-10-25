package main

import (
	"github.com/MKuchum/tada-testing/internal/web"
	"net/http"
)

func main() {
	s := web.NewServer()
	if err := http.ListenAndServe("localhost:8080", s); err != nil {
		panic(err)
	}
}
