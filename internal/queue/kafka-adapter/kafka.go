package kafkaadapter

// Adapter ...
type Adapter struct{}

// New ...
func New() (*Adapter, error) {
	return &Adapter{}, nil
}

// Close ...
func (a *Adapter) Close() error {
	return nil
}

// Subscribe ..
func (a *Adapter) Subscribe(topic string) (chan []byte, error) {
	return nil, nil
}

// Publish ...
func (a *Adapter) Publish(topic string, data []byte) error {
	return nil
}
