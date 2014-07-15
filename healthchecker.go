package gol

import (
	"fmt"
	"net"
)

func TCPHealthcheck(s Server) bool {
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		return false
	}
	c.Close()
	return true
}
