package port

import (
	"fmt"
)

type Port = uint16

const (
	Redis           = 6379
	UsersGrpc  Port = 10100
	GsmGrpc         = 10200
	GsmHttp         = 10201
	PushGrpc        = 10300
	AssistGrpc      = 10400
	S3Grpc          = 10500
)

func Format(port Port) string {
	return fmt.Sprintf("%d", port)
}

func FormatTCP(port Port) string {
	return fmt.Sprintf(":%d", port)
}

func FormatLocal(port Port) string {
	return fmt.Sprintf("localhost:%d", port)
}
