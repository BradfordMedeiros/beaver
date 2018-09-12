package mainlogic

import "fmt"

/*
	parse config and get config
	add it to mainlogic here
	open server to start listening for alert statuses
	when we get the status we call the relevant command hooks here
	command hooks here trigger rebuilds or whatever outside
*/
// add all resources here
// add resource id
// remove resource id
// get resource status
// on resources status change
/*
	deleting a resource will also delete all of the child resources --> maybe should support multiple targets
	// when this supports more than one target we can support this

	someid: Resource,
	anotherid: Resource,
	etc

	and maybe the dependency graph right here
		somideid
		/		\
	anotherid	something here

	tree has someething like this:?


	// maybe ready with depndencies vs ready without deps?
	-> onReady()
		ready
	ready  not-ready

	goes to
		ready
	build not-ready

		ready
	complete complete

		build
	complete complete

		complete
	complete complete



*/

type Resource struct {
	// needs to have resource id,
	// dependencies
	// basically clone the config value here (but it's decoupled i guess, maybe better way to do it/ without sharing code betwen)
}

type MainLogic struct {
	Dependencies map[string]string
}

// func AddResource(resource Resource){  } // nicer interface
func (logic *MainLogic) AddResource(resourceName string, resourceValue string) {
	logic.Dependencies[resourceName] = resourceValue
}
func (logic *MainLogic) RemoveResource(resourceName string, resourceValue string) {
	panic("not yet implemented")
}

func Test() {
	fmt.Println("hello world")
}
