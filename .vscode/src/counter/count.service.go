package counter

import (
	"fmt"
	"net/http"
)

const counterPath = "counter"

// ServeHTTP Serves the HTTP request
func (h *CountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

// SetupCountRoute sets up the rout
func SetupCountRoute(basePath string) {
	http.Handle(fmt.Sprintf("%s/%s", basePath, counterPath), new(CountHandler))
}
