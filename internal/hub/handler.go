package hub

import "net"

type RequestHandler interface {
	HandleRequest(l net.Conn) error
}

type (
	ListHandler         struct{ *Hub }
	JoinHandler         struct{ *Hub }
	WriteMessageHandler struct{ *Hub }
)
