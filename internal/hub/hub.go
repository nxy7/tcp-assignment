package hub

import (
	"bufio"
	"fmt"
	"net"

	"github.com/labstack/gommon/log"
	"github.com/nxy7/tcp-assignment/internal/client"
)

type Hub struct {
	Clients map[string]client.Client
}

func (h *Hub) Serve(port string) error {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer l.Close()

	for l, err := l.Accept(); err == nil; {
		go ServeRequest(l)
	}

	return nil
}

func ServeRequest(l net.Conn) {

	var data []byte
	// var buf bytes.Buffer
	b := bufio.NewScanner(l)
	for {
		_, err := l.Read(data)
		if err != nil {
			log.Error(err)
			return
		}

		// io.Copy(&buf, l)
		for b.Scan() {
			fmt.Println(b.Text())
		}

		r := "someResp"
		_, err = l.Write([]byte(r))
		if err != nil {
			log.Error(err)
			return
		}
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
