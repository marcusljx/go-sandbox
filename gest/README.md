# gest
SetUp + TearDown for standard go tests. Allows for before-after wrapping of test cases without having to move away from standard golang test development (*a la* `testify Suite`)

## Usage
Apply the following in your test package's `TestMain`:
```
func TestMain(m *testing.M) {
  code := gest.New("MyTestSuite").
    TestSetUp(setUpFunc).
    TestTearDown(tearDownFunc).
    Run(m)
    
  os.Exit(code)
}

func setUpFunc(t *testing.T) {
  // Do set up here...
}

func tearDownFunc(t *testing.T) {
  // Do tear down here...
}
```

## Supported Builders
Each new `gest` Suite supports the following:
- [x] TestSetUp
- [x] TestTearDown
- [x] BenchmarkSetUp (runs on every iteration, not a general "before whole case") `// See TODO`
- [x] BenchmarkTearDown (runs on every iteration, not a general "after whole case") `// See TODO`
- [x] ExampleSetUp
- [x] ExampleTearDown
- [x] BenchmarkCaseSetUp
- [x] BenchmarkCaseTearDown

## `TODO`
- [ ] Test/Benchmark/Example Function name injection into setup/teardown methods (breaks API)
