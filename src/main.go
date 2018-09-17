package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"path/filepath"
)
import "./parseConfig"
import "./mainlogic"
import "./plugins"
import "./plugins/pluginResource"
import "./ioLoop"


func main() {	
	// INIT PART
	fmt.Println("----------INIT SECTION----------\n")

	config, err := parseConfig.ParseYamlConfig("./examples/simple-config.yaml")
	logic := mainlogic.New(func(nodeIdChange string){
		fmt.Println("node id change: ", nodeIdChange)
	})
	if err != nil {
		panic("could not parse config")
	}

	// Add dependencies from the configuration
	fmt.Println(config)

	resourceName := config.ResourceName
	fmt.Println("add resource name here: ", resourceName)
	logic.AddResource(resourceName)

	folderPath, _ := filepath.Abs("./res/plugins")
	pluginGroup, _ := plugins.Load(folderPath, func(eventName string){
		fmt.Println("event: ", eventName)
	})
	pluginGroup.Setup()

	// setup individual slaves
	// for each slave, check if valid resource, then add resource
	fmt.Println("setup slaves placeholder")

	var options []pluginResource.PluginOption
	for _, option := range config.Options {
		options = append(options, pluginResource.PluginOption{
			Option: option.Option,
			Value:  option.Value,
		})
	}

	abspath, _ := filepath.Abs("./commonScripts/alert-ready.sh")
	fmt.Println(abspath)
	errAddRes := pluginGroup.AddResource(config.PluginType, "test-id", options, abspath)
	fmt.Println("error: ", errAddRes)
	if errAddRes != nil {
		panic("resource error add")
	}

	// DURING LIFE OF PROGRAM PART

	// wait for sigterm (aka ctrl-c from terminal)
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	fmt.Println("----------SIGNAL WAITING SECTION----------\n")

	listenMessages := func(){
		commandChannel := make(chan string) 
		go ioLoop.StartRepl(commandChannel)
		for true {
			select {
				case commandString := <-commandChannel: {
					if commandString == "quit" {
						os.Exit(0)
					}
				}
			}
		}
	}

	go listenMessages()
	<-signalChannel
	
	//<-signalChannel
	fmt.Println("----------SIGNAL CAUGHT SECTION----------\n")


	// DEINIT
	fmt.Println("----------DEINIT SECTION----------\n")

	errRemRes := pluginGroup.RemoveResource(config.PluginType, "test-id", options)
	if errRemRes != nil {
		panic("resource error remove")
	}

	// teardown slaves after signal received
	pluginGroup.Teardown()

}
