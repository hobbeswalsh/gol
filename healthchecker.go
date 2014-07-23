package gol

import (
	"fmt"
	"log"
	"net"
	"time"
)

func TCPHealthcheck(s *Server) bool {
	t := 5 * time.Second
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port), t)
	if err != nil {
		return false
	}
	c.Close()
	return true
}

func startHealthchecks(v Vip, t time.Duration) {
	for _, server := range v.Servers {
		actualServer := serverMap[server]
		go loopHealthcheck(actualServer, v.Healthcheck, t)
	}
}

func loopHealthcheck(s Server, hc Healthcheck, t time.Duration) {
	for {
		if !hc(&s) {
			log.Printf("Failed health check for %s", s.Id)
			unhealthyChan <- s.Id
		} else {
			log.Printf("Succeeded health check for %s", s.Id)
			healthyChan <- s.Id
		}
		time.Sleep(t)

	}

}
