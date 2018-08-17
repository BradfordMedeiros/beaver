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
import "fmt"

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
	nodes map[string]*Node
}

func NewNode(nodeId string) *Node {
	node := Node{NodeId: nodeId, NodeState: NOTREADY, parents: make([]*Node, 0), dependencies: make([]*Node, 0)}
	return &node
}
func New() *RootNode {
	rootNode := RootNode{nodes: make(map[string]*Node)}
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
func traverseDependencies(nodeId string, visitedNodes *map[string]*Node, allNodes map[string]*Node) {
	_, nodeVisited := (*visitedNodes)[nodeId]
	if nodeVisited {
		return
	}

	node := allNodes[nodeId]
	(*(visitedNodes))[nodeId] = node

	for _, dependencyNode := range node.dependencies {
		traverseDependencies(dependencyNode.NodeId, visitedNodes, allNodes)
	}
}
func (rootnode *RootNode) GetDependencies(nodeId string) map[string]*Node {
	dependencies := make(map[string]*Node, 0)
	node := rootnode.nodes[nodeId]
	for _, depNode := range node.dependencies {
		traverseDependencies(depNode.NodeId, &dependencies, rootnode.nodes)
	}
	return dependencies
}
func (rootnode *RootNode) GetNumImmediateParents(nodeId string) (int, error) {
	node, ok := rootnode.nodes[nodeId]
	fmt.Println("ok: ", ok)
	if !ok {
		return -0, errors.New("node " + nodeId + " is not in the graph")
	}

	return len(node.parents), nil
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
	// maybe i should explicitly have an add target instead of doing that implicitly?
	parentNode, parentOk := rootnode.nodes[nodeId]
	if !parentOk {	
		parentNode = NewNode(nodeId)
		rootnode.nodes[nodeId] = parentNode
	}

	dependencyNode, depOk := rootnode.nodes[nodeIdDep]
	if !depOk {
		dependencyNode = NewNode(nodeIdDep)
		rootnode.nodes[nodeIdDep] = dependencyNode
	}
	////////////////////////////////////////////////////////////

	if !rootnode.CanAddDependency(nodeId, nodeIdDep) {
		return errors.New("circular dependency, adding [" + nodeIdDep + "] as dependency to [" + nodeId + "]")
	}

	parentNode.dependencies = append(parentNode.dependencies, dependencyNode)
	dependencyNode.parents = append(dependencyNode.parents, parentNode)

	return nil
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
