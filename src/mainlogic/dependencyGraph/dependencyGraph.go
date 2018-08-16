/*
	tree structure to  represent dependency graph within the program

	idea is we mark individual nodes are resolved
	cannot mark a node as resolved w/ out dependent nodes being "resolved"

	elsewhere this will be used by listening for the commit hooks
	those commit hooks processed serially
	and we resolve them if possible
	choosing the nodes that have unresolved dependencies first


	How we add nodes:
	- graph must be acyclic
	- if you are dumb and have circular dependencies, this can be resolved in "user-land" by creating a separate project

	- to determine if graph is acyclic we:
	- let dep(A) => all dependencies (including child dependencies of A)
	- let dep(A) includes X => X in the set of dependencies of A
	- add(A,B)   => add B as a dependency of A
	-	if dep(B) ! includes A and A != B
	-
*/

package dependencyGraph

import "errors"

type State int

const (
	NOTREADY  State = 0 // the plugin has not yet alerted as being ready for build (alert not called)
	READY     State = 1 // the plugin has called alert, is ready to  build (build has not yet been called)
	INPROGESS State = 2 // the plugin has called build, but has notyet called complete
	COMPLETE  State = 3 // the plugin has finished the build
	ERROR     State = 4 // the plugin has declared an error, this probably needs manual reset
)

type Node struct {
	NodeId       string
	NodeState    State
	dependencies []*Node
	parents      []*Node
}

func (node *Node) UpdateNodeState() {
	// this should go to the parent, check the children, update new value,
	// and then update that fucker, if change repeat until nil parents
	// we shoujld not broadcast a new state until this full computation is done
}
func (node *Node) SetNotReady() {
	node.NodeState = NOTREADY
	node.UpdateNodeState()
}
func (node *Node) SetReady() {
	node.NodeState = READY
	node.UpdateNodeState()
}
func (node *Node) SetInProgress() {
	node.NodeState = INPROGESS
	node.UpdateNodeState()
}
func (node *Node) SetComplete() {
	node.NodeState = COMPLETE
	node.UpdateNodeState()
}
func (node *Node) SetError() {
	node.NodeState = ERROR
	node.UpdateNodeState()
}

type RootNode struct {
	Node  *Node
	nodes map[string]*Node
}

func NewNode(nodeId string) *Node {
	node := Node{NodeId: nodeId, NodeState: NOTREADY, parents: make([]*Node, 0), dependencies: make([]*Node, 0)}
	return &node
}
func New() *RootNode {
	node := Node{NodeId: "0", NodeState: NOTREADY, parents: make([]*Node, 0), dependencies: make([]*Node, 0)}
	rootNode := RootNode{Node: &node, nodes: make(map[string]*Node)}
	return &rootNode
}

func (rootnode *RootNode) CanAddDependency(nodeId string, nodeIdDep string) bool {
	_, depNodeInGraph := rootnode.nodes[nodeIdDep]
	if !depNodeInGraph {
		return true
	}
	dependencies := rootnode.GetDependencies(nodeIdDep)
	_, hasDependency := dependencies[nodeId]
	return !hasDependency
}
func (rootnode *RootNode) GetNumTargets() int {
	return len(rootnode.Node.dependencies)
}
func traverseDependencies(nodeId string, visitedNodes map[string]*Node, allNodes map[string]*Node) {
	_, nodeVisited := visitedNodes[nodeId]
	if nodeVisited {
		return
	}

	node := allNodes[nodeId]
	for _, dependencyNode := range node.dependencies {
		visitedNodes[dependencyNode.NodeId] = dependencyNode
		traverseDependencies(dependencyNode.NodeId, visitedNodes, allNodes)
	}
}
func (rootnode *RootNode) GetDependencies(nodeId string) map[string]*Node {
	dependencies := make(map[string]*Node, 0)
	node := rootnode.nodes[nodeId]
	traverseDependencies(node.NodeId, dependencies, rootnode.nodes)
	return dependencies
}
func (rootnode *RootNode) GetDepString(nodeId string) string {
	dependencies := rootnode.GetDependencies(nodeId)
	dependenciesAsString := ""
	for _, depNode := range dependencies {
		dependenciesAsString = dependenciesAsString + depNode.NodeId + " "
	}
	return dependenciesAsString
}
func (rootnode *RootNode) AddDependency(nodeId string, nodeIdDep string) error {
	// when we add a dependency, we add it as not ready, update all the parents
	// also should probably check if a dependency is added twice
	//  also  prevent circular dependencies here too (check if the node youre dependending on depends on has ancestors to you)
	// if someone needs a hackey shitty circular thing, should be done via overloading the config type probably

	// this is bad because we add the parent as a target even if the dependency fails later on
	// !important this is bad code.  in practice it shouldnt be hit, but i don't like that it can throw an error and fail, but
	// have this sied
	parentNode, ok := rootnode.nodes[nodeId]

	// this if statement may be better served in a function called AddTarget (maybe)
	// concept of target may also be stupid an unnecessary
	if !ok {
		// use create rootTarget  instead
		parentNode = NewNode(nodeId)
		rootnode.nodes[nodeId] = parentNode

		// children of the rootnode are  "targets"
		rootnode.Node.dependencies = append(rootnode.Node.dependencies, parentNode)
	}
	//////

	if !rootnode.CanAddDependency(nodeId, nodeIdDep) {
		return errors.New("circular dependency, adding [" + nodeIdDep + "] as dependency to [" + nodeId + "]")
	}
	dependencyNode, depOk := rootnode.nodes[nodeIdDep]
	if !depOk {
		dependencyNode = NewNode(nodeIdDep)
		rootnode.nodes[nodeIdDep] = dependencyNode
	}
	parentNode.dependencies = append(parentNode.dependencies, dependencyNode)
	dependencyNode.parents = append(dependencyNode.parents, parentNode)

	return nil
}
func (node *Node) RemoveDependency(nodeId string, nodeIdDep string) { // remove chop the dependency off, and update the parents

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
