package gol

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	// "time"
)

func Listen(ip net.IP, port int) {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
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
		go echo(conn)
	}
}

func echo(c net.Conn) {
	s1 := Server{
		net.IPv4(127, 0, 0, 1),
		8080,
	}
	s2 := Server{
		net.IPv4(127, 0, 0, 1),
		8081,
	}

	servers := []Server{s1, s2}

	s := servers[rand.Intn(len(servers))]

	sc, err1 := net.Dial("tcp", fmt.Sprintf("%s:%s", s.Ip, s.Port))

	go io.Copy(sc, c)
	go io.Copy(c, sc)

	// io.Copy(c2, c)

	// io.Copy(c, c2)
	// fmt.Println("Copied back")
	// // now := time.Now()
	// timeout := now.Add(1000 * time.Millisecond)
	// c.SetDeadline(timeout)
	// Shut down the connection.
	defer c.Close()
	// c.Close()
}
