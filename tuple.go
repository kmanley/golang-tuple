/*
Features:
 - python-style indexing
 - python-style comparison
 - helper functions like popleft, popright, reverse

TODO:
// TODO:
// func (this *Tuple) Zip(...*Tuple) Tuple {}
// func (this *Tuple) Map(func) ? or Apply?
// Dict() convert to Dict which I'm also working on
// append - or better pushleft, pushright
// insert/remove
// flatten?
// group?
// chunk?
// coalesce
// append, extend
// count, contains/index, sort

*/
package tuple

import (
	"fmt"
	_ "math"
	"reflect"
)

type Tuple struct {
	data []interface{}
}

// Creates a new empty Tuple of length n
func NewTuple(n int) *Tuple {
	t := &Tuple{}
	t.data = make([]interface{}, n)
	return t
}

// Creates a new Tuple from an existing slice
func NewTupleFromSlice(slice []interface{}) *Tuple {
	t := &Tuple{}
	t.data = slice
	return t
}

// Creates a new tuple from a literal sequence of items
func NewTupleFromItems(items ...interface{}) *Tuple {
	t := NewTuple(len(items))
	for i, item := range items {
		t.Set(i, item)
	}
	return t
}

// Returns the number of elements in the Tuple
func (this *Tuple) Len() int {
	return len(this.data)
}

// Returns the internal slice
func (this *Tuple) Slice() []interface{} {
	return this.data
}

// Convert n to an index into the internal slice.
// Negative numbers are supported, e.g. -1 points to the last item
func (this *Tuple) Index(n int) int {
	// allow negative indexing as in Python
	if n < 0 {
		n = this.Len() + n
	}
	return n
}

// Set the item at index n
func (this *Tuple) Set(n int, item interface{}) {
	this.data[this.Index(n)] = item
}

// Get the item at index n
func (this *Tuple) Get(n int) interface{} {
	item := this.data[this.Index(n)]
	return item
}

// Returns a string representation of the Tuple
func (this *Tuple) String() string {
	return fmt.Sprintf("%v", this.data)
}

// Pops the leftmost item from the Tuple and
// returns it
func (this *Tuple) PopLeft() interface{} {
	if this.Len() < 1 {
		return nil
	}
	ret := this.data[0]
	this.data = this.data[1:]
	return ret
}

// Pops the rightmost item from the Tuple and
// returns it
func (this *Tuple) PopRight() interface{} {
	if this.Len() < 1 {
		return nil
	}
	idx := this.Index(-1)
	ret := this.data[idx]
	this.data = this.data[:idx]
	return ret
}

// Reverses the Tuple in place
func (this *Tuple) Reverse() {
	for i, j := 0, this.Len()-1; i < j; i, j = i+1, j-1 {
		this.data[i], this.data[j] = this.data[j], this.data[i]
	}
}

func TupleElemEq(lhsi interface{}, rhsi interface{}) bool {
	lhsv, rhsv := reflect.ValueOf(lhsi), reflect.ValueOf(rhsi)
	// IsNil panics if type is not interface-y, so use IsValid instead
	if lhsv.IsValid() != rhsv.IsValid() {
		return false
	}
	switch lhsi.(type) {
	case nil:
		if rhsv.IsValid() {
			return false
		}
	case string:
		if lhsi.(string) != rhsi.(string) {
			return false
		}
	case int, int8, int16, int32, int64:
		if lhsv.Int() != rhsv.Int() {
			return false
		}
	case uint, uintptr, uint8, uint16, uint32, uint64:
		if lhsv.Uint() != rhsv.Uint() {
			return false
		}
	case float32, float64:
		if lhsv.Float() != rhsv.Float() {
			return false
		}
	case *Tuple:
		if lhsi.(*Tuple).Ne(rhsi.(*Tuple)) {
			return false
		}
	default:
		//if !lhsv.IsValid() && !rhsv.IsValid() {
		//return false
		//}
		// TODO: allow user-defined callback for unsupported types
		panic(fmt.Sprintf("unsupported type %#v for Eq in Tuple", lhsi))
	}
	return true
}

// Returns True if this Tuple is elementwise equal to other
func (this *Tuple) Eq(other *Tuple) bool {
	if this.Len() != other.Len() {
		return false
	}
	//return reflect.DeepEqual(this.data, other.data)
	for i := 0; i < this.Len(); i++ {
		lhsi, rhsi := this.Get(i), other.Get(i)
		if !TupleElemEq(lhsi, rhsi) {
			return false
		}
	}
	return true
}

// Returns True if this Tuple is not elementwise equal to other
func (this *Tuple) Ne(other *Tuple) bool {
	return !this.Eq(other)
}

func TupleElemLt(lhsi interface{}, rhsi interface{}) bool {
	lhsv, rhsv := reflect.ValueOf(lhsi), reflect.ValueOf(rhsi)
	if lhsv.IsValid() && !rhsv.IsValid() {
		// zero value is considered least
		return false
	}
	switch lhsi.(type) {
	case nil:
		if rhsv.IsValid() {
			return true
		}
	case string:
		if lhsi.(string) < rhsi.(string) {
			return true
		}
	case int, int8, int16, int32, int64:
		if lhsv.Int() < rhsv.Int() {
			return true
		}
	case uint, uintptr, uint8, uint16, uint32, uint64:
		if lhsv.Uint() < rhsv.Uint() {
			return true
		}
	case float32, float64:
		if lhsv.Float() < rhsv.Float() {
			return true
		}
	case *Tuple:
		if lhsi.(*Tuple).Lt(rhsi.(*Tuple)) {
			return true
		}
	default:
		// TODO: allow user-defined callback for unsupported types
		panic(fmt.Sprintf("unsupported type %#v for Lt in Tuple", lhsi))
	}
	return false
}

// Returns True if this Tuple is elementwise less than other
// TODO: use Sortable (sp?) interface instead
func (this *Tuple) Lt(other *Tuple) bool {
	tlen, olen := this.Len(), other.Len()
	var n int
	if tlen < olen {
		n = tlen
	} else {
		n = olen
	}
	for i := 0; i < n; i++ {
		lhsi, rhsi := this.Get(i), other.Get(i)
		if TupleElemLt(lhsi, rhsi) {
			return true
		}
	}
	// if we get here then they matched up to n
	if tlen < olen {
		return true
	}
	return false
}

// Returns True if this Tuple is elementwise less than
// or equal to other
func (this *Tuple) Le(other *Tuple) bool {
	return this.Lt(other) || this.Eq(other)
}

// Returns True if this Tuple is elementwise greater than other
func (this *Tuple) Gt(other *Tuple) bool {
	return !this.Le(other)
}

// Returns True if this Tuple is elementwise greater than
// or equal to other
func (this *Tuple) Ge(other *Tuple) bool {
	return !this.Lt(other)
}
