package pagination

import "math"

// Pager represents an object that can divide items into discrete pages.
type Pager struct {
	currentPage, pages, perPage, total int
}

// NewPager returns a new Pager or an error.
// The currentPage and perPage must be positive integers.
// The total must be a zero or positive integer.
func NewPager(currentPage, perPage, total int) (*Pager, error) {
	return &Pager{
		currentPage: currentPage,
		pages:       int(math.Ceil(float64(total) / float64(perPage))),
		perPage:     perPage,
		total:       total,
	}, nil
}

// CurrentPage returns the current page number.
func (p Pager) CurrentPage() int {
	return p.currentPage
}

// LastPage returns the last page number in the pageset.
func (p Pager) LastPage() int {
	return p.pages
}

// NextPage returns the next page number in the pageset.
// A returned zero indicates there are no further pages.
func (p Pager) NextPage() int {
	if p.currentPage >= p.pages {
		return 0
	}

	return p.currentPage + 1
}

// PerPage returns the number of items per page in the pageset.
func (p Pager) PerPage() int {
	return p.perPage
}

// PrevPage returns the previous page in the pageset.
func (p Pager) PrevPage() int {
	return p.currentPage - 1
}

// Total returns the total number of items in the pageset.
func (p Pager) Total() int {
	return p.total
}
