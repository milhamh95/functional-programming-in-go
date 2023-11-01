package e_currying

import "fmt"

type (
	ServerOptions func(optionsCfg options) options
	TransportType int
)

const (
	UDP TransportType = iota
	TCP
)

type Server struct {
	options
	isAlive bool
}

type options struct {
	MaxConnection int
	TransportType TransportType
	Name          string
}

func NewServer(os ...ServerOptions) Server {
	opts := options{
		TransportType: TCP,
	}

	for _, option := range os {
		opts = option(opts)
	}

	return Server{
		options: opts,
		isAlive: true,
	}
}

func MaxConnection(n int) ServerOptions {
	return func(optionsCfg options) options {
		optionsCfg.MaxConnection = n
		return optionsCfg
	}
}

func ServerName(n string) ServerOptions {
	return func(o options) options {
		o.Name = n
		return o
	}
}

func Transport(t TransportType) ServerOptions {
	return func(o options) options {
		o.TransportType = t
		return o
	}
}

func createServer() {
	server := NewServer(MaxConnection(10), ServerName("MyFirstServer"))
	fmt.Printf("%+v\n", server)
}
