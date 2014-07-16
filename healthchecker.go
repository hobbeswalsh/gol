package gol

import (
	"fmt"
	"log"
	"net"
	"time"
)

func TCPHealthcheck(s Server) bool {
	t := 5 * time.Second
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port), t)
	if err != nil {
		log.Println("Time out!")
		return false
	}
	c.Close()
	return true
}

func loopHealthcheck(v *Vip, t time.Duration) {
	for {
		fmt.Printf("Performing health check for %s", v.Id)
		for _, server := range v.Servers {
			if !v.Healthcheck(server) {
				fmt.Println("OH DAMN")
				server.fail()
			} else {
				server.succeed()
			}
		}
		time.Sleep(t)
	}
}
