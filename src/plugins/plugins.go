/*
	handles stateful logic of beaver
	for example, should start/stop the main server

	should be able to package the folders

	should have the dependencies
	invoke the triggers for the dependencies, so on

	this should only have the functions that do this

	the actually mechanism to parse and then add should be done elsewhere


*/
package plugins
import "io/ioutil"	
import "path/filepath"	
import "errors"

// These simply load the set up available plugins, and verify z

func LoadPlugins(){

}
func IsValidResource(){

}

type Plugin struct {
	PluginName string;
	PluginFolderPath string;
}

func GetPlugins(pluginFolderPath string) ([]Plugin, error) {
	files, err := ioutil.ReadDir(pluginFolderPath)

	plugins :=  []Plugin { }
	for _, file := range files {
		if file.IsDir() {
			fileName := file.Name()
			fullPath :=  filepath.Join(pluginFolderPath, fileName)
			plugins = append(plugins, Plugin { PluginName: fileName, PluginFolderPath: fullPath })
		}else{
			return nil, errors.New("file found in plugin folder that is not a valid plugin")
		}
	}

	if err != nil{
		return nil, err
	}

	return plugins, nil
}



// These are applied per type during runtime
func SetupPlugin(){

}
func SetupPlugins(){

}
func TeardownPlugin(){

}
func TeardownPlugins(){

}