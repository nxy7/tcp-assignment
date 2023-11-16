package command

import (
	"bufio"
	"fmt"
	"net"
)

func (l *List) MakeHeader() string {
	return "Cmd: List\n"
}
func ListHeader() string {
	return (&List{}).MakeHeader()
}

func (l *List) Do(conn net.Conn) (string, error) {
	fmt.Println("Executing list command")
	_, err := fmt.Fprintf(conn, l.MakeHeader())
	if err != nil {
		return "", err
	}
	fmt.Println("Wrote list header")
	var entries []string
	buf := bufio.NewReader(conn)
	metadata_line, err := buf.ReadString('\n')
	if err != nil {
		return "", err
	}
	fmt.Println(metadata_line)
	for {
		r, err := buf.ReadString('\n')
		if err != nil {
			return "", err
		}
		if r == "DONE\n" {
			break
		} else {
			entries = append(entries, r)
		}
	}
	fmt.Println(entries)
	return "", nil
}
