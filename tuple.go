package tuple

import (
	"fmt"
	_ "math"
)

type Tuple struct {
	data []interface{}
}

func NewTuple(n int) *Tuple {
	t := &Tuple{}
	t.data = make([]interface{}, n)
	return t
}

func NewTupleFromSlice(slice []interface{}) *Tuple {
	t := &Tuple{}
	t.data = slice
	return t
}

func NewTupleFromItems(items ...interface{}) *Tuple {
	t := NewTuple(len(items))
	for i, item := range items {
		t.Set(i, item)
	}
	return t
}

func (this *Tuple) Len() int {
	return len(this.data)
}

func (this *Tuple) Slice() []interface{} {
	return this.data
}

func (this *Tuple) Index(n int) int {
	// allow negative indexing as in Python
	if n < 0 {
		n = this.Len() + n
	}
	return n
}

func (this *Tuple) Set(n int, item interface{}) {
	this.data[this.Index(n)] = item
}

func (this *Tuple) Get(n int) interface{} {
	item := this.data[this.Index(n)]
	return item
}

func (this *Tuple) String() string {
	return fmt.Sprintf("%v", this.data)
}

func (this *Tuple) PopLeft() interface{} {
	if this.Len() < 1 {
		return nil
	}
	ret := this.data[0]
	this.data = this.data[1:]
	return ret
}

func (this *Tuple) PopRight() interface{} {
	if this.Len() < 1 {
		return nil
	}
	idx := this.Index(-1)
	ret := this.data[idx]
	this.data = this.data[:idx]
	return ret
}

func (this *Tuple) Reverse() {
	for i, j := 0, this.Len()-1; i < j; i, j = i+1, j-1 {
		this.data[i], this.data[j] = this.data[j], this.data[i]
	}
}

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

func (this *Tuple) Ne(other *Tuple) bool {
	return !this.Eq(other)
}

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
		if this.Get(i) != other.Get(i) {
			return true
		}
	}
	// if we get here then they matched up to n
	if tlen < olen {
		return true
	}
	return false
}

func (this *Tuple) Le(other *Tuple) bool {
	return this.Lt(other) || this.Eq(other)
}

func (this *Tuple) Gt(other *Tuple) bool {
	return !this.Le(other)
}

func (this *Tuple) Ge(other *Tuple) bool {
	return !this.Lt(other)
}
