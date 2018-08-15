/*
	tree structure to  represent dependency graph within the program

	idea is we mark individual nodes are resolved
	cannot mark a node as resolved w/ out dependent nodes being "resolved"

	elsewhere this will be used by listening for the commit hooks
	those commit hooks processed serially
	and we resolve them if possible
	choosing the nodes that have unresolved dependencies first

	i think
*/

package dependencyGraph


type State int;
const (
	NOTREADY State = 0	// the plugin has not yet alerted as being ready for build (alert not called)
	READY State = 1		// the plugin has called alert, is ready to  build (build has not yet been called)
	INPROGESS State = 2 // the plugin has called build, but has notyet called complete
	COMPLETE State = 3	// the plugin has finished the build
	ERROR State = 4		// the plugin has declared an error, this probably needs manual reset
)

type Node struct {
	NodeId  string;
	NodeState State;
	dependencies [] *Node;
	parent *Node;
}

func (node* Node) UpdateNodeState(){
	// this should go to the parent, check the children, update new value, 
	// and then update that fucker, if change repeat until nil parents
	// we shoujld not broadcast a new state until this full computation is done
}
func (node *Node) SetNotReady(){
	node.NodeState = NOTREADY
	node.UpdateNodeState()
}
func (node *Node) SetReady(){
	node.NodeState = READY
	node.UpdateNodeState()
}
func (node *Node) SetInProgress(){
	node.NodeState = INPROGESS
	node.UpdateNodeState()
}
func (node *Node) SetComplete(){
	node.NodeState = COMPLETE
	node.UpdateNodeState()
}
func(node *Node) SetError(){
	node.NodeState = ERROR
	node.UpdateNodeState()
}


type RootNode struct {
	Node *Node;
	statusHandler func();
}
func New(changeHandler func()) *RootNode {
	node := Node { NodeId: "0", NodeState: NOTREADY, parent: nil, dependencies: make([]*Node, 0) }
	rootNode := RootNode { Node: &node, statusHandler: changeHandler }
	return &rootNode
}
func (node *Node) AddDependency(dependency  *Node){	
	// when we add a dependency, we add it as not ready, update all the parents
	// also should probably check if a dependency is added twice
	//  also  prevent circular dependencies here too (check if the node youre dependending on depends on has ancestors to you)
	// if someone needs a hackey shitty circular thing, should be done via overloading the config type probably
}
func (node *Node) RemoveDependency(dependency *Node){	// remove chop the dependency off, and update the parents
	type Thing struct {
		a string;
	}
}



/*

		()------- \
		/\         |
	   () ()       |
	    \ /        |
	     ()---------
*/