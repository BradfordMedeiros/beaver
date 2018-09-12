package main

import "fmt"
import "./parseConfig"
import "./mainlogic"

func AddDependenciesToLogic(config parseConfig.Config){
}
func SetupSlaves(logic mainlogic.MainLogic){
	// for each slave, check if valid resource, then add resource
}
func TeardownSlaves (logic mainlogic.MainLogic){
	// for each slave, call teardown
}


func main() {	
	config, err := parseConfig.ParseYamlConfig("./examples/simple-config.yaml")
	mainsystem := mainlogic.New(func(nodeIdChange string){
		fmt.Println("node id change: ", nodeIdChange)
	})
	if err != nil {
		panic("could not parse config")
	}
	AddDependenciesToLogic(config)
	SetupSlaves(mainsystem)


	fmt.Println(config)
	fmt.Println(err)
	fmt.Println(mainsystem)
}
