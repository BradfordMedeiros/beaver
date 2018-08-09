package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//import "./parseCommand"
//import "./dependencyGraph"
//import "./parseConfig"
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

	list := func() {
		plugs, err := plugins.GetPlugins(options.PluginDirectory)
		if err != nil {
			fmt.Println("error reading plugins")
			fmt.Println(err)
			os.Exit(1)
		}

		jsonPlugs, _ := json.Marshal(plugs)
		fmt.Println(string(jsonPlugs))
	}

	setup := func() {
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
	teardown := func() {
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
	exit := func() {
		os.Exit(0)
	}

	build := func(){
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

	commandMap := map[string]func(){
		"list":     list,
		"setup":    setup,
		"teardown": teardown,
		"build": build,
		"exit":     exit,
	}

	commandChannel := make(chan string)
	if options.LoopType == "repl" {
		go ioLoop.StartRepl(commandChannel)
		for true {
			command := <-commandChannel
			commandToExecute := commandMap[command]
			if commandToExecute == nil {
				fmt.Println("invalid command")
			} else {
				commandToExecute()
			}
		}
	}

}
