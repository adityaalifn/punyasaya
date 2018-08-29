package domain

import (
	"fmt"

	"github.com/tokopedia/tdk/go/app/resource"
)

type Posts struct {
	OrderID   int
	ProductID int
	Quantity  int
	Invoice   string
}
