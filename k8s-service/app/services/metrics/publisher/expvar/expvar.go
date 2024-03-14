package expvar

import (
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Expvar struct {
	log    *zap.SugaredLogger
	server http.Server
	data   map[string]interface{}
	mu     sync.Mutex
}

func New(log *zap.SugaredLogger, host string, route string, readTimeout,
	writeTimeout time.Duration, idleTimeout time.Duration) *Expvar {

	return &Expvar{log: nil}
}

func (exp *Expvar) Stop(shutdownTimeout time.Duration) {

}

func (exp *Expvar) Publish(data map[string]interface{}) {

}

func (exp *Expvar) handler(w http.ResponseWriter, r *http.Request) {

}
