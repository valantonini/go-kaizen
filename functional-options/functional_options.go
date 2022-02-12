package functional_options

type options struct {
	port *int
}

type Option func(o *options) error

type Server struct {
	Host string
}

func NewServer(host string) (*Server, error) {
	return &Server{host}, nil
}
