package gest

import (
	"reflect"
	"testing"
)

// Suite is a builder type for wrapping test cases with SetUp and TearDown methods
type Suite struct {
	Name string `json:"name"`

	setUpTestFunc    func(*testing.T)
	tearDownTestFunc func(*testing.T)

	setUpBenchmarkFunc    func(*testing.B)
	tearDownBenchmarkFunc func(*testing.B)

	setUpExampleFunc    func()
	tearDownExampleFunc func()
}

// New begins a new Suite
func New(name string) *Suite {
	return &Suite{
		Name: name,
	}
}

// TestSetUp adds a Set Up method for Tests
func (s *Suite) TestSetUp(BeforeEachTestFunc func(*testing.T)) *Suite {
	s.setUpTestFunc = BeforeEachTestFunc
	return s
}

// TestTearDown adds a Tear Down method for Tests
func (s *Suite) TestTearDown(AfterEachTestFunc func(*testing.T)) *Suite {
	s.tearDownTestFunc = AfterEachTestFunc
	return s
}

// BenchmarkSetUp adds a Set Up method for Benchmarks
func (s *Suite) BenchmarkSetUp(BeforeEachBenchmarkFunc func(*testing.B)) *Suite {
	s.setUpBenchmarkFunc = BeforeEachBenchmarkFunc
	return s
}

// BenchmarkTearDown adds a Tear Down method for Benchmarks
func (s *Suite) BenchmarkTearDown(AfterEachBenchmarkFunc func(*testing.B)) *Suite {
	s.tearDownBenchmarkFunc = AfterEachBenchmarkFunc
	return s
}

// ExampleSetUp adds a Set Up method for Examples
func (s *Suite) ExampleSetUp(BeforeEachExampleFunc func()) *Suite {
	s.setUpExampleFunc = BeforeEachExampleFunc
	return s
}

// ExampleTearDown adds a Tear Down method for Examples
func (s *Suite) ExampleTearDown(AfterEachExampleFunc func()) *Suite {
	s.tearDownExampleFunc = AfterEachExampleFunc
	return s
}

// Run takes in a *testing.M object, instruments it according to the builder rules, and runs it.
// The return integer is the result of the instrumented M's .Run() method.
func (s *Suite) Run(m *testing.M) int {
	newM := s.instrument(m)
	return newM.Run()
}

// instrument instruments the testing.M object.
func (s *Suite) instrument(m *testing.M) *testing.M {
	v := reflect.Indirect(reflect.ValueOf(m))

	// extract m internal data
	tests := extractUnexportedField("tests", v).Interface().([]testing.InternalTest)
	benchmarks := extractUnexportedField("benchmarks", v).Interface().([]testing.InternalBenchmark)
	examples := extractUnexportedField("examples", v).Interface().([]testing.InternalExample)

	// bind all SetUp and TearDowns
	newTests := s.bindTests(tests)
	newBenchmarks := s.bindBenchmarks(benchmarks)
	newExamples := s.bindExamples(examples)

	return testing.MainStart(&TestDeps{}, newTests, newBenchmarks, newExamples)
}
