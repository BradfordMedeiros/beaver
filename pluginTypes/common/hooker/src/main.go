package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ConfigFlag struct {
	isSet        bool
	fileLocation string
}

func (flag *ConfigFlag) Set(s string) error {
	flag.isSet = true
	flag.fileLocation = s
	return nil
}
func (flag *ConfigFlag) String() string {
	return flag.fileLocation
}

type Options struct {
	port           uint
	url            string
	scriptLocation string
	verbose        bool
	config         ConfigFlag
}

// @todo parse flags for real using flag package
func getOptions() Options {
	port := flag.Uint("p", 8000, "port number to listen on")
	url := flag.String("u", "/", "url to listen on")
	verbose := flag.Bool("v", false, "enable verbose output")
	scriptLocation := flag.String("s", "./hook.sh", "location of script to call when hook triggered")
	configFile := ConfigFlag{isSet: false}
	flag.Var(&configFile, "c", "location of config file")

	flag.Parse()

	option := Options{
		port:           *port,
		url:            *url,
		scriptLocation: *scriptLocation,
		verbose:        *verbose,
		config:         configFile,
	}

	return option
}

// this should just create a map[string]func
func createHttpServer(port uint, endpointToRequest map[string]func()) {
	for endpointUrl, functionHandler := range endpointToRequest {
		handle := functionHandler
		http.HandleFunc(endpointUrl, func(w http.ResponseWriter, r *http.Request) {
			handle()
			w.Write([]byte("ok"))
		})
	}
	http.ListenAndServe(":"+strconv.Itoa(int(port)), nil)
}

// @todo execute script for real
// maybe want rate limiter too, but idk for now
func getCallExternalScript(scriptLocation string) func() {
	return func() {
		err := exec.Command("/bin/sh", "-c", scriptLocation).Run()
		if err != nil {
			fmt.Println("error executing script")
			fmt.Println(err)
		}

	}
}

func generateRequestMapFromConfig(configFileLocation string) (map[string]func(), error) {
	fileContent, err := ioutil.ReadFile(configFileLocation)
	if err != nil {
		return nil, err
	}

	configFileString := strings.TrimSpace(string(fileContent))
	configPairs := strings.Split(configFileString, "\n")

	requestMap := make(map[string]func())
	for _, config := range configPairs {
		configPairSplit := strings.SplitN(config, " ", 2)
		if len(configPairSplit) != 2 {
			fmt.Println(len(configPairSplit))
			return make(map[string]func()), errors.New("invalid config")
		}
		fmt.Println("url: ", configPairSplit[0])
		fmt.Println("action: ", configPairSplit[1])
		requestMap[configPairSplit[0]] = getCallExternalScript(configPairSplit[1])
	}

	return requestMap, nil
}

func main() {
	option := getOptions()

	if option.verbose {
		fmt.Println("port is: ", option.port)
		fmt.Println("url is: ", option.url)
		fmt.Println("script is: ", option.scriptLocation)
	}

	var requestMap map[string]func()

	if option.config.isSet == false {
		requestMap = map[string]func(){
			option.url: getCallExternalScript(option.scriptLocation),
		}
	} else {
		fmt.Println("using config  file")
		funcMap, err := generateRequestMapFromConfig(option.config.fileLocation)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		requestMap = funcMap
	}

	//var requestMap map[string]func
	createHttpServer(option.port, requestMap)
}
