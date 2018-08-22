package tests

import (
	"testing"
)
import . "../src/mainlogic/dependencyGraph"

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



