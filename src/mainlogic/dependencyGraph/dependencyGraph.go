
/*

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
*/

package dependencyGraph
import "./acyclicGraph"

type DepGraph struct {
	acyclicGraph *acyclicGraph.RootNode;
}

func New() *DepGraph  {
	return &DepGraph { acyclicGraph: acyclicGraph.New() }
}

