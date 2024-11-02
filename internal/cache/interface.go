package cache

import "errors"

// ICacheAdapter ...
type ICacheAdapter interface {
	Close() error
}

var (
	// ErrClose ...
	ErrClose = errors.New("cache close error")
	// ErrDependency ...
	ErrDependency = errors.New("service dependency error")
)
