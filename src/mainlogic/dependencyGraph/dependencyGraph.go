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

import "errors"
import "fmt"


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
	parents []*Node;		
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
	nodes map[string] *Node;
	statusHandler func();
}
func NewNode(nodeId string) *Node {
	node := Node { NodeId: nodeId, NodeState: NOTREADY, parents: make([] *Node, 0), dependencies: make([]*Node, 0) }
	return &node
}
func New(changeHandler func()) *RootNode {
	node := Node { NodeId: "0", NodeState: NOTREADY, parents: make([] *Node,  0), dependencies: make([]*Node, 0) }
	rootNode := RootNode { Node: &node, statusHandler: changeHandler , nodes: make(map[string]*Node)}
	return &rootNode
}

func (node *RootNode) AddTarget(nodeId string) {

}
func (node *RootNode) HasDependency(nodeId string, nodeIdDep string) bool{
		return false
}
func (rootnode *RootNode) GetNumTargets() int {
	return len(rootnode.Node.dependencies)
}
func (rootnode *RootNode) GetDependencies(nodeId string) [] *Node {
	node := rootnode.nodes[nodeId]
	fmt.Println(node)
	return make([] *Node, 0)
}
func (rootnode *RootNode) GetDepString(nodeId string) string {
	return "[test get dep string]"
}
func (rootnode *RootNode) AddDependency(nodeId string, nodeIdDep string) error{	
	// when we add a dependency, we add it as not ready, update all the parents
	// also should probably check if a dependency is added twice
	//  also  prevent circular dependencies here too (check if the node youre dependending on depends on has ancestors to you)
	// if someone needs a hackey shitty circular thing, should be done via overloading the config type probably	
	parentNode, ok := rootnode.nodes[nodeId]

	// this if statement may be better served in a function called AddTarget (maybe)
	// concept of target may also be stupid an unnecessary
	if !ok {
		fmt.Println("nodeId: ", nodeId, " not yet a target")

		// use create rootTarget  instead
		parentNode = NewNode(nodeId)
		rootnode.nodes[nodeId] = parentNode

		// children of the rootnode are  "targets"
		rootnode.Node.dependencies = append(rootnode.Node.dependencies, parentNode)
	}
	//////

	if rootnode.HasDependency(nodeId, nodeIdDep){
		return errors.New("circular dependency")
	}
	dependencyNode, depOk := rootnode.nodes[nodeIdDep]
	fmt.Println("1: ", dependencyNode)
	if !depOk {
		dependencyNode = NewNode(nodeIdDep)
		fmt.Println("2: ", dependencyNode)
		rootnode.nodes[nodeIdDep] = dependencyNode
	}
	fmt.Println("3: ", dependencyNode)

	parentNode.dependencies = append(parentNode.dependencies, dependencyNode)
	dependencyNode.parents = append(dependencyNode.parents, parentNode)
	fmt.Println("len parents is: ", len(dependencyNode.parents))
	fmt.Println(rootnode)

	return nil
}
func (node *Node) RemoveDependency(nodeId string, nodeIdDep string){	// remove chop the dependency off, and update the parents
	
}



/*

		()------- \
		/ \       |
	    V  V
	   ()-> ()       |
	    \ /        |
	     V
	     ()---------
*/