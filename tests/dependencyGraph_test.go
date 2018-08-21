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
	readinessStorkAuto, readyErrStorkAuto := graph.GetNodeGlobalReadiness("stork-automate")
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

	readinessStorkAuto, readyErrStorkAuto := graph.GetNodeGlobalReadiness("stork-automate")
	if readyErrStorkAuto != nil {
		test.Error(nil)
	}
	if readinessStorkAuto != NOTREADY {
		test.Error("expected stork-automate to be not ready")
	}

	readinessStork, _ := graph.GetNodeGlobalReadiness("stork")
	if readinessStork != NOTREADY {
		test.Error("expected stork to be not ready")
	}

	readinessAuto, _ := graph.GetNodeGlobalReadiness("automate")
	if readinessAuto != NOTREADY {
		test.Error("expected automate to be not ready")
	}
}
func TestSetOneNodeReady(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")

	readinessStorkAuto, _ := graph.GetNodeGlobalReadiness("stork")
	if readinessStorkAuto != NOTREADY {
		test.Error("expected stork to be not ready before set")
	}

	graph.SetNodeStateReady("stork")
	readinessStorkAfterSet, _ := graph.GetNodeGlobalReadiness("stork")

	if readinessStorkAfterSet != READY {
		test.Error("expected stork to be ready since set to ready and has no dependencies")
	}
}

func TestSetOneNodeReadyDepNotReady(test *testing.T){
	graph := New()
	graph.AddDependency("stork-automate", "stork")

	readinessStorkAuto, _ := graph.GetNodeGlobalReadiness("stork-automate")
	if readinessStorkAuto != NOTREADY {
		test.Error("expected stork-automate to be not ready before set")
	}

	graph.SetNodeStateReady("stork-automate")
	readinessStorkAfterSet, _ := graph.GetNodeGlobalReadiness("stork-automate")

	if readinessStorkAfterSet != NOTREADY {
		test.Error("expected stork-automate to be not ready since set to ready but has not ready dependencies")
	}
}



