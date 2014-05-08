package tuple

import "fmt"

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

/*
func main() {
  fmt.Println("hi")

  t1 := []int{1,2,3}
  t2 := []int{2,3,4}

  fmt.Println(t1, t2)


  if t1 < t2 {
    fmt.Println("yes")
  }

  t := NewTuple(3, 1, 2, 3)
  t.Set(0, "kevin")
  t.Set(1, "thomas")
  t.Set(2, "manley")
  fmt.Println(t.String())
}
*/
