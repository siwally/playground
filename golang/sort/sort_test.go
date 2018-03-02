package sort

import "testing"
import "fmt"
import "log"

func ExampleSort() {

	res, err := Sort([]int{22, 5, 11, 7, 1, 15, 3, -111, 30, 42, 0, 2})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", res)
	// Output: [-111 0 1 2 3 5 7 11 15 22 30 42]
}

func TestNil(t *testing.T) {
	testScenario(t, nil, nil)
}

func TestSingleValue(t *testing.T) {
	testScenario(t, []int{22}, []int{22})
}

func TestWithTwoValuesUnordered(t *testing.T) {
	testScenario(t, []int{22, 11}, []int{11, 22})
}

func TestWithTwoValuesAlreadyOrdered(t *testing.T) {
	testScenario(t, []int{11, 22}, []int{11, 22})
}

func TestSortOddLength(t *testing.T) {
	testScenario(t, []int{22, 11, 5},
		[]int{5, 11, 22})
}

func TestSortEvenLength(t *testing.T) {
	testScenario(t, []int{22, 11, 5, 1},
		[]int{1, 5, 11, 22})
}

func TestSortAlreadyOrdered(t *testing.T) {
	testScenario(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func testScenario(t *testing.T, input, expected []int) {
	res, err := Sort(input)

	if err != nil {
		t.Errorf("Error returned when sorting %v", input)
	}

	if !matches(res, expected) {
		t.Errorf("Expected %v but was %v", expected, res)
	}
}
func matches(res []int, exp []int) bool {
	if len(res) != len(exp) {
		return false
	}

	for i, v := range res {
		if v != exp[i] {
			return false
		}
	}

	return true
}

func BenchmarkSort(b *testing.B) {
	in := [1000000]int{}

	for i, v := 0, len(in); i < len(in); i++ {
		in[i] = v
		v--
	}

	algInUse = seqAlg{}

	for i := 0; i < b.N; i++ {
		Sort(in[0:])
	}
}

func BenchmarkSortParallel(b *testing.B) {
	in := [1000000]int{}

	for i, v := 0, len(in); i < len(in); i++ {
		in[i] = v
		v--
	}

	algInUse = parAlg{}

	for i := 0; i < b.N; i++ {
		Sort(in[0:])
	}
}
