package main

import (
	"code.google.com/p/log4go"
	"fmt"
	"net/http"
	"os"
)

const Port = "VCAP_APP_PORT"

func main() {
	log := make(log4go.Logger)
	log.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<h1>Listening on %s</h1>", os.Getenv(Port))
	})
	
	var p string
	if p = os.Getenv(Port); p == "" {
		p = "8080"
	}
	
	log.Debug("Listening at port %v\n", p)
	if err := http.ListenAndServe(":"+p, nil); err != nil {
		panic(err)
	}
}
