package main

import (
	"net/http"
	"go-to-do/api/tasker"
)

func main() {
	r := tasker.NewRouter()

	if err := http.ListenAndServe(":8888", r); err != nil {
		panic(err)
	}
}