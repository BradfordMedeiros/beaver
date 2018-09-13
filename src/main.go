package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)
import "./parseConfig"
import "./mainlogic"
import "./plugins"

func AddDependenciesToLogic(logic mainlogic.MainLogic, config parseConfig.Config){
	fmt.Println("add dependencies placeholder")
	fmt.Println(config)

	resourceName := config.ResourceName
	logic.AddResource(resourceName)

	pluginValues, _ := plugins.GetPlugins("./res/plugins")
	for _, plugin := range(pluginValues){
		plugin.Setup("0") // what is the point of this string? should get rid of no id is needed to setup...
	}

	var options []plugins.PluginOption
	for _, option := range config.Options {
		fmt.Println("new option ", option)
	}

	/*	fmt.Println("config options length is: ", len(config.Options))
		for _, option := range config.Options {
			options = append(options, plugins.PluginOption{
				Option: option.Option,
				Value:  option.Value,
			})
		}

		abspath, err := filepath.Abs("./commonScripts/alert-ready.sh")
		err1 := plugin.AddResource(id, options, abspath+" "+id)*/
	fmt.Println(len(options))	
	


}
func SetupSlaves(logic mainlogic.MainLogic){
	// for each slave, check if valid resource, then add resource
	fmt.Println("setup slaves placeholder")
}
func TeardownSlaves (logic mainlogic.MainLogic){
	// for each slave, call teardown
	fmt.Println("teardown slaves placeholder")
}


func main() {	
	config, err := parseConfig.ParseYamlConfig("./examples/simple-config.yaml")
	mainsystem := mainlogic.New(func(nodeIdChange string){
		fmt.Println("node id change: ", nodeIdChange)
	})
	if err != nil {
		panic("could not parse config")
	}
	AddDependenciesToLogic(mainsystem, config)
	SetupSlaves(mainsystem)

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	TeardownSlaves(mainsystem)
}
