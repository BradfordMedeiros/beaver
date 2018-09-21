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
import "./mainlogic/dependencyGraph"
import "./plugins"
import "./plugins/pluginResource"
import "./ioLoop"


func main() {	
	// INIT PART
	fmt.Println("----------INIT SECTION----------\n")

	folderPath, _ := filepath.Abs("./res/plugins")
		pluginGroup, _ := plugins.Load(folderPath, func(eventName string){
		fmt.Println("event: ", eventName)
	})

	config, err := parseConfig.ParseYamlConfig("./examples/simple-config.yaml")

	var options []pluginResource.PluginOption
	for _, option := range config.Options {
		options = append(options, pluginResource.PluginOption{
			Option: option.Option,
			Value:  option.Value,
		})
	}

	logic := mainlogic.New(func(nodeId string, oldState dependencyGraph.GlobalState, newState dependencyGraph.GlobalState){
		fmt.Println("node id change: id: ", nodeId, "old state : ", oldState, " new state: ", newState)
		// add in old state to this, and then new state
		// then we can do things like if becomes ready, we set to queued, if queued set to in progress and invoke build, etc

		if (oldState == dependencyGraph.NOTREADY && newState == dependencyGraph.READY){
			fmt.Println(nodeId, " became ready")
			pluginGroup.Build(config.PluginType, nodeId, options)

		}else if (oldState == dependencyGraph.READY && newState == dependencyGraph.NOTREADY){
			fmt.Println(nodeId, " became not ready")
		}
	})
	if err != nil {
		panic("could not parse config")
	}

	// Add dependencies from the configuration
	fmt.Println(config)

	resourceName := config.ResourceName
	fmt.Println("add resource name here: ", resourceName)
	logic.AddResource(resourceName)

	
	// calls the setup.sh scripts for each plugin
	pluginGroup.Setup()

	// setup individual slaves
	// for each slave, check if valid resource, then add resource
	fmt.Println("setup slaves placeholder")

	

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
					}else if commandString == "ready1" {
						logic.SetNodeReady(resourceName)	// these are race conditions.  be careful. ok for now
					}else if commandString == "notready1" {
						logic.SetNodeNotReady(resourceName)
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

	// calls the teardown.sh scripts for each plugin
	pluginGroup.Teardown()

}
