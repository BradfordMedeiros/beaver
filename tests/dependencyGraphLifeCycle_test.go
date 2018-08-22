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
}

func TestLifeCycle_AdvanceInProgressUnlessInProgress_NoDeps{

}