package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Goodbye is a simple handler
type Goodbye struct {
	l *log.Logger
}

// NewHGoodbye creates a new goodbye handler with the given logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Goodbye request")

	// write the response
	fmt.Fprintf(rw, "Goodbye")
}
