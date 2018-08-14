package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

//import "./parseCommand"
//import "./dependencyGraph"
import "./parseConfig"
import "./options"
import "./plugins"
import "./ioLoop"

type Command struct {
	commandType string
}

func main() {
	options, err := options.GetOptions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonOptions, _ := json.Marshal(*options)
	if options.Verbose {
		fmt.Println("Options:")
		fmt.Println(string(jsonOptions), "\n")
	}

	//testpackage.Test()
	/*command := parseCommand.ParseCommand()
	fmt.Println("command  type is: " , command.Test)
	jsonObj, _ := json.Marshal(command)
	fmt.Println(string(jsonObj))
	var newCommand parseCommand.Command = parseCommand.Command { Test: "yello" }
	fmt.Println(newCommand.Test)


	graph := dependencyGraph.New()
	fmt.Println(graph.Size())

	val, _ := json.Marshal(*graph)
	valString := string(val)
	fmt.Println(valString)

	parseConfig.ParseConfig()
	*/

	list := func(val string) {
		plugs, err := plugins.GetPlugins(options.PluginDirectory)
		if err != nil {
			fmt.Println("error reading plugins")
			fmt.Println(err)
			os.Exit(1)
		}

		jsonPlugs, _ := json.Marshal(plugs)
		fmt.Println(string(jsonPlugs))
	}

	setup := func(val string) {
		plugs, err := plugins.GetPlugins(options.PluginDirectory)
		if err != nil {
			fmt.Println("error reading plugins")
			fmt.Println(err)
			os.Exit(1)
		}

		for _, plug := range plugs {
			err := plug.Setup("testid")
			if err != nil {
				fmt.Println("error: ", err)
			}
		}
	}
	teardown := func(val string) {
		plugs, err := plugins.GetPlugins(options.PluginDirectory)
		if err != nil {
			fmt.Println("error reading plugins")
			fmt.Println(err)
			os.Exit(1)
		}

		for _, plug := range plugs {
			err := plug.Teardown("testid")
			if err != nil {
				fmt.Println("error: ", err)
			}
		}
	}
	exit := func(val string) {
		os.Exit(0)
	}

	build := func(val string){
		plugs, err := plugins.GetPlugins(options.PluginDirectory)
		if err != nil {
			fmt.Println("error reading plugins")
			fmt.Println(err)
			os.Exit(1)
		}

		for _, plug := range plugs {
			err := plug.Build("testid", "https://github.com/BradfordMedeiros/stork.git stork")
			if err != nil {
				fmt.Println("error: ", err)
			}
		}
	}
	parse := func(val string){
		res, err := parseConfig.ParseYamlConfig("./test.yaml")
		if err != nil {
			fmt.Println("error: ", err)
		}else{
			fmt.Println("name:  ", res.ResourceName)
			fmt.Println("dependencies: ", len(res.Dependencies))
			fmt.Println("type:  ", res.PluginType)
			fmt.Println("options: ", len(res.Options))
		}
	}
	add := func(id string){
		config, err := parseConfig.ParseYamlConfig("./test.yaml")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(config.ResourceName)
		plugs, err := plugins.GetPlugins(options.PluginDirectory)
		if err != nil {
			fmt.Println(err)
			return
		}

		plugin := plugs[0]
		fmt.Println(plugin.PluginName)
		var options []plugins.PluginOption
		fmt.Println("config options length is: ", len(config.Options))
		for _, option := range(config.Options){
			options = append(options, plugins.PluginOption { 
				Option: option.Option, 
				Value: option.Value,
			} )
		}

		abspath, err := filepath.Abs("./commonScripts/alert-ready.sh")
		err1 := plugin.AddResource(id, options, abspath )
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		fmt.Println("success")
	}
	remove := func(id string){
		config, err := parseConfig.ParseYamlConfig("./test.yaml")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(config.ResourceName)
		plugs, err := plugins.GetPlugins(options.PluginDirectory)
		if err != nil {
			fmt.Println(err)
			return
		}

		plugin := plugs[0]
		fmt.Println(plugin.PluginName)
		var options []plugins.PluginOption
		fmt.Println("config options length is: ", len(config.Options))
		for _, option := range(config.Options){
			options = append(options, plugins.PluginOption { 
				Option: option.Option, 
				Value: option.Value,
			} )
		}
		fmt.Println("removed id: ", id)
		err1 := plugin.RemoveResource(id, options)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		fmt.Println("success")
	}

	commandMap := map[string]func(string){
		"list":     list,
		"setup":    setup,
		"teardown": teardown,
		"build": build,
		"exit": exit,
		"parse": parse,
		"add": add,
		"remove": remove,
	}

	commandChannel := make(chan string)
	if options.LoopType == "repl" {
		go ioLoop.StartRepl(commandChannel)
		for true {
			commandString := <-commandChannel
			commandArray := strings.SplitN(commandString, " ", 2)
			command := commandArray[0]
			
			option := ""
			if len(commandArray) > 1 {
				option = commandArray[1]
			}


			fmt.Println("option is: ", option)
			commandToExecute := commandMap[command]
			if commandToExecute == nil {
				fmt.Println("invalid command")
			} else {
				commandToExecute(option)
			}
		}
	}

}
