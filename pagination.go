package pagination

import "math"

// Pager represents an object that can divide items into discrete pages.
type Pager struct {
	Total       int  `json:"total"`
	PerPage     int  `json:"per_page"`
	CurrentPage int  `json:"current_page"`
	LastPage    int  `json:"last_page"`
	NextPage    *int `json:"next_page"`
	PrevPage    *int `json:"prev_page"`
}

// NewPager returns a new Pager.
// The currentPage and perPage must be positive integers.
// The total must be a zero or positive integer.
func NewPager(currentPage, perPage, total int) *Pager {
	p := &Pager{
		Total:       total,
		PerPage:     perPage,
		CurrentPage: currentPage,
		LastPage:    int(math.Ceil(float64(total) / float64(perPage))),
	}

	if currentPage < p.LastPage {
		p.NextPage = new(int)

		*p.NextPage = currentPage + 1
	}
	if currentPage > 1 {
		p.PrevPage = new(int)

		*p.PrevPage = currentPage - 1
	}

	return p
}
