package colony

import (
	"reflect"
	"testing"
)

var (
	arr4       = [][]int{{1, 3, 4, 0}, {1, 2, 5, 6, 0}}
	TestCases2 = []struct {
		name         string
		numberOfAnts int
		arr          [][]int
		res          [][]string
	}{
		{"Test1", 3, arr4, [][]string{{"L1-3 L2-2"}, {"L1-4 L2-5 L3-3"}, {"L1-0 L2-6 L3-4"}, {"L2-0 L3-0"}}},
		//	{"Test2", arr2, [][]string{{"L1-2"}, {"L1-3 L2-2"}, {"L1-1 L2-3 L3-2"}, {"L2-1 L3-3"}, {"L3-1"}}},
		//	{"Test3", arr3, [][]string{{"L1-2 L2-3"}, {"L1-1 L2-1 L3-2"}, {"L3-1"}}},
	}
)

func TestPath(t *testing.T) {
	for _, tc := range TestCases2 {
		t.Run(tc.name, func(t *testing.T) {
			got := Path(tc.arr, tc.numberOfAnts)

			if !reflect.DeepEqual(got, tc.res) {
				t.Errorf("Expected the paths:\n%v\nbut got the paths:\n%v\n", tc.res, got)
			}
		})
	}
}
