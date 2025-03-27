// You can edit this code!
// Click here and start typing.
package main

import (
	"reflect"
	"testing"
)

func TestSortLinkedList(t *testing.T) {
	test_data := [][][]int {
		{{2, 3, 4, 5, 6}, {1, 5, 8, 11, 13}},
	}

	want := [][]int{
		{1,2,3,4,5,5,6,8,11,13},
	}

	for i := range test_data {
		got := SortLinkedList(test_data[i][0], test_data[i][1])
		var temp []int
		for got != nil {
			temp = append(temp, got.num)
			got = got.next
		}

		if !reflect.DeepEqual(temp, want[i]) {
			t.Errorf("Test Failed: Output: %v | Expected: %v | Data: %v", temp, want[i], test_data[i])
		}
	}
}