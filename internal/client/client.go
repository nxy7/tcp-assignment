package client

import (
	"fmt"
	"net"

	"github.com/nxy7/tcp-assignment/internal/command"
)

type ResponseWriter interface {
	WriteResponse() error
}

type Client struct {
	self_id string
	conn    net.Conn
}

func (c *Client) ExecuteCommand(cmd command.Command) (string, error) {
	return cmd.Do(c.conn)
}

func MakeClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection established")
	defer func() {
		if err != nil {
			conn.Close()
		}
	}()

	j := command.Join{}
	id, err := j.Do(conn)
	if err != nil {
		return nil, err
	}

	c := Client{
		self_id: id,
		conn:    conn,
	}

	return &c, nil
}
