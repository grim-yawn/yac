package yac

import (
	"fmt"

	"github.com/weitbelou/yac/params"
)

// Method name to handler
type Methods map[string]Handler

// Pattern to methods map
type Routes map[string]Methods

func NewRoutes() Routes {
	return make(Routes)
}

// Adds route to routes
// Returns error if there is such route.
//
// See: Has(pattern, method)
func (rs Routes) Add(pattern, method string, h Handler) error {
	if rs.Has(pattern, method) {
		return ErrRouteAlreadyExists{pattern: pattern, method: method}
	}

	if !rs.HasPattern(pattern) {
		rs[pattern] = make(Methods)
	}

	rs[pattern][method] = h
	return nil
}

// Returns handler for 'pattern' 'method'
// Returns ErrPatternNotFound if pattern that matches path not found
// Returns ErrMethodNotAllowed if this method not found for matching pattern
func (rs Routes) Get(path, method string) (Handler, params.Container, error) {
	pattern, p := rs.GetPatternAndParamsByPath(path)

	if pattern == "" {
		return nil, nil, ErrPathNotFound{path: path}
	}

	if !rs.Has(pattern, method) {
		return nil, nil, ErrMethodNotAllowed{path: path, method: method}
	}

	return rs[pattern][method], p, nil
}

// Returns first pattern that matches this path
// If not found returns empty string
func (rs Routes) GetPatternAndParamsByPath(path string) (string, params.Container) {
	for pattern := range rs {
		p, err := params.New(pattern, path)
		if err != nil {
			continue
		}

		return pattern, p
	}

	return "", nil
}

// Checks if route in routes
func (rs Routes) Has(pattern, method string) bool {
	_, ok := rs[pattern][method]
	return ok
}

// Checks if pattern in routes
func (rs Routes) HasPattern(pattern string) bool {
	_, ok := rs[pattern]
	return ok
}

type ErrRouteAlreadyExists struct {
	pattern string
	method  string
}

func (e ErrRouteAlreadyExists) Error() string {
	return fmt.Sprintf("route with pattern '%s' and method '%s' already exists", e.pattern, e.method)
}

type ErrPathNotFound struct {
	path string
}

func (e ErrPathNotFound) Error() string {
	return fmt.Sprintf("pattern that matches '%s' not found", e.path)
}

type ErrMethodNotAllowed struct {
	path   string
	method string
}

func (e ErrMethodNotAllowed) Error() string {
	return fmt.Sprintf("method '%s' not found for path '%s'", e.method, e.path)
}
