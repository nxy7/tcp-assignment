package command

import (
	"bufio"
	"fmt"
	"net"
)

func (l *Join) MakeHeader() string {
	return "Cmd: Join\n"
}

func JoinHeader() string {
	return (&Join{}).MakeHeader()
}

func (l *Join) Do(conn net.Conn) (string, error) {
	fmt.Println("Starting join command")
	b := bufio.NewReader(conn)
	_, err := fmt.Fprint(conn, l.MakeHeader())
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote header")
	r, err := b.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("my new id: %v\n", r)

	return r, nil
}
