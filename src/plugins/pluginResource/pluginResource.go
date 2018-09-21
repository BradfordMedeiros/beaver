/*
	handles stateful logic of beaver
	for example, should start/stop the main server

	should be able to package the folders

	should have the dependencies
	invoke the triggers for the dependencies, so on

	this should only have the functions that do this

	the actually mechanism to parse and then add should be done elsewhere


*/
package pluginResource

import "path/filepath"
import "fmt"
import "os"
import "os/exec"

// These simply load the set up available plugins, and verify z

type Plugin struct {
	PluginName       string;
	PluginFolderPath string;
}

type PluginOption struct {
	Option string;
	Value string;
}

// @todo would be nice to log stderr, stdout somewhere
func (plugin *Plugin) Setup(id string) error {
	fmt.Println("plugin setup: ", plugin.PluginName)
	payload := plugin.getSetupLocation() + " " + id
	fmt.Println("payload: ", payload)
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
func (plugin *Plugin) Build(id string, options[] PluginOption, alertBuiltLocation string) error {
	fmt.Println("plugin build: ", plugin.PluginName)
	payload := plugin.getBuildLocation()
	command := exec.Command("/bin/sh", "-c", payload)
	command.Dir = plugin.PluginFolderPath
	command.Env = append(command.Env, "ID="+id)	// should check about properly escaping this?
	command.Env = append(command.Env, "OPTIONS=" + pluginOptionsToString(options))
	command.Env = append(command.Env, "FINISHED=" + alertBuiltLocation)

	// should probably add ENV of the build folder here so you can do  cp -r somefile $DESTINATION
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
func (plugin *Plugin) AddResource(id string, options []PluginOption, alertScriptLocation string) error {
	fmt.Println("plugin add resource: ", plugin.PluginName)
	payload := plugin.getAddResourceLocation() 
	command := exec.Command("/bin/sh", "-c", payload)
	command.Dir = plugin.PluginFolderPath
	command.Env = os.Environ()
	command.Env = append(command.Env, "ID="+id)	// should check about properly escaping this?
	command.Env = append(command.Env, "OPTIONS=" + pluginOptionsToString(options))
	command.Env = append(command.Env, "ALERT_READY=" + alertScriptLocation)
	err := command.Run()
	return err
}
func (plugin *Plugin) RemoveResource(id string, options []PluginOption) error {
	fmt.Println("plugin remove resource: ", plugin.PluginName)
	payload := plugin.getRemoveResourceLocation()
	command := exec.Command("/bin/sh", "-c", payload)
	command.Dir = plugin.PluginFolderPath
	command.Env = os.Environ()
	command.Env = append(command.Env, "ID="+id)
	command.Env = append(command.Env, "OPTIONS=" + pluginOptionsToString(options))
	err := command.Run()
	return err
}

// valid plugin needs setup.sh, teardown.sh, and build.sh
func (plugin *Plugin) IsValidResource() bool {
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
