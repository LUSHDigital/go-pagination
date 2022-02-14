package pagination

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPager(t *testing.T) {
	for _, tc := range []struct {
		currentPage int
		perPage     int
		total       int
		exp         *Pager
	}{
		{
			currentPage: 1,
			perPage:     1,
			total:       10,
			exp: &Pager{
				Total:       10,
				PerPage:     1,
				CurrentPage: 1,
				LastPage:    10,
				NextPage:    intPtr(2),
				PrevPage:    nil,
			},
		},
		{
			currentPage: 10,
			perPage:     1,
			total:       10,
			exp: &Pager{
				Total:       10,
				PerPage:     1,
				CurrentPage: 10,
				LastPage:    10,
				NextPage:    nil,
				PrevPage:    intPtr(9),
			},
		},
		{
			currentPage: 1,
			perPage:     10,
			total:       10,
			exp: &Pager{
				Total:       10,
				PerPage:     10,
				CurrentPage: 1,
				LastPage:    1,
				NextPage:    nil,
				PrevPage:    nil,
			},
		},
		{
			currentPage: 1,
			perPage:     1,
			total:       0,
			exp: &Pager{
				Total:       0,
				PerPage:     1,
				CurrentPage: 1,
				LastPage:    0,
				NextPage:    nil,
				PrevPage:    nil,
			},
		},
		{
			currentPage: 2,
			perPage:     33,
			total:       100,
			exp: &Pager{
				Total:       100,
				PerPage:     33,
				CurrentPage: 2,
				LastPage:    4,
				NextPage:    intPtr(3),
				PrevPage:    intPtr(1),
			},
		},
		{
			currentPage: 99,
			perPage:     1,
			total:       10,
			exp: &Pager{
				Total:       10,
				PerPage:     1,
				CurrentPage: 99,
				LastPage:    10,
				NextPage:    nil,
				PrevPage:    intPtr(98),
			},
		},
	} {
		t.Run(fmt.Sprintf("NewPager(%d,%d,%d)", tc.currentPage, tc.perPage, tc.total), func(t *testing.T) {
			p := NewPager(tc.currentPage, tc.perPage, tc.total)

			if !reflect.DeepEqual(p, tc.exp) {
				t.Errorf("got %#v, want %#v", p, tc.exp)
			}
		})
	}
}

func intPtr(i int) *int {
	return &i
}
