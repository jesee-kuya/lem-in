package colony

import (
	"reflect"
	"testing"
)

var (
	content1 = `3
##start
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

	content2 = `3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1`

	content3 = `3
2 5 0
##start
0 1 2
##end
1 9 2
3 5 4
0-2
0-3
2-1
3-1
2-3`

	TestCases = []struct {
		name    string
		content string
		arr     [][]int
	}{
		{"Test1", content1, [][]int{{1, 3, 4, 0}, {1, 2, 4, 0}, {1, 2, 7, 4, 0}, {1, 2, 7, 6, 0}, {1, 3, 5, 6, 0}, {1, 3, 5, 2, 4, 0}, {1, 3, 5, 2, 7, 4, 0}, {1, 3, 5, 2, 7, 6, 0}}},
		{"Test2", content2, [][]int{{0, 2, 3, 1}}},
		{"Test3", content3, [][]int{{0, 3, 1}, {0, 2, 1}}},
	}
)

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
