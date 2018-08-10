/*
	used to parse the build config object types
	these configs represent dependencies
*/

package parseConfig

import (
	"fmt"
	"io/ioutil"
)
import "gopkg.in/yaml.v2"	


type Config struct {
	ResourceName string `yaml:"ResourceName"`
	Dependencies []Config `yaml:"Dependencies"`
}

func parseYamlString(yamlConfig []byte) (Config, error) {
	fmt.Print("parse config placeholder")
	var resource Config
	err := yaml.Unmarshal(yamlConfig, &resource)
	if err != nil {
		return Config {}, err
	}

	return resource, nil
}

func ParseYamlConfig(yamlConfig string) (Config, error) {
	fileContent, err := ioutil.ReadFile(yamlConfig)
	if err != nil {
		return Config{}, err
	}
	
	parsedConfig, errParsing := parseYamlString(fileContent)
	if errParsing != nil {
		return Config{}, err
	}

	return parsedConfig, nil
}


