package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

type Options struct {
	port           uint
	url            string
	scriptLocation string
	verbose        bool
}

// @todo parse flags for real using flag package
func getOptions() Options {
	port := flag.Uint("p", 8000, "port number to listen on")
	url := flag.String("u", "/", "url to listen on")
	verbose := flag.Bool("v", false, "enable verbose output")
	scriptLocation := flag.String("s", "./test.sh", "location of script to call when hook triggered")

	flag.Parse()

	option := Options{
		port:           *port,
		url:            *url,
		scriptLocation: *scriptLocation,
		verbose:        *verbose,
	}

	return option
}

func createHttpServer(endpoint string, port uint, onRequest func()) {
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		onRequest()
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":"+strconv.Itoa(int(port)), nil)
}

// @todo execute script for real
// maybe want rate limiter too, but idk for now
func getCallExternalScript(scriptLocation string) func() {
	return func() {
		fmt.Println("call external placeholder")
	}
}

func main() {
	option := getOptions()

	if option.verbose {
		fmt.Println("port is: ", option.port)
		fmt.Println("url is: ", option.url)
		fmt.Println("script is: ", option.scriptLocation)
	}

	createHttpServer(option.url, option.port, getCallExternalScript(option.scriptLocation))
}
