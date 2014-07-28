package gol

import (
	"fmt"
	"io"
	"log"
	"net"
)

func serve(vip Vip, c net.Conn) {
	s, err := vip.Select()
	if err != nil {
		log.Println(err.Error())
		c.Close()
		return
	}

	sc, err1 := net.Dial("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err1 != nil {
		log.Println(err1.Error())
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
