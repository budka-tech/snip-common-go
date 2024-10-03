package client

import (
	"github.com/budka-tech/snip-common-go/port"
	"google.golang.org/grpc"
)

type Client[T any] struct {
	conn   *grpc.ClientConn
	client T
	port   port.Port
}

func NewClient[T any](port port.Port) *Client[T] {
	return &Client[T]{
		conn:   nil,
		client: *new(T),
		port:   port,
	}
}

func (receiver Client[T]) Init(f func(c *grpc.ClientConn) T) {
	conn, err := grpc.NewClient(port.FormatLocal(receiver.port))

	if err != nil {
		panic(err)
	}

	receiver.conn = conn
	receiver.client = f(conn)
}

func (receiver Client[T]) Dispose() {
	receiver.conn.Close()
}
