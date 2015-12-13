package matrices

import "fmt"

// sides defines side values of a matrix nxn
type matrixSides struct {
	leftvalues, rightvalues []int
}

// SideValues prints out the left and right values of a matrix of size n
func SideValues(n int) ([]int, []int, error) {
	if n < 0 {
		return nil, nil,
			fmt.Errorf("Can't form a matrix with %d sides.", n)
	}
	m := &matrixSides{
		leftvalues:  GetLeftSide(n),
		rightvalues: GetRightSide(n),
	}
	return m.leftvalues, m.rightvalues, nil
}

// GetLeftSide returns the
// left column for a n x n.
func GetLeftSide(n int) []int {
	ls := make([]int, 0, n)

	for i := 0; i < n; i++ {
		ls = append(ls, i*n)
	}
	return ls
}

// GetRightSide returns the
// right column for a n x n.
func GetRightSide(n int) []int {
	rs := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		rs = append(rs, ((n * i) - 1))
	}
	return rs
}

// IsLeftSide returns true if i is on
// the left column on a matrix n x n.
func IsLeftSide(i, n int) bool {
	flag := false
	for j := 0; j < n; j++ {
		if n*j == i {
			flag = true
			break
		}
	}
	return flag
}

// IsRightSide returns true if i is on
// the right column on a matrix n x n.
func IsRightSide(i, n int) bool {
	flag := false
	for j := 1; j <= n; j++ {
		if (n*j)-1 == i {
			flag = true
			break
		}
	}
	return flag
}
