package tuple

import (
	"fmt"
	"testing"
)

func assertEq(t *testing.T, lhs interface{}, rhs interface{}) {
	if lhs != rhs {
		t.Error(lhs, " != ", rhs)
	}
}

func TestNew(t *testing.T) {
	tup := NewTuple(3)
	s := tup.String()

	fmt.Println("t is", s)
	x := tup.Len()
	assertEq(t, x, 3)

	/*
	   t.Set(0, "kevin")
	   t.Set(1, "thomas")
	   t.Set(2, "manley")
	*/

}
