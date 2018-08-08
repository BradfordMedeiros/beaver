
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Options struct {
	port uint;
	url string;
}

func getOptions () Options{
	option := Options { port: 3000, url: "/someendpoint" }
	return option
}

func createHttpServer(endpoint string, port uint, onRequest func()) {
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request){
		onRequest()
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":"+strconv.Itoa(int(port)), nil)
}

func getCallExternalScript(scriptLocation string) func(){
	return func(){
		fmt.Println("call external placeholder")
	}
}

func main() {
	option := getOptions()

	fmt.Println("port is: ", option.port)
	fmt.Println("url is: ", option.url)

	createHttpServer(option.url, option.port, func(){
		fmt.Println("external func called")
	})
}

