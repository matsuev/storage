package natsadapter

import (
	"log"
	"storage-service/internal/queue"

	"github.com/nats-io/nats.go"
)

// Config ..
type Config struct {
	URI   string
	Token string
	Depth int
}

// Adapter ...
type Adapter struct {
	config *Config
	conn   *nats.Conn
}

// New ...
func New(cfg *Config) (*Adapter, error) {
	if cfg == nil {
		return nil, queue.ErrDependency
	}

	conn, err := nats.Connect(cfg.URI, nats.Token(cfg.Token))
	if err != nil {
		log.Println(err)
		return nil, queue.ErrConnection
	}

	return &Adapter{
		config: cfg,
		conn:   conn,
	}, nil
}

// Close ...
func (a *Adapter) Close() error {
	if err := a.conn.Drain(); err != nil {
		log.Println(err)
		return queue.ErrClose
	}

	return nil
}

// Subscribe ..
func (a *Adapter) Subscribe(topic string) (chan []byte, error) {
	subChan := make(chan []byte, a.config.Depth)

	if _, err := a.conn.Subscribe(topic, func(msg *nats.Msg) {
		subChan <- msg.Data
	}); err != nil {
		log.Println(err)
		return nil, queue.ErrSubscribe
	}

	return subChan, nil
}

// Publish ...
func (a *Adapter) Publish(topic string, data []byte) error {
	if err := a.conn.Publish(topic, data); err != nil {
		log.Println(err)
		return queue.ErrPublish
	}

	return nil
}
