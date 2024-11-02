package service

import (
	"errors"
	"log"
	"storage-service/internal/cache"
	"storage-service/internal/database"
	"storage-service/internal/queue"
)

// Service ...
type Service struct {
	queue    queue.IQueueAdapter
	cache    cache.ICacheAdapter
	database database.IDatabaseAdapter
}

// New ...
func New(
	q queue.IQueueAdapter,
	d database.IDatabaseAdapter,
	c cache.ICacheAdapter,
) (*Service, error) {
	if q == nil || d == nil || c == nil {
		return nil, errors.New("service dependency error")
	}

	return &Service{
		queue:    q,
		database: d,
		cache:    c,
	}, nil
}

// Start ...
func (s *Service) Start() error {

	log.Println("storage service started")

	return nil
}

// Shutdown ...
func (s *Service) Shutdown() error {

	log.Println("storage service stopped")

	return nil
}
