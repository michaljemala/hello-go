package main

import (
	"code.google.com/p/log4go"
	"fmt"
	"launchpad.net/goyaml"
	"net/http"
	"os"
	"runtime"
)

const (
	HostVar = "VCAP_APP_HOST"
	PortVar = "VCAP_APP_PORT"
)

type T struct {
	A string
	B []int
}

func main() {
	log := make(log4go.Logger)
	log.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())

	http.HandleFunc("/", hello)
	var port string
	if port = os.Getenv(PortVar); port == "" {
		port = "8080"
	}
	log.Debug("Listening at port %v\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	// Dump Go version
	fmt.Fprintf(res, "%v\n\n", runtime.Version())

	// Dump ENV
	env := os.Environ()
	for _, e := range env {
		fmt.Fprintln(res, e)
	}

	//Dump some YAML
	t := T{A: "Foo", B: []int{1, 2, 3}}
	if d, err := goyaml.Marshal(&t); err != nil {
		fmt.Fprintf(res, "Unable to dump YAML")
	} else {
		fmt.Fprintf(res, "\n\n--- \n%s", d)
	}
}
