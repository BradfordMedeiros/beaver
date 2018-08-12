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
import "fmt"
import "os"
import "os/exec"

// These simply load the set up available plugins, and verify z

type Plugin struct {
	PluginName       string
	PluginFolderPath string
}

type PluginOption struct {
	Option string;
	Value string;
}

// @todo would be nice to log stderr, stdout somewhere
func (plugin *Plugin) Setup(id string) error {
	fmt.Println("plugin setup: ", plugin.PluginName)
	payload := plugin.getSetupLocation() + " " + id
	command := exec.Command("/bin/sh", "-c", payload)
	command.Dir = plugin.PluginFolderPath
	err := command.Run()
	return err
}
func (plugin *Plugin) Teardown(id string) error {
	fmt.Println("plugin teardown: ", plugin.PluginName)
	payload := plugin.getTeardownLocation() + " " + id
	command := exec.Command("/bin/sh", "-c", payload)
	command.Dir = plugin.PluginFolderPath
	err := command.Run()
	return err
}
func (plugin *Plugin) Build(id string, options string) error {
	fmt.Println("plugin build: ", plugin.PluginName)
	payload := plugin.getBuildLocation() + " " + id + " '" + options +"'" 
	command := exec.Command("/bin/sh", "-c", payload)
	command.Dir = plugin.PluginFolderPath
	err := command.Run()
	return err
}
func pluginOptionsToString(options []PluginOption) string {
	var optionsString string
	fmt.Println("optitons length is: ", len(options))
	for _, option := range(options){
		optionName := option.Option
		optionValue := option.Value
		optionsString = optionsString + optionName + " " + optionValue + "\n"
	}
	return optionsString
}
func (plugin *Plugin) AddResource(id string, options []PluginOption) error {
	fmt.Println("plugin add resource: ", plugin.PluginName)
	payload := plugin.getAddResourceLocation() 
	command := exec.Command("/bin/sh", "-c", payload)
	command.Dir = plugin.PluginFolderPath
	command.Env = os.Environ()
	command.Env = append(command.Env, "ID=testid")
	command.Env = append(command.Env, "OPTIONS=" + pluginOptionsToString(options))
	err := command.Run()
	return err
}
func (plugin *Plugin) RemoveResource(id string, options []PluginOption) error {
	fmt.Println("plugin remove resource: ", plugin.PluginName)
	payload := plugin.getRemoveResourceLocation()
	command := exec.Command("/bin/sh", "-c", payload)
	command.Dir = plugin.PluginFolderPath
	command.Env = os.Environ()
	command.Env = append(command.Env, "ID=testid")
	command.Env = append(command.Env, "OPTIONS=" + pluginOptionsToString(options))
	err := command.Run()
	return err
}

// valid plugin needs setup.sh, teardown.sh, and build.sh
func (plugin *Plugin) isValidResource() bool {
	setupLocation := plugin.getSetupLocation()
	teardownLocation := plugin.getTeardownLocation()
	buildLocation := plugin.getBuildLocation()
	addResLocation := plugin.getAddResourceLocation()
	removeResLocation := plugin.getRemoveResourceLocation()

	_, errSetup := os.Stat(setupLocation)
	setupExists := errSetup == nil

	_, errTeardown := os.Stat(teardownLocation)
	teardownExists := errTeardown == nil

	_, errBuild := os.Stat(buildLocation)
	buildExists := errBuild == nil

	_, errAddRes := os.Stat(addResLocation)
	addResExists := errAddRes == nil

	_, errRemoveRes := os.Stat(removeResLocation)
	removeResExists := errRemoveRes == nil

	return setupExists && teardownExists && buildExists && addResExists &&  removeResExists
}
func (plugin *Plugin) getSetupLocation() string {
	return filepath.Join(plugin.PluginFolderPath, "setup.sh")
}
func (plugin *Plugin) getTeardownLocation() string {
	return filepath.Join(plugin.PluginFolderPath, "teardown.sh")
}
func (plugin *Plugin) getBuildLocation() string {
	return filepath.Join(plugin.PluginFolderPath, "build.sh")
}
func (plugin *Plugin) getAddResourceLocation() string {
	return filepath.Join(plugin.PluginFolderPath, "add-resource.sh")
}
func (plugin *Plugin) getRemoveResourceLocation() string{
	return filepath.Join(plugin.PluginFolderPath, "remove-resource.sh")
}

func GetPlugins(pluginFolderPath string) ([]Plugin, error) {
	files, err := ioutil.ReadDir(pluginFolderPath)

	plugins := []Plugin{}
	for _, file := range files {
		if file.IsDir() {
			fileName := file.Name()
			fullPath := filepath.Join(pluginFolderPath, fileName)

			loadedPlugin := Plugin{PluginName: fileName, PluginFolderPath: fullPath}
			if !loadedPlugin.isValidResource() {
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
