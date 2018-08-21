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

package acyclicGraph

import "errors"

type Node struct {
	NodeId       string
	dependencies []*Node
	parents      []*Node
}

type RootNode struct {
	nodes map[string]*Node
}

func NewNode(nodeId string) *Node {
	node := Node{NodeId: nodeId,  parents: make([]*Node, 0), dependencies: make([]*Node, 0)}
	return &node
}
func New() *RootNode {
	rootNode := RootNode{nodes: make(map[string]*Node)}
	return &rootNode
}
func(node *Node) GetParents() []*Node{
	return node.parents
}
func(node *Node) GetDependencies() []*Node{
	return node.dependencies
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

func (rootnode *RootNode) HasNode(nodeId string) bool{
	_, nodeInGraph := rootnode.nodes[nodeId]
	return nodeInGraph
}
func (rootnode *RootNode) GetNode(nodeId string) (*Node, error) {
	node, nodeInGraph := rootnode.nodes[nodeId]
	if nodeInGraph == false {
		return (&Node{}), errors.New("node not in graph")
	}
	return node, nil
}

