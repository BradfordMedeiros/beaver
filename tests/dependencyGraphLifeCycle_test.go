package tests

import (
	"testing"
)
import . "../src/mainlogic/dependencyGraph"

func TestSingleNodeLifecycle(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")
	readinessStork1, _ := graph.GetNodeGlobalState("stork")
	if readinessStork1 != NOTREADY {
		test.Error("expected stork to be not ready")
	}

	graph.SetNodeStateLocalReady("stork")
	readinessStork2, _ := graph.GetNodeGlobalState("stork")
	if readinessStork2 != READY {
		test.Error("expected stork to be ready")
	}

	graph.AdvanceNodeStateQueued("stork")
	readinessStork3, _ := graph.GetNodeGlobalState("stork")
	if readinessStork3 != QUEUED {
		test.Error("expected stork to be queued")
	}

	graph.AdvanceNodeStateInProgress("stork")
	readinessStork4, _ := graph.GetNodeGlobalState("stork")
	if readinessStork4 != INPROGRESS{
				test.Error("expected stork to be in progress")
	}
	graph.AdvanceNodeStateComplete("stork")
	readinessStork5, _ := graph.GetNodeGlobalState("stork")
	if readinessStork5 != COMPLETE {
				test.Error("expected stork to be complete")
	}
}

func TestLifeCycle_AdvanceQueuedUnlessGlobalReady_NoDeps(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")
	readinessStork1, _ := graph.GetNodeGlobalState("stork")
	if readinessStork1 != NOTREADY {
		test.Error("expected stork to be not ready")
	}

	err := graph.AdvanceNodeStateQueued("stork")
	if err == nil {
		test.Error("expected an error because advanced a non-ready node to queue, got nil")
	}

	graph.SetNodeStateLocalReady("stork")
	err1 := graph.AdvanceNodeStateQueued("stork")
	if err1 != nil {
		test.Error("expected an advance queue to work now since local ready")
	}
	
	finalState, _ := graph.GetNodeGlobalState("stork")
	if finalState != QUEUED {
		test.Error("node not queued")
	}
}

func TestLifeCycle_AdvanceInProgressUnlessQueued_NoDeps(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")
	err := graph.AdvanceNodeStateInProgress("stork")
	if err == nil {
		test.Error("expected an error because advanced a non-ready node to in progress")
	}
	graph.SetNodeStateLocalReady("stork")
	err1 := graph.AdvanceNodeStateInProgress("stork")
	if err1 == nil {
		test.Error("expected an error because advanced a non-queued node to in progress")
	}
	graph.AdvanceNodeStateQueued("stork")
	err2 := graph.AdvanceNodeStateInProgress("stork")
	if err2 != nil {
		test.Error("expected to be able to advance to in progress since already queued")
	}
	finalState, _ := graph.GetNodeGlobalState("stork")
	if finalState != INPROGRESS {
		test.Error("node not in progress")
	}
}

func TestLifeCycle_AdvanceCompleteUnlessInProgess_NoDeps(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")
	err := graph.AdvanceNodeStateComplete("stork")
	if err == nil {
		test.Error("expected an error because advanced a non-ready node to in progress")
	}
	graph.SetNodeStateLocalReady("stork")
	err1 := graph.AdvanceNodeStateComplete("stork")
	if err1 == nil {
		test.Error("expected an error because advanced a non-queued node to in progress")
	}
	graph.AdvanceNodeStateQueued("stork")
	err2 := graph.AdvanceNodeStateInProgress("stork")
	if err2 != nil {
		test.Error("expected to be able to advance to in progress since already queued")
	}
	graph.AdvanceNodeStateInProgress("stork")
	err3 := graph.AdvanceNodeStateComplete("stork")
	if err3 != nil {
		test.Error("expected to be able to advance to complete since in progress")
	}
	finalState, _ := graph.GetNodeGlobalState("stork")
	if finalState != COMPLETE {
		test.Error("node not complete")
	}
}