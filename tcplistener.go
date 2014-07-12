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
		go loadBalance(conn)
	}
}

func loadBalance(c net.Conn) {
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

	sc, err1 := net.Dial("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err1 != nil {
		fmt.Println(err1)
		c.Close()
		return
	}

	serverDoneChan := make(chan bool)
	clientDoneChan := make(chan bool)

	go copyStream(sc, c, serverDoneChan)
	go copyStream(c, sc, clientDoneChan)

	for {
		select {

		case <-serverDoneChan:
			defer sc.Close()
			defer c.Close()
			return

		case <-clientDoneChan:
			defer c.Close()
			defer sc.Close()
			return

		default:

		}
	}

	defer c.Close()
}

func copyStream(src, dst net.Conn, dc chan bool) {
	// log.Println(src.LocalAddr().String())
	// log.Println(src.RemoteAddr().String())
	io.Copy(dst, src)
	dc <- true
}
