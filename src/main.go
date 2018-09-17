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



func main() {	
	config, err := parseConfig.ParseYamlConfig("./examples/simple-config.yaml")
	mainlogic.New(func(nodeIdChange string){
		fmt.Println("node id change: ", nodeIdChange)
	})
	if err != nil {
		panic("could not parse config")
	}

	// Add dependencies from the configuration
	fmt.Println(config)

	resourceName := config.ResourceName
	fmt.Println("add resource name here: ", resourceName)
	//logic.AddResource(resourceName)

	pluginValues, _ := plugins.GetPlugins("./res/plugins")
	for _, plugin := range(pluginValues){
		plugin.Setup("0") // what is the point of this string? should get rid of no id is needed to setup...
	}


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

	// wait for siterm (aka ctrl-c from terminal)
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	// teardown slaves after signal received
	fmt.Println("teardown slaves placeholder")

}
