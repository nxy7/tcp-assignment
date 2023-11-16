package command

import (
	"fmt"
	"net"
)

func (l *WriteMessage) MakeHeader() string {
	return "Cmd: Write\n"
}
func WriteHeader() string {
	return (&WriteMessage{}).MakeHeader()
}
func (l *WriteMessage) Do(conn net.Conn) (string, error) {
	_, err := fmt.Fprintf(conn, l.MakeHeader())
	if err != nil {
		return "", err
	}
	_, err = fmt.Fprintf(conn, l.Message)
	if err != nil {
		return "", err
	}
	_, err = fmt.Fprintf(conn, "DONE")
	if err != nil {
		return "", err
	}
	return "", nil
}
