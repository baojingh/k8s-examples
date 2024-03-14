package collector

import "net/http"

type Expvar struct {
	host   string
	tr     *http.Transport
	client *http.Client
}

// New creates a Expvar for collection metrics.
func New(host string) (*Expvar, error) {
	return &Expvar{}, nil
}

// Collect captures metrics on the host configure to this endpoint.
func (exp *Expvar) Collect() (map[string]interface{}, error) {
	return nil, nil
}
