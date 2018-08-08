
package main

import (
	"fmt"
	"encoding/json"
)

//import "./parseCommand"
//import "./executeCommand"
//import "./dependencyGraph"
//import "./parseConfig"
import "./options"
import "./plugins"

type Command struct {
	commandType string;
}


func main(){
	options, err := options.GetOptions()
	if err != nil {
		panic("got nil options")
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

	executeCommand := executeCommand.GetExecuteCommand()

	executeCommand()

	graph := dependencyGraph.New()
	fmt.Println(graph.Size())

	val, _ := json.Marshal(*graph)
	valString := string(val)
	fmt.Println(valString)

	parseConfig.ParseConfig()
	*/
	plugs := plugins.GetPlugins(options.PluginDirectory)

	if options.Verbose {
		fmt.Println("Number of plugins: \n", len(plugs))
	}

}