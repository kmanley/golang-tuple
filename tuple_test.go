package tuple

import (
	"fmt"
	"sort"
	"testing"
)

func assertEq(t *testing.T, lhs interface{}, rhs interface{}) {
	if lhs != rhs {
		t.Error(lhs, " != ", rhs)
	}
}

func assertEqTuple(t *testing.T, lhs *Tuple, rhs *Tuple) {
	if lhs.Ne(rhs) {
		t.Error(lhs, ".Ne(", rhs, ")")
	}
}

func assertNeTuple(t *testing.T, lhs *Tuple, rhs *Tuple) {
	if lhs.Eq(rhs) {
		t.Error(lhs, ".Eq(", rhs, ")")
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
	slice := []interface{}{"a", 1, "c"}
	tup := NewTupleFromSlice(slice)
	assertEq(t, tup.Get(0), "a")
	assertEq(t, tup.Get(1), 1)
	assertEq(t, tup.Get(2), "c")
}

func TestNewTupleFromItems(t *testing.T) {
	tup := NewTupleFromItems(100, 200, 300, 400)
	assertEq(t, tup.Get(0), 100)
	assertEq(t, tup.Get(-1), 400)
}

func TestCopy(t *testing.T) {
	t1 := NewTupleFromItems(1, 2, 3)
	t2 := t1.Copy()
	assertEqTuple(t, t1, t2)
	t1.Set(1, 20)
	assertEqTuple(t, t1, NewTupleFromItems(1, 20, 3))
	assertEqTuple(t, t2, NewTupleFromItems(1, 2, 3))
}

func TestData(t *testing.T) {
	tup := NewTupleFromItems(3, 2, 1)
	assertEq(t, fmt.Sprintf("%x", tup.Data()), fmt.Sprintf("%x", []int{3, 2, 1}))
}

func TestSlice(t *testing.T) {
	tup := NewTupleFromItems(100, 200, 300, 400)
	assertEqTuple(t, tup.Slice(0, 0), NewTuple(0))
	assertEqTuple(t, tup.Slice(0, 1), NewTupleFromItems(100))
	assertEqTuple(t, tup.Slice(0, 3), NewTupleFromItems(100, 200, 300))
	assertEqTuple(t, tup.Slice(0, 10), NewTupleFromItems(100, 200, 300, 400))
	assertEqTuple(t, tup.Slice(1, 2), NewTupleFromItems(200))
	assertEqTuple(t, tup.Slice(2, 4), NewTupleFromItems(300, 400))
	assertEqTuple(t, tup.Slice(4, 4), NewTuple(0))
	assertEqTuple(t, tup.Slice(10, 100), NewTuple(0))
	assertEqTuple(t, tup.Slice(-1, 100), NewTupleFromItems(400))
	assertEqTuple(t, tup.Slice(-3, -1), NewTupleFromItems(200, 300))
}

func TestLeft(t *testing.T) {
	tup := NewTupleFromItems(100, 200, 300, 400)
	assertEqTuple(t, tup.Left(0), NewTuple(0))
	assertEqTuple(t, tup.Left(1), NewTupleFromItems(100))
	assertEqTuple(t, tup.Left(3), NewTupleFromItems(100, 200, 300))
	assertEqTuple(t, tup.Left(10), NewTupleFromItems(100, 200, 300, 400))
}

func TestRight(t *testing.T) {
	tup := NewTupleFromItems(100, 200, 300, 400)
	assertEqTuple(t, tup.Right(0), NewTuple(0))
	assertEqTuple(t, tup.Right(1), NewTupleFromItems(400))
	assertEqTuple(t, tup.Right(3), NewTupleFromItems(200, 300, 400))
	assertEqTuple(t, tup.Right(10), NewTupleFromItems(100, 200, 300, 400))
}

func TestOffset(t *testing.T) {
	tup := NewTuple(10)
	assertEq(t, tup.Offset(0), 0)
	assertEq(t, tup.Offset(5), 5)
	assertEq(t, tup.Offset(-1), 9)
	assertEq(t, tup.Offset(-2), 8)
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
	assertEqTuple(t, tup, NewTupleFromItems(4, 6, 8))
}

func TestPopRight(t *testing.T) {
	tup := NewTupleFromItems(1, 3, 5, 7)
	x := tup.PopRight()
	assertEq(t, x, 7)
	assertEq(t, tup.Len(), 3)
	assertEqTuple(t, tup, NewTupleFromItems(1, 3, 5))
}

func TestEq(t *testing.T) {
	tup1 := NewTupleFromItems(3, 6, 9)
	tup2 := NewTuple(3)
	tup2.Set(0, 3)
	tup2.Set(1, 6)
	tup2.Set(2, 9)
	assertEqTuple(t, tup1, tup2)
	assertEq(t, tup1.Ne(tup2), false)

	tup5 := NewTupleFromItems(int16(100), int32(200), int64(300))
	tup6 := NewTupleFromItems(int8(100), int16(200), int32(300))
	assertEqTuple(t, tup5, tup6)

	tup7 := NewTupleFromItems(nil, nil, nil)
	tup8 := NewTuple(3)
	assertEqTuple(t, tup7, tup8)

	tup9 := NewTupleFromItems(NewTupleFromItems(1, 2), NewTupleFromItems(5, 10), NewTupleFromItems(10, 20))
	tup10 := NewTupleFromItems(NewTupleFromItems(1, 2), NewTupleFromItems(5, 10), NewTupleFromItems(10, 20))
	assertEqTuple(t, tup9, tup10)
	tup11 := NewTupleFromItems(NewTupleFromItems(1, 2), NewTupleFromItems(3, 10), NewTupleFromItems(10, 20))
	assertNeTuple(t, tup10, tup11)

}

func TestLt(t *testing.T) {
	tup1 := NewTupleFromItems(10, 20, 30)
	tup2 := NewTupleFromItems(10, 20, 30, 40)
	tup3 := NewTupleFromItems(10, 20, 50)
	tup4 := NewTupleFromItems(10, 20, 30)
	assertEq(t, tup1.Lt(tup2), true)
	assertEq(t, tup1.Lt(tup3), true)
	assertEq(t, tup1.Lt(tup4), false)

	tup5 := NewTupleFromItems(int16(90), int32(200), int64(300))
	tup6 := NewTupleFromItems(int8(100), int16(200), int32(300))
	assertEq(t, tup5.Lt(tup6), true)

	tup7 := NewTuple(3)
	tup8 := NewTupleFromItems(1, 2, 3)
	assertEq(t, tup7.Lt(tup8), true)
	assertEq(t, tup8.Lt(tup7), false)

	tup9 := NewTupleFromItems(NewTupleFromItems(1, 2), NewTupleFromItems(5, 10), NewTupleFromItems(10, 20))
	tup10 := NewTupleFromItems(NewTupleFromItems(1, 2), NewTupleFromItems(5, 10), NewTupleFromItems(10, 20))
	assertEq(t, tup9.Lt(tup10), false)
	tup11 := NewTupleFromItems(NewTupleFromItems(1, 2), NewTupleFromItems(5, 9), NewTupleFromItems(10, 20))
	assertEq(t, tup11.Lt(tup10), true)

	tup12 := NewTupleFromItems(10, 20, 30)
	tup13 := NewTupleFromItems(1, 2, 3, 4, 5, 6, 7)
	assertEq(t, tup12.Lt(tup13), false)
	assertEq(t, tup13.Lt(tup12), true)
}

func TestLe(t *testing.T) {
	tup1 := NewTupleFromItems(10, 20, 30)
	tup2 := NewTupleFromItems(10, 20, 30, 40)
	tup3 := NewTupleFromItems(10, 20, 50)
	tup4 := NewTupleFromItems(10, 20, 30)
	assertEq(t, tup1.Le(tup2), true)
	assertEq(t, tup1.Le(tup3), true)
	assertEq(t, tup1.Le(tup4), true)
}

func TestGt(t *testing.T) {
	tup1 := NewTupleFromItems(10, 20, 30)
	tup2 := NewTupleFromItems(10, 20, 30, 40)
	tup3 := NewTupleFromItems(10, 20, 50)
	tup4 := NewTupleFromItems(10, 20, 30)
	assertEq(t, tup1.Gt(tup2), false)
	assertEq(t, tup1.Gt(tup3), false)
	assertEq(t, tup1.Gt(tup4), false)
}

func TestGe(t *testing.T) {
	tup1 := NewTupleFromItems(10, 20, 30)
	tup2 := NewTupleFromItems(10, 20, 30, 40)
	tup3 := NewTupleFromItems(10, 20, 50)
	tup4 := NewTupleFromItems(10, 20, 30)
	assertEq(t, tup1.Ge(tup2), false)
	assertEq(t, tup1.Ge(tup3), false)
	assertEq(t, tup1.Ge(tup4), true)
}

func TestReverse(t *testing.T) {
	tup1 := NewTupleFromItems(1, 3, 5, 7, 9, 11, 13)
	tup1.Reverse()
	tup2 := NewTupleFromItems(13, 11, 9, 7, 5, 3, 1)
	assertEqTuple(t, tup1, tup2)
}

func TestIndex(t *testing.T) {
	tup1 := NewTupleFromItems(10, 20, 30, 40, 30)
	assertEq(t, tup1.Index(10, 0), 0)
	assertEq(t, tup1.Index(20, 0), 1)
	assertEq(t, tup1.Index(30, 0), 2)
	assertEq(t, tup1.Index(30, 2), 2)
	assertEq(t, tup1.Index(30, 3), 4)
	assertEq(t, tup1.Index(100, 0), -1)
}

func TestCount(t *testing.T) {
	tup1 := NewTupleFromItems(10, 20, 30, 40, 30, 20, 30, 40, 50)
	assertEq(t, tup1.Count(10, 0), 1)
	assertEq(t, tup1.Count(20, 0), 2)
	assertEq(t, tup1.Count(30, 0), 3)
	assertEq(t, tup1.Count(30, 5), 1)
	assertEq(t, tup1.Count(120, 0), 0)
}

func TestSortInternal(t *testing.T) {
	tup1 := NewTupleFromItems(1, 9, 7, 2, 3, 10, 5, 4, 8, 6)
	sort.Sort(tup1)
	assertEqTuple(t, tup1, NewTupleFromItems(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func TestSortTuples(t *testing.T) {
	tups := make([]*Tuple, 3)
	tup0 := NewTupleFromItems(10, 20, 30, 40, 50)
	tup1 := NewTupleFromItems(1, 2, 3, 4, 5, 6, 7)
	tup2 := NewTupleFromItems(10, 20, 30)
	tups[0] = tup0
	tups[1] = tup1
	tups[2] = tup2
	sort.Sort(ByElem(tups))
	assertEqTuple(t, tups[0], tup1)
	assertEqTuple(t, tups[1], tup2)
	assertEqTuple(t, tups[2], tup0)
}

func TestAppend(t *testing.T) {
	tup0 := NewTupleFromItems(1, 2, 3)
	tup1 := NewTupleFromItems("a", "b", "c")
	tup0.Append(tup1)
	assertEqTuple(t, tup0, NewTupleFromItems(1, 2, 3, "a", "b", "c"))
	tup1.AppendItems("d", "e", "f")
	assertEqTuple(t, tup1, NewTupleFromItems("a", "b", "c", "d", "e", "f"))
	// TODO: try with reference elements, show they are not deep copied during the append
}

func TestInsert(t *testing.T) {
	tup0 := NewTupleFromItems(1, 2, 3)
	tup1 := NewTupleFromItems("a", "b", "c")
	tup0.Insert(0, tup1)
	assertEqTuple(t, tup0, NewTupleFromItems("a", "b", "c", 1, 2, 3))

	tup2 := NewTupleFromItems(10, 20)
	tup1.Insert(1, tup2)
	assertEqTuple(t, tup1, NewTupleFromItems("a", 10, 20, "b", "c"))

	tup1.Insert(-1, NewTupleFromItems("x", "y", "z"))
	assertEqTuple(t, tup1, NewTupleFromItems("a", 10, 20, "b", "x", "y", "z", "c"))
}
