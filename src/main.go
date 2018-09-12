package main

import "fmt"
import "./parseConfig"

func AddDependenciesToLogic(config parseConfig.Config){

}

func main() {	
	config, err := parseConfig.ParseYamlConfig("./examples/simple-config.yaml")
	if err != nil {
		panic("could not parse config")
	}
	AddDependenciesToLogic(config)
	fmt.Println(config)
	fmt.Println(err)
}
