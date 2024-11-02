package redisadapter

// Config ...
type Config struct {
	Host string
	Port int
	User string
	Pass string
	DB   int
}

// Adapter ...
type Adapter struct {
	config *Config
}

// New ...
func New(cfg *Config) (*Adapter, error) {
	return &Adapter{
		config: cfg,
	}, nil
}

// Close ...
func (a *Adapter) Close() error {
	return nil
}
