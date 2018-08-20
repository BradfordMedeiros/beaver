
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

type State int

const (
	NOTREADY  State = 0 // the plugin has not yet alerted as being ready for build (alert not called)
	READY     State = 1 // the plugin has called alert, is ready to  build (build has not yet been called)
	QUEUED 	  State = 2
	INPROGESS State = 3 // the plugin has called build, but has notyet called complete
	COMPLETE  State = 4 // the plugin has finished the build
	ERROR     State = 5 // the plugin has declared an error, this probably needs manual reset
)

type DepGraph struct {
	acyclicGraph *acyclicGraph.RootNode;
	nodeIdToLocalState map[string] State;
	nodeIdToGlobalState map[string] State;
}

func New() *DepGraph  {
	graph := &DepGraph { 
		acyclicGraph: acyclicGraph.New(), 
		nodeIdToLocalState: make(map[string]State), 
		nodeIdToGlobalState: make(map[string]State),
	}
	return graph
}

func (graph *DepGraph) AddDependency(nodeId string, depNodeId string) error{
	err := graph.acyclicGraph.AddDependency(nodeId, depNodeId)
	if err !=nil {
		return err
	}
	graph.nodeIdToLocalState[nodeId] = NOTREADY
	graph.nodeIdToGlobalState[nodeId] = NOTREADY
	return nil
}

func (graph *DepGraph) SetNodeStateLocalNotReady(nodeId string) error{
	return nil
}
func (graph *DepGraph) SetNodeStateLocalReady(nodeId string) error{
	return nil
}

func (graph *DepGraph) SetNodeNotReady(nodeId string) error {
	// @todo
	// not ready, so we should traverse all parent nodes and mark them as global not ready
	// can come from any state
	return nil
}
func (graph *DepGraph) SetNodeStateReady(nodeId string) error {
	// @todo
	// node global ready, so we should simply mark this.  do not do anything since parents need completion
	// must come from not ready state
	return nil
}
func (graph *DepGraph) SetNodeStateQueued(nodeId string) error{
	// @todo
	// simple marker
	// must come from ready state
	return nil
}
func (graph *DepGraph) SetNodeStateInProgress(nodeId string) error{
	// @todo 
	// simple marker
	// must come from queued state
	return nil
}
func (graph *DepGraph) SetNodeStateComplete(nodeId string) error {
	// @todo 
	// should traverse parents, and if the parent has all complete childen, set the parent ready
	// must come from in progress state
	return nil
}
func (graph *DepGraph) SetNodeStateError(nodeId string) error {
	return nil
}

