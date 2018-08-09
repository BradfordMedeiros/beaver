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
func (plugin *Plugin) Build() error {
	fmt.Println("plugin build: ", plugin.PluginName)
	payload := plugin.getBuildLocation()
	command := exec.Command("/bin/sh", "-c", payload)
	command.Dir = plugin.PluginFolderPath
	err := command.Run()
	return err
}

// valid plugin needs setup.sh, teardown.sh, and build.sh
func (plugin *Plugin) isValidResource() bool {
	setupLocation := plugin.getSetupLocation()
	teardownLocation := plugin.getTeardownLocation()
	buildLocation := plugin.getBuildLocation()

	_, errSetup := os.Stat(setupLocation)
	setupExists := errSetup == nil

	_, errTeardown := os.Stat(teardownLocation)
	teardownExists := errTeardown == nil

	_, errBuild := os.Stat(buildLocation)
	buildExists := errBuild == nil

	return setupExists && teardownExists && buildExists
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
