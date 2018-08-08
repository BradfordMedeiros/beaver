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

// These simply load the set up available plugins, and verify z
	
func LoadPlugins(){

}
func IsValidResource(){

}

type Plugin struct {
	pluginName string;
	pluginFolderPath string;
}

func GetPlugins(pluginFolderPath string) []Plugin {
	plugins := []Plugin{ }
	return plugins
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