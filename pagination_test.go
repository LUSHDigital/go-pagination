package pagination

import "testing"

func TestCurrentPage(t *testing.T) {
	p, err := NewPager(1, 10, 20)
	if err != nil {
		t.Fatal(err)
	}

	want := 1

	if p.CurrentPage() != want {
		t.Errorf("got: %v, want: %v", p.CurrentPage(), want)
	}
}

func TestPerPage(t *testing.T) {
	p, err := NewPager(1, 10, 20)
	if err != nil {
		t.Fatal(err)
	}

	want := 10

	if p.PerPage() != want {
		t.Errorf("got: %v, want: %v", p.PerPage(), want)
	}
}

func TestTotal(t *testing.T) {
	p, err := NewPager(1, 10, 20)
	if err != nil {
		t.Fatal(err)
	}

	want := 20

	if p.Total() != want {
		t.Errorf("got: %v, want: %v", p.Total(), want)
	}
}

func TestZeroTotal(t *testing.T) {
	p, err := NewPager(1, 10, 0)
	if err != nil {
		t.Fatal(err)
	}

	if p.NextPage() != 0 {
		t.Errorf("got %v, want %v", p.NextPage(), 0)
	}

	if p.PrevPage() != 0 {
		t.Errorf("got %v, want %v", p.PrevPage(), 2)
	}
}

func TestNonZeroTotal(t *testing.T) {
	type args struct {
		currentPage, perPage, total int
	}

	type exp struct {
		lastPage, nextPage, prevPage int
	}

	for _, tc := range []struct {
		in   args
		want exp
	}{
		{
			in: args{
				currentPage: 2,
				perPage:     5,
				total:       2271,
			},
			want: exp{
				lastPage: 455,
				nextPage: 3,
				prevPage: 1,
			},
		},
		{
			in: args{
				currentPage: 10,
				perPage:     10,
				total:       100,
			},
			want: exp{
				lastPage: 10,
				nextPage: 0,
				prevPage: 9,
			},
		},
		{
			in: args{
				currentPage: 1,
				perPage:     10,
				total:       10,
			},
			want: exp{
				lastPage: 1,
				nextPage: 0,
				prevPage: 0,
			},
		},
		{
			in: args{
				currentPage: 7,
				perPage:     10,
				total:       50,
			},
			want: exp{
				lastPage: 5,
				nextPage: 0,
				prevPage: 6,
			},
		},
	} {
		p, err := NewPager(tc.in.currentPage, tc.in.perPage, tc.in.total)
		if err != nil {
			t.Fatal(err)
		}

		if p.LastPage() != tc.want.lastPage {
			t.Errorf("LastPage(%+v) got: %v, want: %v", tc.in, p.LastPage(), tc.want.lastPage)
		}

		if p.NextPage() != tc.want.nextPage {
			t.Errorf("NextPage(%+v) got: %v, want %v", tc.in, p.NextPage(), tc.want.nextPage)
		}

		if p.PrevPage() != tc.want.prevPage {
			t.Errorf("PrevPage(%+v) got: %v, want: %v", tc.in, p.PrevPage(), tc.want.prevPage)
		}
	}
}
