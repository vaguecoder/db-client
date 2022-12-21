package mongowire

import (
	"fmt"
	"net"
)

const tcp = `tcp`

func ConnectTCP(addr string) (net.Conn, error) {
	conn, err := net.Dial(tcp, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to open TCP connection: %v", err)
	}

	return conn, nil
}
