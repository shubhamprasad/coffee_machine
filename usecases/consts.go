package usecases

import "sync"

var (
	mutex sync.Mutex
)

// string constants
const (
	Outlets = "outlets"
	CountN = "count_n"
	TotalItemsQuantity = "total_items_quantity"
	Beverages = "beverages"
)