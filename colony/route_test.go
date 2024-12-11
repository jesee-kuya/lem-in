package colony

import (
	"reflect"
	"testing"
)

var content1 = `##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5`

var TestCases = []struct {
	name    string
	content string
	arr     [][]int
}{
	{"Test1", content1, [][]int{{1, 3, 4, 0}, {1, 2, 4, 0}, {1, 2, 7, 4, 0}, {1, 2, 7, 6, 0}, {1, 3, 5, 6, 0}, {1, 3, 5, 2, 4, 0}, {1, 3, 5, 2, 7, 4, 0}, {1, 3, 5, 2, 7, 6, 0}}},
}

func TestRoute(t *testing.T) {
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			arr, _ := Route(tc.content)
			if !reflect.DeepEqual(tc.arr, arr) {
				t.Errorf("Expected the routes:\n%v\nBut got the routes:\n%v\n", tc.arr, arr)
			}
		})
	}
}
