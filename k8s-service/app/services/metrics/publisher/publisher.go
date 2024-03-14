package publisher

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

// Collector defines a contract a collector must support
// so a consumer can retrieve metrics.
type Collector interface {
	Collect() (map[string]interface{}, error)
}

type Publisher func(map[string]interface{})

type Publish struct {
	log       *zap.SugaredLogger
	collector Collector
	publisher []Publisher
	wg        sync.WaitGroup
	timer     *time.Timer
	shutdown  chan struct{}
}

// New creates a Publish for consuming and publishing metrics.
func New(log *zap.SugaredLogger, collector Collector, interval time.Duration,
	publisher ...Publisher) (*Publish, error) {

	return &Publish{}, nil
}

type Stdout struct {
	log *zap.SugaredLogger
}

func NewStdout(log *zap.SugaredLogger) *Stdout {
	return &Stdout{}
}

func (p *Publish) Stop() {

}

// Publish publishers for writing to stdout.
func (s *Stdout) Publish(data map[string]interface{}) {}
