
package options

import "flag"
import "path/filepath"

type Options struct {
	Verbose bool;
	PluginDirectory string;
}

func GetOptions() (*Options, error) {
	verbose := flag.Bool("v", false, "enable verbose output")
	pluginLocation := flag.String("p", "./plugins", "location of plugin folder")

	flag.Parse()

	pluginDirectory, err := filepath.Abs(*pluginLocation)

	if err != nil {
		return nil, err
	}

	option := Options { 
		Verbose: *verbose, 
		PluginDirectory: pluginDirectory,
	}

	return &option, nil
}