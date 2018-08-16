
package tests

import (
	"testing"
)
import "../src/mainlogic/dependencyGraph"

func TestAddDependency(test *testing.T) {
	graph := dependencyGraph.New()
	err := graph.AddDependency("stork-automate", "stork")
	if err != nil {
		test.Error(err)
	}
		
}
func TestBasicCircularDependency(test *testing.T){
	graph := dependencyGraph.New()
	graph.AddDependency("stork-automate", "stork")
	err := graph.AddDependency("stork", "stork-automate")
	if err == nil {
		test.Error("did not error when expecting circular dependency error")
	}
}

