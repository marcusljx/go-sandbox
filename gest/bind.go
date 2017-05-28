package gest

import "testing"

// bindTests wraps the Tests with the suite's setUpTestFunc and tearDownTestFunc functions
func (s *Suite) bindTests(tests []testing.InternalTest) (bindedTests []testing.InternalTest) {
	for _, c := range tests {
		newC := c // make copy
		newC.F = func(t *testing.T) {
			s.setUpTestFunc(t)
			defer s.tearDownTestFunc(t)
			c.F(t)
		}
		bindedTests = append(bindedTests, newC)
	}
	return
}

// bindBenchmarks wraps the Benchmarks with the suite's setUpBenchmarkFunc and tearDownBenchmarkFunc functions
func (s *Suite) bindBenchmarks(benchmarks []testing.InternalBenchmark) (bindedBenchmarks []testing.InternalBenchmark) {
	for _, c := range benchmarks {
		newC := c // make copy
		newC.F = func(b *testing.B) {
			s.setUpBenchmarkFunc(b)
			defer s.tearDownBenchmarkFunc(b)
			c.F(b)
		}
		bindedBenchmarks = append(bindedBenchmarks, newC)
	}
	return
}

// bindExamples wraps the Examples with the suite's setUpExampleFunc and tearDownExampleFunc functions
func (s *Suite) bindExamples(examples []testing.InternalExample) (bindedExamples []testing.InternalExample) {
	for _, c := range examples {
		newC := c // make copy
		newC.F = func() {
			s.setUpExampleFunc()
			defer s.tearDownExampleFunc()
			c.F()
		}
		bindedExamples = append(bindedExamples, newC)
	}
	return
}
