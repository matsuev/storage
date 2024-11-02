package postgresadapter

import "storage-service/internal/database"

// Config ...
type Config struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

// Adapter ...
type Adapter struct {
	config *Config
}

// New ...
func New(cfg *Config) (*Adapter, error) {
	if cfg == nil {
		return nil, database.ErrDependency
	}

	return &Adapter{
		config: cfg,
	}, nil
}

// Close ...
func (a *Adapter) Close() error {
	return nil
}
