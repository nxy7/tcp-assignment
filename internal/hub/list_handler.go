package hub

import (
	"fmt"
	"net"
)

func (h *ListHandler) HandleRequest(l net.Conn) error {
	fmt.Println("Handling list request")
	_, err := fmt.Fprintf(l, "Client amount: %v\n", len(h.Clients))
	if err != nil {
		return err
	}

	for _, u := range h.Clients {
		_, err := fmt.Fprint(l, u.id+"\n")
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(l, "DONE\n")
	if err != nil {
		return err
	}
	return nil
}
