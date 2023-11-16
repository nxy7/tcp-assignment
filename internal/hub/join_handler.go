package hub

import (
	"fmt"
	"net"

	"github.com/nxy7/tcp-assignment/internal/utils"
)

func (h *JoinHandler) HandleRequest(l net.Conn) error {
	fmt.Println("Handling Join Request")
	newClient := Client{
		id:   utils.RandomId(),
		conn: l,
	}
	h.Clients[newClient.id] = newClient
	_, err := l.Write([]byte(newClient.id + "\n"))
	if err != nil {
		return err
	}
	return nil
}
