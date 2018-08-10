/*
	used to parse the build config object types
	these configs represent dependencies
*/

package parseConfig

import "fmt"
import "gopkg.in/yaml.v2"	


type Config struct {
	ResourceName string `yaml:"ResourceName"`
}

func ParseYamlConfig(yamlConfig string) {

}

func ParseYamlString(yamlConfig string) (error, Config) {
	fmt.Print("parse config placeholder")
	var resource Config
	err := yaml.Unmarshal([]byte(yamlConfig), &resource)
	if err != nil {
		return err, Config {}
	}

	return nil, resource
}
