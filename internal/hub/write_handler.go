package hub

import (
	"bufio"
	"fmt"
	"net"
)

func (h *WriteMessageHandler) HandleRequest(l net.Conn) error {
	fmt.Println("Handling Write Command")
	buf := bufio.NewReader(l)
	msg, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Println(msg)

	return nil
}
