package handleConnection

import (
	"net"
)

func ConnectionWithVictim(IP string, Port string) (connection net.Conn, err error) {
	LocalAddress := IP + ":" + Port

	listener, err := net.Listen("tcp", LocalAddress)
	if err != nil {
		return connection, err
	}

	connection, err = listener.Accept()

	if err != nil {
		return connection, err
	} else {
		return
	}
}
