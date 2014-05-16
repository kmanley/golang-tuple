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

// TODO:
// func (this *Tuple) Zip(...*Tuple) Tuple {}
// func (this *Tuple) Map(func) ? or Apply?
// Dict() convert to Dict which I'm also working on
// append - or better pushleft, pushright
// insert
// flatten?
// group?
// chunk?
// coalesce

// Returns True if this Tuple is elementwise equal to other
func (this *Tuple) Eq(other *Tuple) bool {
	if this.Len() != other.Len() {
		return false
	}
	for i := 0; i < this.Len(); i++ {
		if this.Get(i) != other.Get(i) {
			return false
		}
	}
	return true
}

// Returns True if this Tuple is not elementwise equal to other
func (this *Tuple) Ne(other *Tuple) bool {
	return !this.Eq(other)
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
		lhs := this.Get(i)
		typ := reflect.TypeOf(lhs)
		if lhs.(typ) < other.Get(i).(typ) {
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
