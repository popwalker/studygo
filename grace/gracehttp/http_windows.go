// +build windows
package gracehttp

import (
	"net/http"
)

// Serve is a wrapper around standard ListenAndServe method.
func Serve(s *http.Server) error {
	return s.ListenAndServe()
}
