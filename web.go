package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "OK")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
