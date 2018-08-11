package options

import "flag"
import "path/filepath"
import "errors"

type Options struct {
	Verbose         bool
	LoopType        string
	PluginDirectory string
}

func GetOptions() (*Options, error) {
	verbose := flag.Bool("v", false, "enable verbose output")
	loopType := flag.String("l", "none", "ability to add command loop <none, repl>")
	pluginLocation := flag.String("p", "./res/plugins", "location of plugin folder")

	flag.Parse()

	if *loopType != "none" && *loopType != "repl" {
		return nil, errors.New("invalid looptype " + *loopType)
	}

	pluginDirectory, err := filepath.Abs(*pluginLocation)

	if err != nil {
		return nil, err
	}

	option := Options{
		Verbose:         *verbose,
		PluginDirectory: pluginDirectory,
		LoopType:        *loopType,
	}

	return &option, nil
}
