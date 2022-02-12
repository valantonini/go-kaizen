package functional_options

import "errors"

type options struct {
	port   *int
	scheme string
}

type Option func(o *options) error

type Server struct {
	Host   string
	Port   int
	Scheme string
}

var PortLessThanZeroError = errors.New("port must be 0 or greater")

func WithPort(p int) Option {
	return func(o *options) error {
		if p < 0 {
			return PortLessThanZeroError
		}
		o.port = &p
		return nil
	}
}

func WithScheme(s string) Option {
	return func(o *options) error {
		o.scheme = s
		return nil
	}
}

// NewServer - based on functional options pattern
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
// “2.11.3 Functional options pattern”
// 100 Go Mistakes and How to Avoid Them MEAP V09
func NewServer(host string, opts ...Option) (*Server, error) {
	if host == "" {
		host = "localhost"
	}
	var options options

	for _, o := range opts {
		err := o(&options)
		if err != nil {
			return nil, err
		}
	}

	var port int
	if options.port == nil {
		port = 80
		options.port = &port
	}

	if options.scheme == "" {
		options.scheme = "http"
	}

	server := Server{Host: host, Port: *options.port, Scheme: options.scheme}

	return &server, nil
}
