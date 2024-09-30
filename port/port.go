package port

import (
	"fmt"
)

type Port = uint16

const (
	Users   Port = 10000
	UsersWs      = 3000
	Gsm          = 10100
	Assist       = 10200
)

func Format(port Port) string {
	return fmt.Sprintf("%d", port)
}

func FormatTCP(port Port) string {
	return fmt.Sprintf(":%d", port)
}
