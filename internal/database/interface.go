package database

import (
	"context"
	"errors"
)

// IDatabaseAdapter ...
type IDatabaseAdapter interface {
	Close() error
	DeleteUserByID(ctx context.Context, id int64) error
	GetUserByID(ctx context.Context, id int64) (storage.User, error)
}

var (
	// ErrClose ...
	ErrClose = errors.New("database close error")
	// ErrDependency ...
	ErrDependency = errors.New("service dependency error")
)
