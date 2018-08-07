
package main

import (
	"fmt"
	"encoding/json"
)

import "./parseCommand"
import "./executeCommand"
import "./dependencyGraph"

type Command struct {
	commandType string;
}


func getFunction(seed int) func(some int) int {
	return func(some int) int {
		return seed + some
	}
}

func Test(num int) int{
	return num + 2
}

func main(){
	fmt.Println("Hello!")
	//testpackage.Test()
	command := parseCommand.ParseCommand()
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

}