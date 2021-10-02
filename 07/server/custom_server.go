package server

import (
	"fmt"
	"net/http"
)

// Viết custom handle
type handle struct{}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("Home page"))
	} else if r.URL.Path == "/about" {
		w.Write([]byte("About page"))
	} else {
		w.Write([]byte("Page not found"))
	}
}

func DemoCustomHandle() {
	handle := new(handle)
	fmt.Println(http.ListenAndServe(":3000", handle))
}
