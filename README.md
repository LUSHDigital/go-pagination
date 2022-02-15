# go-pagination

This package provides a simple pager object for paginating result sets.

## Installation

```bash
go get github.com/LUSHDigital/go-pagination
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/LUSHDigital/go-pagination"
)

func main() {
	currentPage := 2
	perPage := 2

	results := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	total := len(results)

	pager := pagination.NewPager(currentPage, perPage, total)

	fmt.Println(pager.PrevPage) // 1
	fmt.Println(pager.NextPage) // 3
	fmt.Println(pager.LastPage) // 5
}
```
