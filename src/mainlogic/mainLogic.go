package mainlogic
import "./dependencyGraph"

type Resource struct {
	// needs to have resource id,
	// dependencies
	// basically clone the config value here (but it's decoupled i guess, maybe better way to do it/ without sharing code betwen)
}

type MainLogic struct {
	dependencyGraph dependencyGraph.DepGraph
}

func New(onResourceStateChange func(nodeId string, newState dependencyGraph.GlobalState)) MainLogic {
	return MainLogic { dependencyGraph: *dependencyGraph.New(onResourceStateChange)}
}
// func AddResource(resource Resource){  } // nicer interface
func (logic *MainLogic) AddResource(resourceName string) {
	logic.dependencyGraph.AddNode(resourceName)
}

func (logic *MainLogic) AddDependency(resourceName string, resourceNameDep string){
	logic.dependencyGraph.AddDependency(resourceName, resourceNameDep)
}

func(logic * MainLogic) SetNodeReady(resourceName  string) {
	logic.dependencyGraph.SetNodeStateLocalReady(resourceName)
}
func (logic *MainLogic) SetNodeNotReady(resourceName string){
	logic.dependencyGraph.SetNodeStateLocalNotReady(resourceName)
}

