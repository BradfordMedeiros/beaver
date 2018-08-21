
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
import "errors"
import "fmt"

type GlobalState int;
const (
	LOCAL_NOTREADY = 0
	LOCAL_READY = 1
)

type State int

const (
	NOTREADY  State = 0 // the plugin has not yet alerted as being ready for build (alert not called)
	READY     State = 1 // the plugin has called alert, is ready to  build (build has not yet been called)
	QUEUED 	  State = 2 // the scheduler has queued this item for buildqa
	INPROGESS State = 3 // the schedule has invokved build, but has notyet called complete
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
	graph.nodeIdToLocalState[nodeId] = LOCAL_NOTREADY
	graph.nodeIdToGlobalState[nodeId] = NOTREADY
	return nil
}

func (graph *DepGraph) SetNodeStateLocalNotReady(nodeId string) {
	graph.nodeIdToLocalState[nodeId] = LOCAL_NOTREADY
}
func (graph *DepGraph) SetNodeStateLocalReady(nodeId string) {
	graph.nodeIdToLocalState[nodeId] = READY
}
func (graph *DepGraph) UpdateNodeState(nodeId string) {
	nodeState, _ := graph.nodeIdToLocalState[nodeId]
	node, _ := graph.acyclicGraph.GetNode(nodeId)
	if nodeState == LOCAL_NOTREADY {
		graph.nodeIdToGlobalState[nodeId] = NOTREADY
		// update parents to be not ready
		parents := node.GetParents()
		fmt.Println("node marked not ready")
		fmt.Println("node has ", len(parents), " parents")
	}else if nodeState == READY {
		dependencies := node.GetDependencies()
		fmt.Println("warning bypassing children check for now")
		fmt.Println("node marked ready")
		fmt.Println("node has ", len(dependencies), " dependencies")
	}
}
func (graph *DepGraph) SetNodeStateNotReady(nodeId string) error {
	// @todo
	// not ready, so we should traverse all parent nodes and mark them as global not ready
	// can come from any state
	graph.SetNodeStateLocalNotReady(nodeId)
	graph.UpdateNodeState(nodeId)
	return nil
}
func (graph *DepGraph) SetNodeStateReady(nodeId string) error {
	// @todo
	// node global ready, so we should simply mark this.  do not do anything since parents need completion
	// must come from not ready state
	graph.SetNodeStateLocalReady(nodeId)
	graph.UpdateNodeState(nodeId)

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
func (graph *DepGraph) GetNodeGlobalReadiness(nodeId string) (State, error) {
	nodeState, nodeOk := graph.nodeIdToGlobalState[nodeId]
	if !nodeOk {
		return 0, errors.New("invalid node")
	}
	return nodeState,nil
}

