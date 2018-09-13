package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)
import "./parseConfig"
import "./mainlogic"

func AddDependenciesToLogic(logic mainlogic.MainLogic, config parseConfig.Config){
	fmt.Println("add dependencies placeholder")
	fmt.Println(config)

	//resourceName := config.ResourceName

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
