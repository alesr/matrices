package matrices

import (
	"fmt"
	"testing"
)

var sideValuesTestCases = []struct {
	n                           int
	expectedLeft, expectedRight []int
}{
	{5, []int{0, 5, 10, 15, 20}, []int{4, 9, 14, 19, 24}},
	{2, []int{0, 2}, []int{1, 3}},
	{3, []int{0, 3, 6}, []int{2, 5, 8}},
}

var isSideValueTestCase = []struct {
	value, size                 int
	expectedLeft, expectedRight bool
}{
	{3, 3, true, false},
	{7, 4, false, true},
	{2, 2, true, false},
	{3, 2, false, true},
	{1, 3, false, false},
	{8, 4, true, false},
	{24, 5, false, true},
	{10, 5, true, false},
}

func TestSideValues(t *testing.T) {

	for _, test := range sideValuesTestCases {
		ls := GetLeftSide(test.n)
		rs := GetRightSide(test.n)

		fmt.Printf("n = %d\nFor left side got: %v\nExpecting: %v\nFor right, got: %v\nExpected: %v\n\n",
			test.n, ls, test.expectedLeft, rs, test.expectedRight)

	}
}

func TestIsSideValue(t *testing.T) {
	for _, test := range isSideValueTestCase {
		leftResult := IsLeftSide(test.value, test.size)
		rightResult := IsRightSide(test.value, test.size)

		if test.expectedLeft != leftResult || test.expectedRight != rightResult {
			t.Errorf("error")
		}
	}
}

func BenchmarkIsMatrixSide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range isSideValueTestCase {
			IsLeftSide(test.value, test.size)
			IsRightSide(test.value, test.size)
		}
	}
}
