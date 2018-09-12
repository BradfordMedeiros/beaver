package main

import "fmt"
import "./parseConfig"

func main() {	
	config, err := parseConfig.ParseYamlConfig("./examples/simple-config.yaml")
	fmt.Println(config)
	fmt.Println(err)
}
