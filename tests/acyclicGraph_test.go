package tests

import (
	"strconv"
	"testing"
)
import "../src/mainlogic/dependencyGraph/acyclicGraph"

func TestBasicAddDependencyOneTarget(test *testing.T) {
	graph := acyclicGraph.New()
	err := graph.AddDependency("stork-automate", "stork")
	if err != nil {
		test.Error(err)
	}
}
func TestBasicGetDependenciesOneTarget(test *testing.T) {
	graph := acyclicGraph.New()
	graph.AddDependency("stork-automate", "stork")
	graph.AddDependency("stork-automate", "test")
	dependencies := graph.GetDependencies("stork-automate")
	if len(dependencies) != 2 {
		test.Error("expected 2 dependency, got " + strconv.Itoa(len(dependencies)))
	}
}
func TestBasicGetNumParentOneTarget(test *testing.T){
	graph := acyclicGraph.New()
	graph.AddDependency("stork-automate", "stork")
	graph.AddDependency("stork-automate", "test")

	numStorkAuto, _ := graph.GetNumImmediateParents("stork-automate")
	if numStorkAuto != 0 {
		test.Error("expected stork-automate to have no parents")
	}

	numStork, _ := graph.GetNumImmediateParents("stork")
	if numStork != 1 {
		test.Error("expected stork to have one parent")
	}
}
func TestBasicGetNumParentNilTarget(test *testing.T){
	graph := acyclicGraph.New()
	graph.AddDependency("stork-automate", "stork")
	graph.AddDependency("stork-automate", "test")

	_, err := graph.GetNumImmediateParents("cat")
	if err == nil {
		test.Error("expected error for nil target")
	}

	
}
func TestBasicGetDependenciesTwoTargetsNoSharedDependency(test *testing.T) {
	graph := acyclicGraph.New()
	graph.AddDependency("stork-automate", "stork")
	graph.AddDependency("stork-automate", "test")

	graph.AddDependency("automate-beaver", "beaver")
	graph.AddDependency("automate-beaver", "when-do")
	graph.AddDependency("automate-beaver", "when-thing")

	dependenciesSA := graph.GetDependencies("stork-automate")
	dependenciesAB := graph.GetDependencies("automate-beaver")

	if len(dependenciesSA) != 2 {
		test.Error("expected 2 dependencies for stork-automate got " + strconv.Itoa(len(dependenciesSA)))
	}
	if len(dependenciesAB) != 3 {
		test.Error("expected 3 dependencies for automate-beaver got " + strconv.Itoa(len(dependenciesAB)))

	}
}
func TestBasicGetDependenciesTwoTargetsSharedDependency(test *testing.T) {
	graph := acyclicGraph.New()
	graph.AddDependency("stork-automate", "stork")
	graph.AddDependency("stork-automate", "test")

	graph.AddDependency("automate-beaver", "beaver")
	graph.AddDependency("automate-beaver", "when-do")
	graph.AddDependency("automate-beaver", "stork")

	dependenciesSA := graph.GetDependencies("stork-automate")
	dependenciesAB := graph.GetDependencies("automate-beaver")

	if len(dependenciesSA) != 2 {
		test.Error("expected 2 dependencies for stork-automate got " + strconv.Itoa(len(dependenciesSA)))
	}
	if len(dependenciesAB) != 3 {
		test.Error("expected 3 dependencies for automate-beaver got " + strconv.Itoa(len(dependenciesAB)))

	}
}

func TestComplexAddDependency(test *testing.T) {
	graph := acyclicGraph.New()
	err1 := graph.AddDependency("stork-automate", "stork")
	err2 := graph.AddDependency("automate", "scheduler")
	err3 := graph.AddDependency("automate", "logic")
	err4 := graph.AddDependency("automate", "cron")
	err5 := graph.AddDependency("stork", "automate")
	err6 := graph.AddDependency("scheduler", "cron")
	if err1 != nil {
		test.Error(err1)
	}
	if err2 != nil {
		test.Error(err2)
	}
	if err3 != nil {
		test.Error(err3)
	}
	if err4 != nil {
		test.Error(err4)
	}
	if err5 != nil {
		test.Error(err5)
	}
	if err6 != nil {
		test.Error(err6)
	}
}

func TestBasicCircularDependency(test *testing.T) {
	graph := acyclicGraph.New()
	graph.AddDependency("stork-automate", "stork")
	err := graph.AddDependency("stork", "stork-automate")
	if err == nil {
		test.Error("did not error when expecting circular dependency error")
	}
}

//          stork-automate
func TestComplexCircularDependency(test *testing.T) {
	graph := acyclicGraph.New()
	err1 := graph.AddDependency("stork-automate", "stork")
	err2 := graph.AddDependency("automate", "scheduler")
	err3 := graph.AddDependency("automate", "logic")
	err4 := graph.AddDependency("automate", "cron")
	err5 := graph.AddDependency("stork", "automate")
	err6 := graph.AddDependency("scheduler", "stork")
	if err1 != nil {
		test.Error(err1)
	}
	if err2 != nil {
		test.Error(err2)
	}
	if err3 != nil {
		test.Error(err3)
	}
	if err4 != nil {
		test.Error(err4)
	}
	if err5 != nil {
		test.Error(err5)
	}
	if err6 == nil {
		test.Error("Expected circular dependency error, stork depends on automate which depends on scheduler, so cannot make scheduler depend on stork")
	}
}
