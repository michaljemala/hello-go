package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	HostVar = "VCAP_APP_HOST"
	PortVar = "VCAP_APP_PORT"
)

func main() {
	http.HandleFunc("/", hello)
	var port string
	if port = os.Getenv(PortVar); port == "" {
		port = "8080"
	}
	fmt.Printf("Listening at port %v\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, world\n\n")
	env := os.Environ()
	for _, e := range env {
		fmt.Fprintln(res, e)
	}
}
