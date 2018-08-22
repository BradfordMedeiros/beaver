
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

type State int;
const (
	LOCAL_NOTREADY State = 0
	LOCAL_READY State = 1
)

type GlobalState int
const (
	NOTREADY  GlobalState = 0 // the plugin has not yet alerted as being ready for build (alert not called)
	READY     GlobalState = 1 // the plugin has called alert, is ready to  build (build has not yet been called)
	QUEUED 	  GlobalState = 2 // the scheduler has queued this item for buildqa
	INPROGRESS GlobalState = 3 // the schedule has invokved build, but has notyet called complete
	COMPLETE  GlobalState = 4 // the plugin has finished the build
	ERROR     GlobalState = 5 // the plugin has declared an error, this probably needs manual reset
)

type DepGraph struct {
	acyclicGraph *acyclicGraph.RootNode;
	nodeIdToLocalState map[string] State;
	nodeIdToGlobalState map[string] GlobalState;
}

func New() *DepGraph  {
	graph := &DepGraph { 
		acyclicGraph: acyclicGraph.New(), 
		nodeIdToLocalState: make(map[string]State), 
		nodeIdToGlobalState: make(map[string]GlobalState),
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

func (graph *DepGraph) UpdateNodeState(nodeId string, localNodeState State) {
	// we know which node changed, it either became not ready
	
	node, _ := graph.acyclicGraph.GetNode(nodeId)
	if localNodeState == LOCAL_READY {
		graph.nodeIdToLocalState[nodeId] = LOCAL_READY

		allReady := true
		for _, dependency := range(node.GetDependencies()){
			nodeState, _ := graph.nodeIdToGlobalState[dependency.NodeId]
			if nodeState != COMPLETE {
				allReady = false
				break
			}
		}
		if allReady {
			graph.nodeIdToGlobalState[nodeId] = READY
		}
	}else if localNodeState == LOCAL_NOTREADY {
		graph.nodeIdToLocalState[nodeId] = LOCAL_NOTREADY

	}else{
		// this shouldn't ever be reached
		// probably could just use better types to avoid
		panic("got state besides local not ready and ready")
	}

	/*nodeState, _ := graph.nodeIdToLocalState[nodeId]
	node, _ := graph.acyclicGraph.GetNode(nodeId)
	if nodeState == LOCAL_NOTREADY {
		graph.nodeIdToGlobalState[nodeId] = NOTREADY
		// update parents to be not ready
		parents := node.GetParents()
		for _, parent := range(parents){
			fmt.Println("updating parent: ", parent.NodeId)
		}

	}else if nodeState == LOCAL_READY {
		dependencies := node.GetDependencies()
		
		allReady := true
		for _, dependency := range(dependencies){
			nodeId := dependency.NodeId;
			nodeState, _ = graph.nodeIdToGlobalState[nodeId] // need to handle errors better here
			if nodeState != COMPLETE {
				allReady = false
			}
		}

		if allReady {
			graph.nodeIdToGlobalState[nodeId] = READY
		}
	}*/
}
func (graph *DepGraph) SetNodeStateLocalNotReady(nodeId string) {
	//graph.nodeIdToLocalState[nodeId] = LOCAL_NOTREADY
	fmt.Print("local not read")
	graph.UpdateNodeState(nodeId, LOCAL_NOTREADY)
}
func (graph *DepGraph) SetNodeStateLocalReady(nodeId string) {
	//graph.nodeIdToLocalState[nodeId] = LOCAL_READY
	graph.UpdateNodeState(nodeId, LOCAL_READY)
}
func (graph *DepGraph) setNodeStateNotReady(nodeId string) error {
	// @todo
	// not ready, so we should traverse all parent nodes and mark them as global not ready
	// can come from any state
	//graph.SetNodeStateLocalNotReady(nodeId)
	//graph.UpdateNodeState(nodeId)
	return nil
}
func (graph *DepGraph) AdvanceNodeStateQueued(nodeId string) error{
	// check if node was ready, if so advance it as queued, call a onqueue callback
	return nil
}
func (graph *DepGraph) AdvanceNodeStateInProgress(nodeId string) error{
	// must come from queued state
	// check if node was queued, if so advance as in progress, call callback
	return nil
}
// on complete, then need to advance
func (graph *DepGraph) AdvanceNodeStateComplete(nodeId string) error {
	// check if node was in progress, if so advance it
	// call update for parent nodes
	return nil
}
func (graph *DepGraph) SetNodeStateError(nodeId string) error {
	return nil
}
func (graph *DepGraph) GetNodeGlobalState(nodeId string) (GlobalState, error) {
	nodeState, nodeOk := graph.nodeIdToGlobalState[nodeId]
	if !nodeOk {
		return 0, errors.New("invalid node")
	}
	return nodeState,nil
}

