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

type Option struct {
	option string `yaml:"option"`;
	value string `yaml:"value"`;
}

type Config struct {
	ResourceName string `yaml:"name"`
	Dependencies []Config `yaml:"depends-on"`
	PluginType string `yaml:"plugin-type"`
	Options []Option `yaml:"options"`
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


