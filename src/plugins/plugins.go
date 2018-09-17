package plugins

import "fmt"
import "io/ioutil"
import "path/filepath"
import "errors"
import "./pluginResource"

type PluginGroup struct {
	pluginMapping map[string] pluginResource.Plugin
	onEvent func(string)
}

func Load(pluginFolderPath string, onEvent func(string))(PluginGroup, error) {
	pluginMapping := make(map[string] pluginResource.Plugin)
	group := PluginGroup { pluginMapping: pluginMapping, onEvent: onEvent }

	files, err := ioutil.ReadDir(pluginFolderPath)

	for _, file := range files {
		if file.IsDir() {
			fileName := file.Name()
			fullPath := filepath.Join(pluginFolderPath, fileName)
			fmt.Println("fullpath: ", fullPath)

			loadedPlugin := pluginResource.Plugin{PluginName: fileName, PluginFolderPath: fullPath}
			if !loadedPlugin.IsValidResource() {
				return group, errors.New("Invalid resources for plugin: " + loadedPlugin.PluginName)
			}
			pluginMapping[fileName] = loadedPlugin
		} else {
			return group, errors.New("file found in plugin folder that is not a valid plugin")
		}
	}

	if err != nil {
		return group, err
	}

	return group, nil
}

// teardown all the plugins 
func (pluginGroup *PluginGroup) Setup(){
	for _, plugin := range(pluginGroup.pluginMapping){
		fmt.Println("setup here: ", plugin)
		plugin.Setup("0")
	}
}

// add a specific resource
func (pluginGroup *PluginGroup) AddResource(
	resourceName string, 
	id string, 
	options []pluginResource.PluginOption, 
	alertScriptLocation string,
) error {
	plugin, hasResource := pluginGroup.pluginMapping[resourceName]
	if !hasResource {
		return errors.New("no resource named " + resourceName)
	}
	fmt.Println(plugin)
	plugin.AddResource(id, options, alertScriptLocation)
	return nil
}

// remove a specific resource
func (pluginGroup *PluginGroup) RemoveResource(
	resourceName string, 
	id string, 
	options []pluginResource.PluginOption, 
) error{
	plugin, hasResource := pluginGroup.pluginMapping[resourceName]
	if !hasResource {
		return errors.New("no resource named " + resourceName)
	}
	fmt.Println(plugin)
	plugin.RemoveResource(id, options)
	return nil
}

// setup all the plugins
func (pluginGroup *PluginGroup) Teardown(){
	for _, plugin := range(pluginGroup.pluginMapping){
		fmt.Println("teardown here: ", plugin)
		plugin.Teardown("0")
	}
}