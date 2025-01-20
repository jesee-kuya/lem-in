package colony

import (
	"reflect"
	"testing"
)

var (
	arr1 = [][]any{{1, 3, 4, 0}, {1, 2, 4, 0}, {1, 2, 7, 4, 0}, {1, 2, 7, 6, 0}, {1, 2, 5, 6, 0}, {1, 3, 5, 6, 0}, {1, 3, 5, 2, 4, 0}, {1, 3, 5, 2, 7, 4, 0}, {1, 3, 5, 2, 7, 6, 0}}
	arr2 = [][]any{{0, 2, 3, 1}}
	arr3 = [][]any{{0, 2, 1}, {0, 3, 1}}

	TestCases1 = []struct {
		name string
		arr  [][]any
		res  [][]any
	}{
		{"Test1", arr1, [][]any{{1, 2, 4, 0}, {1, 3, 5, 6, 0}}},
		{"Test2", arr2, arr2},
		{"Test3", arr3, arr3},
	}
)

func TestClash(t *testing.T) {
	for _, tc := range TestCases1 {
		t.Run(tc.name, func(t *testing.T) {
			got := Clash(tc.arr)

			if !reflect.DeepEqual(got, tc.res) {
				t.Errorf("Expected the routes:\n%v\nBut got the routes:\n%v\n", tc.res, got)
			}
		})
	}
}
