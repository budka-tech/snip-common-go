package port

import (
	"fmt"
)

type Port = uint16

const (
	Users  Port = 10000
	Gsm         = 10100
	Push        = 10100
	Assist      = 10300
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
