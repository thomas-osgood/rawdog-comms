package comms

import "net"

type TcpTransmissionFunc func(net.Conn, []byte, string) error
