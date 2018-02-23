package intersect

import (
	"testing"

	"github.com/bmizerany/assert"
	"fmt"
	"reflect"
)


func TestSimple(t *testing.T) {

	s := Simple([]string{"a","x"}, []string{"a","b"})
	fmt.Println(s)
	t1 := reflect.ValueOf(s)
	for i := 0; i < t1.Len() ; i ++ {
		fmt.Println(t1.Index(i).Interface().(string))
	}
}

func TestSorted(t *testing.T) {
	s := Sorted([]int{1}, []int{2})
	assert.Equal(t, s, []interface{}{})

	s = Sorted([]int{1, 2}, []int{2})
	assert.Equal(t, s, []interface{}{2})
}

func TestHash(t *testing.T) {
	s := Hash([]int{1}, []int{2})
	assert.Equal(t, s, []interface{}{})

	s = Hash([]int{1, 2}, []int{2})
	assert.Equal(t, s, []interface{}{2})
}
