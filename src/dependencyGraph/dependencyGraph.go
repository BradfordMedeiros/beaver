
/*
	tree structure to  represent dependency graph within the program

	
*/

package dependencyGraph;

import "fmt"

type Node struct {
	leftNode *Node;
	rightNode *Node;
	NodeName string;
}

type RootNode struct {
	Node *Node;

	// maybe look up better way to generate ids
	// but this works well with no external deps
	getNextNodeId func() int;	
}


func New() *RootNode{
	node := Node { NodeName: "hello" }

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

func (graph RootNode) Size() int{
	return 10
}

