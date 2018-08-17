package tests

import (
	"strconv"
	"testing"
	"fmt"
)
import "../src/mainlogic/dependencyGraph"

func TestBasicAddDependencyOneTarget(test *testing.T) {
	graph := dependencyGraph.New()
	err := graph.AddDependency("stork-automate", "stork")
	if err != nil {
		test.Error(err)
	}
}
func TestBasicGetDependenciesOneTarget(test *testing.T) {
	graph := dependencyGraph.New()
	graph.AddDependency("stork-automate", "stork")
	graph.AddDependency("stork-automate", "test")
	dependencies := graph.GetDependencies("stork-automate")
	if len(dependencies) != 2 {
		test.Error("expected 2 dependency, got " + strconv.Itoa(len(dependencies)))
	}
}
func TestBasicGetDependenciesTwoTargetsNoSharedDependency(test *testing.T) {
	graph := dependencyGraph.New()
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
	graph := dependencyGraph.New()
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

func TestComplexGetDependencies(test *testing.T) {

}

func TestComplexAddDependency(test *testing.T) {
	graph := dependencyGraph.New()
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
	graph := dependencyGraph.New()
	graph.AddDependency("stork-automate", "stork")
	err := graph.AddDependency("stork", "stork-automate")
	if err == nil {
		test.Error("did not error when expecting circular dependency error")
	}
}

//          stork-automate
func TestComplexCircularDependency(test *testing.T) { 
	fmt.Println("failing test")
	graph := dependencyGraph.New()                             //         stork
	err1 := graph.AddDependency("stork-automate", "stork")     //         /						 //  stork-automate
	err2 := graph.AddDependency("stork", "scheduler")       //	 	automate    
	graph.AddDependency("scheduler", "wow")

	fmt.Println("---------------------------")             
	err3 := graph.AddDependency("scheduler", "stork-automate") //   -------------circular to stork-atuomate------------------
	fmt.Println("========================")

	if err1 != nil {
		test.Error(err1)
	}
	if err2 != nil {
		test.Error(err2)
	}
	if err3 == nil {
		test.Error("expected error due to circular dependency got nil: ", err3)
	}
}
