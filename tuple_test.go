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

func TestNewTuple(t *testing.T) {
	tup := NewTuple(3)
	n := tup.Len()
	assertEq(t, n, 3)

	tup2 := NewTuple(0)
	n = tup2.Len()
	assertEq(t, n, 0)
}

func TestNewTupleFromSlice(t *testing.T) {
	slice := []interface{}{"a", "b", "c"}
	tup := NewTupleFromSlice(slice)
	assertEq(t, tup.Get(0), "a")
	assertEq(t, tup.Get(1), "b")
	assertEq(t, tup.Get(2), "c")
}

func TestNewTupleFromItems(t *testing.T) {
	tup := NewTupleFromItems(100, 200, 300, 400)
	assertEq(t, tup.Get(0), 100)
	assertEq(t, tup.Get(-1), 400)
}

func TestSlice(t *testing.T) {
	tup := NewTupleFromItems(3, 2, 1)
	assertEq(t, fmt.Sprintf("%x", tup.Slice()), fmt.Sprintf("%x", []int{3, 2, 1}))
}

func TestIndex(t *testing.T) {
	tup := NewTuple(10)
	assertEq(t, tup.Index(0), 0)
	assertEq(t, tup.Index(5), 5)
	assertEq(t, tup.Index(-1), 9)
	assertEq(t, tup.Index(-2), 8)
}

func TestSet(t *testing.T) {
	tup := NewTuple(5)
	tup.Set(0, 100)
	tup.Set(-1, 200)
	assertEq(t, tup.String(), NewTupleFromItems(100, nil, nil, nil, 200).String())
}

func TestGet(t *testing.T) {
	tup := NewTupleFromItems("t", "e", "s", "t", "!")
	assertEq(t, tup.Get(0), "t")
	assertEq(t, tup.Get(2), "s")
	assertEq(t, tup.Get(-1), "!")
}

func TestString(t *testing.T) {
	tup := NewTuple(3)
	s := tup.String()
	assertEq(t, s, "[<nil> <nil> <nil>]")

	tup2 := NewTupleFromItems(100, "abc", 200)
	s = tup2.String()
	assertEq(t, s, "[100 abc 200]")
}

func TestPopLeft(t *testing.T) {
	tup := NewTupleFromItems(2, 4, 6, 8)
	x := tup.PopLeft()
	assertEq(t, x, 2)
	assertEq(t, tup.Len(), 3)
	assertEq(t, tup.Eq(NewTupleFromItems(4, 6, 8)), true)
}

func TestPopRight(t *testing.T) {
	tup := NewTupleFromItems(1, 3, 5, 7)
	x := tup.PopRight()
	assertEq(t, x, 7)
	assertEq(t, tup.Len(), 3)
	assertEq(t, tup.Eq(NewTupleFromItems(1, 3, 5)), true)
}

func TestEq(t *testing.T) {
	tup1 := NewTupleFromItems(3, 6, 9)
	tup2 := NewTuple(3)
	tup2.Set(0, 3)
	tup2.Set(1, 6)
	tup2.Set(2, 9)
	assertEq(t, tup1.Eq(tup2), true)
	assertEq(t, tup1.Ne(tup2), false)
}
