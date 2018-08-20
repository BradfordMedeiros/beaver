package tests

import (
	"testing"
)
import "../src/mainlogic/dependencyGraph"

func DepGraphTestBasic(test *testing.T) {
	graph := dependencyGraph.New()
	err := graph.AddDependency("stork-automate", "stork")
	if err != nil {
		test.Error(err)
	}
}
