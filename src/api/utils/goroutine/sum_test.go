package goroutine

import (
	"fmt"
	"testing"
)

func TestReturnSum(t *testing.T) {
	x, y := returnSum()
	fmt.Println(x, y)
}
