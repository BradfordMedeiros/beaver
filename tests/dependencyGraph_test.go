package tests

import (
	"testing"
)
import . "../src/mainlogic/dependencyGraph"

func TestAddSingleNodeNoDependency(test *testing.T){
	graph := New()
	err := graph.AddNode("automate")
	if err != nil {
		test.Error(err)
	}

	readinessAuto, errAuto := graph.GetNodeGlobalState("automate")
	if errAuto != nil {
		test.Error(errAuto)
	}
	if readinessAuto != NOTREADY {
		test.Error("expected automate to be not ready")
	}

	graph.SetNodeStateLocalReady("automate")

	readinessAuto1, _:= graph.GetNodeGlobalState("automate")
	if readinessAuto1 != READY {
		test.Error("expected automate to be ready")
	}

}

func TestInitialStateOneNode(test *testing.T) {
	graph := New()
	err := graph.AddDependency("stork-automate", "stork")
	if err != nil {
		test.Error(err)
	}
	readinessStorkAuto, readyErrStorkAuto := graph.GetNodeGlobalState("stork-automate")
	if readyErrStorkAuto != nil {
		test.Error(nil)
	}
	if readinessStorkAuto != NOTREADY {
		test.Error("expected stork-automate to be not ready")
	}
}
func TestInitialStateMultipleNodes(test *testing.T) {
	graph := New()
	err := graph.AddDependency("stork-automate", "stork")
	err2 := graph.AddDependency("stork-automate", "automate")
	if err != nil {
		test.Error(err)
	}
	if err2 != nil {
		test.Error(err2)
	}

	readinessStorkAuto, readyErrStorkAuto := graph.GetNodeGlobalState("stork-automate")
	if readyErrStorkAuto != nil {
		test.Error(nil)
	}
	if readinessStorkAuto != NOTREADY {
		test.Error("expected stork-automate to be not ready")
	}

	readinessStork, _ := graph.GetNodeGlobalState("stork")
	if readinessStork != NOTREADY {
		test.Error("expected stork to be not ready")
	}

	readinessAuto, _ := graph.GetNodeGlobalState("automate")
	if readinessAuto != NOTREADY {
		test.Error("expected automate to be not ready")
	}
}
func TestSetOneNodeReady(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")

	readinessStorkAuto, _ := graph.GetNodeGlobalState("stork")
	if readinessStorkAuto != NOTREADY {
		test.Error("expected stork to be not ready before set")
	}

	graph.SetNodeStateLocalReady("stork")
	readinessStorkAfterSet, _ := graph.GetNodeGlobalState("stork")

	if readinessStorkAfterSet != READY {
		test.Error("expected stork to be ready since set to ready and has no dependencies")
	}
}

func TestSetOneNodeReadyDepNotReady(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")

	readinessStorkAuto, _ := graph.GetNodeGlobalState("stork-automate")
	if readinessStorkAuto != NOTREADY {
		test.Error("expected stork-automate to be not ready before set")
	}

	graph.SetNodeStateLocalReady("stork-automate")
	readinessStorkAfterSet, _ := graph.GetNodeGlobalState("stork-automate")

	if readinessStorkAfterSet != NOTREADY {
		test.Error("expected stork-automate to be not ready since set to ready but has not ready dependencies")
	}
}

func TestSetBasicDependency(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")
	graph.AddDependency("stork-automate", "automate")

	graph.SetNodeStateLocalReady("stork")
	graph.SetNodeStateLocalReady("stork-automate")
	graph.SetNodeStateLocalReady("automate")


	storkState1, _ := graph.GetNodeGlobalState("stork")
	autoState1, _ := graph.GetNodeGlobalState("automate")
	storkAutoState1, _ := graph.GetNodeGlobalState("stork-automate")

	if storkState1 != READY {
		test.Error("expected stork to be ready")
	}
	if autoState1 != READY {
		test.Error("expected automate to be ready")
	}
	if storkAutoState1 != NOTREADY {
		test.Error("expected stork-automate to be not ready")
	}

	storkState2, _ := graph.GetNodeGlobalState("stork")
	autoState2, _ := graph.GetNodeGlobalState("automate")
	storkAutoState2, _ := graph.GetNodeGlobalState("stork-automate")
	if storkState2 != READY {
		test.Error("expected stork to be ready")
	}
	if autoState2 != READY {
		test.Error("expected automate to be ready")
	}
	if storkAutoState2 != NOTREADY {
		test.Error("expected stork-automate to be not ready")
	}

	graph.AdvanceNodeStateQueued("stork")
	graph.AdvanceNodeStateInProgress("stork")
	graph.AdvanceNodeStateComplete("stork")
	storkAutoState3, _ := graph.GetNodeGlobalState("stork-automate")
	if storkAutoState3 != NOTREADY {
		test.Error("expected stork-automate to be not ready")
	}

	graph.AdvanceNodeStateQueued("automate")
	graph.AdvanceNodeStateInProgress("automate")
	graph.AdvanceNodeStateComplete("automate")
	storkAutoState4, _ := graph.GetNodeGlobalState("stork-automate")
	if storkAutoState4 != READY {
		test.Error("expected stork-automate to be ready, got ")
	}
}

func TestSetNotDependency(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")
	graph.AddDependency("stork-automate", "automate")

	graph.SetNodeStateLocalReady("stork")
	graph.SetNodeStateLocalReady("stork-automate")
	graph.SetNodeStateLocalReady("automate")


	graph.AdvanceNodeStateQueued("stork")
	graph.AdvanceNodeStateInProgress("stork")
	graph.AdvanceNodeStateComplete("stork")
	graph.AdvanceNodeStateQueued("automate")
	graph.AdvanceNodeStateInProgress("automate")
	graph.AdvanceNodeStateComplete("automate")
	storkAutoState4, _ := graph.GetNodeGlobalState("stork-automate")
	if storkAutoState4 != READY {
		test.Error("expected stork-automate to be ready, got ")
	}
	graph.SetNodeStateLocalNotReady("automate")
	automateState, _ := graph.GetNodeGlobalState("automate")
	if automateState != NOTREADY {
		test.Error("expected global automate state to be not ready since set local not ready got ", automateState)
	} 

	storkAutoState5, _ := graph.GetNodeGlobalState("stork-automate")
	if storkAutoState5 != NOTREADY {
		test.Error("expected stork-automate to be not ready, got ", storkAutoState5)
	}
}



