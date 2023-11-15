package hub

import "net"

type RequestHandler interface {
	HandleRequest(l net.Conn) error
}
