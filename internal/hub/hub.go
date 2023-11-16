package hub

import (
	"bufio"
	"fmt"
	"net"

	"github.com/labstack/gommon/log"
	"github.com/nxy7/tcp-assignment/internal/command"
)

type Hub struct {
	Clients map[string]Client
}

func NewHub() Hub {
	return Hub{
		Clients: make(map[string]Client),
	}
}

type Client struct {
	id   string
	conn net.Conn
}

func (h *Hub) Serve(port string) error {
	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		lconn, err := l.Accept()
		if err != nil {
			return err
		}
		fmt.Println("Handling new connection")
		go func() {
			defer lconn.Close()
			err = h.ServeRequest(lconn)
			if err != nil {
				log.Error(err)
			}
		}()
	}
}

// Handles all communication with one client
func (h *Hub) ServeRequest(l net.Conn) error {
	var cmdHandler RequestHandler
	b := bufio.NewReader(l)

	r, err := b.ReadString('\n')
	if err != nil {
		return err
	}
	if r != command.JoinHeader() {
		l.Write([]byte("First message has to be join command\n"))
		return err
	}
	cmdHandler = &JoinHandler{h}
	err = cmdHandler.HandleRequest(l)
	if err != nil {
		fmt.Fprint(l, err)
		return err
	}
	cmdHandler = nil

	fmt.Println("Going into server req loop")
	for {
		line, err := b.ReadString('\n')
		if err != nil {
			fmt.Fprint(l, err)
			fmt.Printf("err: %v\n", err)
			return err
		}

		if line == command.ListHeader() {
			cmdHandler = &ListHandler{h}
			err = cmdHandler.HandleRequest(l)
			if err != nil {
				return err
			}
		} else if line == command.WriteHeader() {
			cmdHandler = &WriteMessageHandler{h}
			err = cmdHandler.HandleRequest(l)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("Unknown command", line)
		}

		cmdHandler = nil
	}
}

func MakeRequest() {
	_, err := net.Dial("tcp", ":22222")
	if err != nil {
		log.Error(err)
		return
	}

	// c.Write()
}
