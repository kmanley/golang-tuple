package tuple

import (
	"fmt"
	_ "math"
)

type Tuple struct {
	data []interface{}
}

func NewTuple(n int, items ...interface{}) *Tuple {
	t := &Tuple{}
	t.data = make([]interface{}, n)
	for i, item := range items {
		t.Set(i, item)
	}
	return t
}

func (t *Tuple) Len() int {
	return len(t.data)
}

func (t *Tuple) Slice() []interface{} {
	return t.data
}

func (t *Tuple) Set(n int, item interface{}) {
	t.data[n] = item
}

func (t *Tuple) Get(n int) interface{} {
	item := t.data[n]
	return item
}

func (t *Tuple) String() string {
	return fmt.Sprintf("%v", t.data)
}

func (t *Tuple) PopLeft() interface{} {
	if t.Len() < 1 {
		return nil
	}
	ret := t.data[0]
	t.data = t.data[1:]
	return ret
}

func (t *Tuple) Eq(other *Tuple) bool {
	if t.Len() != other.Len() {
		return false
	}
	for i := 0; i < t.Len(); i++ {
		if t.Get(i) != other.Get(i) {
			return false
		}
	}
	return true
}

func (t *Tuple) Ne(other *Tuple) bool {
	return !t.Eq(other)
}

func (t *Tuple) Lt(other *Tuple) bool {
	tlen, olen := t.Len(), other.Len()
	var n int
	if tlen < olen {
		n = tlen
	} else {
		n = olen
	}
	for i := 0; i < n; i++ {
		if t.Get(i) != other.Get(i) {
			return true
		}
	}
	// if we get here then they matched up to n
	if tlen < olen {
		return true
	}
	return false
}

func (t *Tuple) Le(other *Tuple) bool {
	return t.Lt(other) || t.Eq(other)
}

func (t *Tuple) Gt(other *Tuple) bool {
	return !t.Le(other)
}

func (t *Tuple) Ge(other *Tuple) bool {
	return !t.Lt(other)
}
