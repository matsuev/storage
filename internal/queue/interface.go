package queue

import "errors"

// IQueueAdapter ...
type IQueueAdapter interface {
	Subscribe(topic string) (chan []byte, error)
	Publish(topic string, data []byte) error
}

var (
	// ErrConnection ...
	ErrConnection = errors.New("queue connection error")
	// ErrSubscribe ...
	ErrSubscribe = errors.New("queue subscribe error")
	// ErrPublish ...
	ErrPublish = errors.New("queue publish error")
	// ErrClose ...
	ErrClose = errors.New("queue close error")
	// ErrDependency ...
	ErrDependency = errors.New("service dependency error")
)
