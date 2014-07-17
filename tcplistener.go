package gol

import (
	"fmt"
	"log"
	"net"
	"time"
)

func Listen(vip Vip) {
	go startHealthchecks(vip, 10*time.Second)
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
