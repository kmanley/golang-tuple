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

	tup5 := NewTupleFromItems(int16(100), int32(200), int64(300))
	tup6 := NewTupleFromItems(int8(100), int16(200), int32(300))
	assertEq(t, tup5.Eq(tup6), true)

	tup7 := NewTupleFromItems(nil, nil, nil)
	tup8 := NewTuple(3)
	assertEq(t, tup7.Eq(tup8), true)

	tup9 := NewTupleFromItems(NewTupleFromItems(1, 2), NewTupleFromItems(5, 10), NewTupleFromItems(10, 20))
	tup10 := NewTupleFromItems(NewTupleFromItems(1, 2), NewTupleFromItems(5, 10), NewTupleFromItems(10, 20))
	assertEq(t, tup9.Eq(tup10), true)
	tup11 := NewTupleFromItems(NewTupleFromItems(1, 2), NewTupleFromItems(3, 10), NewTupleFromItems(10, 20))
	assertEq(t, tup10.Eq(tup11), false)

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
	assertEq(t, tup1.Eq(tup2), true)
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
	assertEq(t, tup1.Eq(NewTupleFromItems(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)), true)
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
	assertEq(t, tups[0].Eq(tup1), true)
	assertEq(t, tups[1].Eq(tup2), true)
	assertEq(t, tups[2].Eq(tup0), true)
}

/*
func TestWTF(t *testing.T) {
	if nil == nil {
		fmt.Println("nil equals nil")
	} else {
		fmt.Println("nil does not equal nil")
	}
}
*/
