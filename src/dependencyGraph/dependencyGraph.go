
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

package dependencyGraph;

import "fmt"

type Node struct {
	leftNode *Node;
	rightNode *Node;
	NodeName string;
	resolved bool;
}

type RootNode struct {
	Node *Node;

	// maybe look up better way to generate ids
	// but this works well with no external deps
	getNextNodeId func() int;	
}


func New() *RootNode{
	node := Node { NodeName: "hello", resolved: false }

	currId := -1
	getNextNodeId := func() int{
		currId = currId + 1
		return currId
	}

	fmt.Print("wow : ",  currId, "|")
	someNode := RootNode { Node: &node, getNextNodeId: getNextNodeId }

	return &someNode
}

func (graph RootNode) AddDependency(node string, node2 string){

}

func (graph RootNode) HasDependency(node string, node2 string){

}

func (graph RootNode) RemoveDependency(node string, node2 string){

}

func (graph RootNode) MarkAsResolved(node string){

}

func (graph RootNode) MarkAsUnresolved(node string){

}

func (graph RootNode) Size() int{
	return 10
}

