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
	t1 := NewTuple(3)
	x := t1.Len()
	assertEq(t, x, 3)

	t2 := NewTuple(3)
	assertEq(t, t1.Len(), t2.Len())

	t3 := NewTuple(3, 1, 2, 3)
	t4 := NewTuple(3, 1, 2, 3)
	if !t3.Eq(t4) {
		t.Error(t3.String(), "!=", t4.String())
	}
	if t3.Ne(t4) {
		t.Error(t3.String(), "!=", t4.String())
	}

	/*
	   t.Set(0, "kevin")
	   t.Set(1, "thomas")
	   t.Set(2, "manley")
	*/

	//z := t3.Get(2).(int)
	//fmt.Println(z * 3)

	fmt.Println(t3.String())
	fmt.Println(t3.PopLeft())
	fmt.Println(t3)
	fmt.Println(t3.PopLeft())
	fmt.Println(t3)
	fmt.Println(t3.PopLeft())
	fmt.Println(t3)
	fmt.Println(t3.PopLeft())
	fmt.Println(t3)

}

func TestString(t *testing.T) {
	tup := NewTuple(3)
	s := tup.String()
	assertEq(t, s, "[<nil> <nil> <nil>]")
}
