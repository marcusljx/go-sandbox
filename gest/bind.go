package gest

import (
	"testing"
)

// bindTests wraps the Tests with the suite's setUpTestFunc and tearDownTestFunc functions
func (s *Suite) bindTests(tests []testing.InternalTest) (bindedTests []testing.InternalTest) {
	for _, c := range tests {
		// make copies1
		newC := c
		newF := c.F
		newName := c.Name

		// wrap test with setup and teardown
		newC.F = func(t *testing.T) {
			s.setUpTestFunc(t, newName)
			defer s.tearDownTestFunc(t, newName)
			newF(t)
		}
		bindedTests = append(bindedTests, newC)
	}
	return
}

// bindBenchmarks wraps the Benchmarks with the suite's setUpBenchmarkFunc and tearDownBenchmarkFunc functions
func (s *Suite) bindBenchmarks(benchmarks []testing.InternalBenchmark) (bindedBenchmarks []testing.InternalBenchmark) {
	for _, c := range benchmarks {
		// make copies1
		newC := c
		newF := c.F
		newName := c.Name

		newC.F = func(b *testing.B) {
			s.setUpBenchmarkFunc(b, newName)
			defer s.tearDownBenchmarkFunc(b, newName)
			newF(b)
		}
		bindedBenchmarks = append(bindedBenchmarks, newC)
	}
	return
}

// bindExamples wraps the Examples with the suite's setUpExampleFunc and tearDownExampleFunc functions
func (s *Suite) bindExamples(examples []testing.InternalExample) (bindedExamples []testing.InternalExample) {
	for _, c := range examples {
		// make copies1
		newC := c
		newF := c.F
		newName := c.Name

		newC.F = func() {
			s.setUpExampleFunc(newName)
			defer s.tearDownExampleFunc(newName)
			newF()
		}
		bindedExamples = append(bindedExamples, newC)
	}
	return
}
