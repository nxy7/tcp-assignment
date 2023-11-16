package command

import "net"

type (
	List         struct{}
	Join         struct{}
	WriteMessage struct {
		Message string
	}
)

type Command interface {
	MakeHeader() string
	Do(l net.Conn) (string, error)
}
