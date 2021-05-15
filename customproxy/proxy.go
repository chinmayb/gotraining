package customproxy

import (
	"net"
)

func createProxy(protocol string, port string) error {
	listener, err := net.Listen(protocol, net.JoinHostPort("0.0.0.0", port))
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()

	}

}

func main() {

}
