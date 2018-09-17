package plugins

import "io/ioutil"
import "path/filepath"
import "errors"
import "./pluginResource"

type PluginGroup struct {

}

func GetPlugins(pluginFolderPath string) ([]pluginResource.Plugin, error) {
	files, err := ioutil.ReadDir(pluginFolderPath)

	plugins := []pluginResource.Plugin{}
	for _, file := range files {
		if file.IsDir() {
			fileName := file.Name()
			fullPath := filepath.Join(pluginFolderPath, fileName)

			loadedPlugin := pluginResource.Plugin{PluginName: fileName, PluginFolderPath: fullPath}
			if !loadedPlugin.IsValidResource() {
				return nil, errors.New("Invalid resources for plugin: " + loadedPlugin.PluginName)
			}
			plugins = append(plugins, loadedPlugin)
		} else {
			return nil, errors.New("file found in plugin folder that is not a valid plugin")
		}
	}

	if err != nil {
		return nil, err
	}

	return plugins, nil
}


func New() PluginGroup {
	group := PluginGroup { }
	return group
}

// teardown all the plugins 
func (pluginGroup *PluginGroup) Setup(){

}

// add a specific resource
func (pluginGroup *PluginGroup) AddResource() {

}

// remove a specific resource
func (pluginGroup *PluginGroup) RemoveResource(){

}

// setup all the plugins
func (pluginGroup *PluginGroup) Teardown(){

}