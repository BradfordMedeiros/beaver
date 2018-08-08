
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"flag"
)

type Options struct {
	port uint;
	url string;
	scriptLocation string;
}

// @todo parse flags for real using flag package
func getOptions () Options{
	port := flag.Uint("p", 8000, "port number to listen on")
	url := flag.String("u", "/", "url to listen on")

	flag.Parse()

	option := Options { 
		port: *port,
		url: *url,
		scriptLocation: "somefilesystempath",
	}

	return option
}

func createHttpServer(endpoint string, port uint, onRequest func()) {
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request){
		onRequest()
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":"+strconv.Itoa(int(port)), nil)
}

// @todo execute script for real
// maybe want rate limiter too, but idk for now
func getCallExternalScript(scriptLocation string) func(){
	return func(){
		fmt.Println("call external placeholder")
	}
}

func main() {
	option := getOptions()

	fmt.Println("port is: ", option.port)
	fmt.Println("url is: ", option.url)
	fmt.Println("script is: ", option.scriptLocation)

	createHttpServer(option.url, option.port, getCallExternalScript(option.scriptLocation))
}

