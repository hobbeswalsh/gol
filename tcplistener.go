package gol

import (
	"fmt"
	"log"
	"net"
)

func Listen(vip Vip) {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", vip.Ip, vip.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go serve(vip, conn)
	}
}
