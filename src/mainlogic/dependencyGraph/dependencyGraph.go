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
	onStateChange func(nodeId string, newState GlobalState);
}

func New(onStateChange func(nodeId string, newState GlobalState)) *DepGraph  {
	graph := &DepGraph { 
		acyclicGraph: acyclicGraph.New(), 
		nodeIdToLocalState: make(map[string]State), 
		nodeIdToGlobalState: make(map[string]GlobalState),
		onStateChange: onStateChange,
	}
	return graph
}

func (graph *DepGraph) AddNode(nodeId string) error {
	err := graph.acyclicGraph.AddNode(nodeId)
	if err !=nil {
		return err
	}
	graph.nodeIdToLocalState[nodeId] = LOCAL_NOTREADY
	graph.setNodeGlobalState(nodeId, NOTREADY)
	return nil
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

// helper function just to set the global state array value and trigger update
// all updates should go through this so we can message all updates externally 
func (graph *DepGraph) setNodeGlobalState(nodeId string, newState GlobalState){
	graph.nodeIdToGlobalState[nodeId] = newState
	graph.onStateChange("state change here", newState)
}

// given a new graph, starting at a change at nodeId, traverse the graph to ensure the effects propogate
func (graph *DepGraph) UpdateNodeGlobalState(nodeId string){
	node, _ := graph.acyclicGraph.GetNode(nodeId)
	// get deps, check if all deps are complete and if we are local ready, if so mark this as ready
	// or if deps are not complete or this is not local ready, mark as not ready, recurse up

	allReady := true
	for _, dependency := range(node.GetDependencies()){
		if graph.nodeIdToGlobalState[dependency.NodeId] != COMPLETE {
			allReady = false
			break
		}	
	}

	globalNodeState, _  := graph.nodeIdToGlobalState[nodeId]
	if allReady && graph.nodeIdToLocalState[nodeId] == LOCAL_READY {
		if globalNodeState == NOTREADY { 		 // if we are not ready, become ready
			graph.setNodeGlobalState(nodeId, READY)
		}else if globalNodeState == READY {		 // if we are ready, just stay the same
			// do nothing
		}else if globalNodeState == QUEUED {	 // if we are queued, just stay the same ()
			// do nothing
		}else if globalNodeState == INPROGRESS { // if we are in progress, stay the same
			// do nothing
		}else if globalNodeState == COMPLETE {   // if we are complete, inform our parents
			// do nothing
		}else{
			panic("case not handled")
		}	
	}else if !allReady || graph.nodeIdToLocalState[nodeId] == LOCAL_NOTREADY {
		if globalNodeState == NOTREADY { 		// if we are not ready, stay not ready
			// do nothing
		}else if globalNodeState == READY {     // if we are ready, become not ready
			graph.setNodeGlobalState(nodeId, NOTREADY)
		}else if globalNodeState == QUEUED {    // if we are queued, just stay the same (we cannot remove a queued build)
			// do nothing
		}else if globalNodeState == INPROGRESS {  // if we are queued, just stay the same
			// do nothing
		}else if globalNodeState == COMPLETE {
			graph.setNodeGlobalState(nodeId, NOTREADY)
		}	
	}else{
		panic("unexpected case")
	}	
	for _, parent := range(node.GetParents()){
		graph.UpdateNodeGlobalState(parent.NodeId)
	}
}

func (graph *DepGraph) triggerStateChange(){
	fmt.Println("trigger state change placeholder: this signified whole graph updated")
}

// this function simply updates one nodes global standing, based upon its local state
func (graph *DepGraph) UpdateNodeState(nodeId string, localNodeState State) {
	
	node, _ := graph.acyclicGraph.GetNode(nodeId)

	// update local ready, so make it global ready if all dependencies complete (or has none)
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
			graph.setNodeGlobalState(nodeId, READY)
		}
	}else if localNodeState == LOCAL_NOTREADY {	
		// make not ready, so we update the global state only if we were in ready position
		// this would only affect parent nodes if we transition from complete, which notready won't trasition from
		graph.nodeIdToLocalState[nodeId] = LOCAL_NOTREADY
		globalNodeState, _ := graph.nodeIdToGlobalState[nodeId]
		if globalNodeState == READY || globalNodeState == COMPLETE  {
			graph.setNodeGlobalState(nodeId, NOTREADY)
		}else if globalNodeState == NOTREADY {
			graph.setNodeGlobalState(nodeId, NOTREADY)
		}else{
			fmt.Println("curr node state: ", globalNodeState)
			panic("not yet implemented behavior of setting a node local not ready besides from basic ready or complete state")
		}
	}else{
		// this shouldn't ever be reached
		// probably could just use better types to avoid
		panic("got state besides local not ready and ready")
	}
}
func (graph *DepGraph) SetNodeStateLocalNotReady(nodeId string) {
	graph.UpdateNodeState(nodeId, LOCAL_NOTREADY)
	graph.UpdateNodeGlobalState(nodeId)
	graph.triggerStateChange()
}
func (graph *DepGraph) SetNodeStateLocalReady(nodeId string) {
	graph.UpdateNodeState(nodeId, LOCAL_READY)
	graph.UpdateNodeGlobalState(nodeId)
	graph.triggerStateChange()
}
func (graph *DepGraph) AdvanceNodeStateQueued(nodeId string) error{
	nodeState, hasNode := graph.nodeIdToGlobalState[nodeId]
	if !hasNode  {
		return errors.New("node does not exist")
	}
	if nodeState != READY {
		return errors.New("nodeState advanced to queued, but node was not ready")
	}
	graph.setNodeGlobalState(nodeId, QUEUED)
	graph.UpdateNodeGlobalState(nodeId)
	graph.triggerStateChange()

	return nil
}
func (graph *DepGraph) AdvanceNodeStateInProgress(nodeId string) error{
	nodeState, hasNode := graph.nodeIdToGlobalState[nodeId]
	if !hasNode  {
		return errors.New("node does not exist")
	}
	if nodeState != QUEUED {
		return errors.New("nodeState advanced to progress, but node was not queued")
	}
	graph.setNodeGlobalState(nodeId, INPROGRESS)
	graph.UpdateNodeGlobalState(nodeId)
	graph.triggerStateChange()

	return nil
}
// on complete, then need to advance
func (graph *DepGraph) AdvanceNodeStateComplete(nodeId string) error {
	nodeState, hasNode := graph.nodeIdToGlobalState[nodeId]
	if !hasNode  {
		return errors.New("node does not exist")
	}
	if nodeState != INPROGRESS {
		return errors.New("nodeState advanced to complete, but node was not inprogress")
	}
	graph.setNodeGlobalState(nodeId, COMPLETE)
	graph.UpdateNodeGlobalState(nodeId)
	graph.triggerStateChange()

	return nil
}
func (graph *DepGraph) SetNodeStateError(nodeId string) error {
	panic("not yet implemented")
	return nil
}
func (graph *DepGraph) GetNodeGlobalState(nodeId string) (GlobalState, error) {
	nodeState, nodeOk := graph.nodeIdToGlobalState[nodeId]
	if !nodeOk {
		return 0, errors.New("invalid node")
	}
	return nodeState,nil
}

