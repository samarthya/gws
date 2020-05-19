package counter

import "sync"

// CountHandler Handles the counter
type CountHandler struct {
	sync.Mutex     // guards n
	n          int //stores the counter
}
