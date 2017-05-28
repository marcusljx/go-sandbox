package gest

import (
	"os"
	"testing"

	"log"

	"fmt"

	"github.com/stretchr/testify/require"
)

const (
	setupMsg    = "SETUP WAS HIT"
	testMsg     = "TEST WAS HIT"
	TearDownMsg = "TEARDOWN WAS HIT"
)

var (
	testSignature      = ""
	benchmarkSignature = 0
	exampleSignature   = ""
)

func TestMain(m *testing.M) {
	code := New("TestTest").
		TestSetUp(testSetUp).
		TestTearDown(testTearDown).
		BenchmarkSetUp(benchmarkSetUp).
		BenchmarkTearDown(benchmarkTearDown).
		ExampleSetUp(exampleSetUp).
		ExampleTearDown(exampleTearDown).
		Run(m)

	//verify teardown msgs
	if testSignature != "" {
		if testSignature != TearDownMsg {
			log.Printf("TestTearDown failed to bind : expected [%v] but got [%v]", TearDownMsg, testSignature)
			code--
		}
	}
	if benchmarkSignature != 0 {
		if benchmarkSignature%3 != 0 {
			log.Printf("BenchmarkTearDown failed to bind : expected [%v] but got [%v]", 0, benchmarkSignature%3)
			code--
		}
	}
	if exampleSignature != "" {
		if exampleSignature != TearDownMsg {
			log.Printf("ExampleTearDown failed to bind : expected [%v] but got [%v]", TearDownMsg, exampleSignature)
			code--
		}
	}

	os.Exit(code)
}

//===================================================
func TestSample(t *testing.T) {
	require.Equal(t, setupMsg, testSignature)
	testSignature = testMsg
}

func testSetUp(t *testing.T) {
	require.Equal(t, "", testSignature)
	testSignature = setupMsg
}

func testTearDown(t *testing.T) {
	require.Equal(t, testMsg, testSignature)
	testSignature = TearDownMsg
}

//===================================================
func BenchmarkSample(b *testing.B) {
	require.True(b, (benchmarkSignature-1)%3 == 0, "BenchmarkSample was not correct : [%v]", benchmarkSignature)
	benchmarkSignature++
}

func benchmarkSetUp(b *testing.B) {
	require.True(b, benchmarkSignature%3 == 0, "BenchmarkSample was not correct : [%v]", benchmarkSignature)
	benchmarkSignature++
}

func benchmarkTearDown(b *testing.B) {
	require.True(b, (benchmarkSignature-2)%3 == 0, "BenchmarkSample was not correct : [%v]", benchmarkSignature)
	benchmarkSignature++
}

//===================================================
func ExampleSample() {
	exampleSignature = testMsg
	fmt.Printf("During Test : %s\n", exampleSignature)
	// Output:
	// During SetUp : SETUP WAS HIT
	// During Test : TEST WAS HIT
	// During TearDown : TEARDOWN WAS HIT
}

func exampleSetUp() {
	exampleSignature = setupMsg
	fmt.Printf("During SetUp : %s\n", exampleSignature)
}

func exampleTearDown() {
	exampleSignature = TearDownMsg
	fmt.Printf("During TearDown : %s\n", exampleSignature)
}
