package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBatching(t *testing.T) {
	testCases := []struct {
		desc string
		in   []int
		size int
		want [][]int
	}{
		{
			desc: "",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			size: 3,
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}},
		},
		{
			desc: "",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			size: 10,
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}},
		},
		{
			desc: "",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			size: 15,
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
		},
		{
			desc: "",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			size: 20,
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
		},
		{
			desc: "",
			in:   []int{},
			size: 3,
			want: nil,
		},
		{
			desc: "",
			in:   nil,
			size: 3,
			want: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Batching(tC.in, tC.size)
			assert.Equal(t, tC.want, got)
		})
	}
}
