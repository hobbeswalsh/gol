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
		go loopHealthcheck(actualServer, v, t)
	}
}

func loopHealthcheck(s Server, v Vip, t time.Duration) {
	var failures = 0
	var successes = 0
	for {
		if !v.Healthcheck(&s) {
			failures += 1
			log.Printf("Failed health check for %s", s.Id)
			if failures >= v.ConsecutiveFailuresBeforeDown {
				log.Printf("Setting server %s to UNHEALTHY", s.Id)
				successes = 0
				unhealthyChan <- s.Id
			}

			time.Sleep(v.FailureInterval)
		} else {
			successes += 1
			failures = 0
			log.Printf("Succeeded health check for %s", s.Id)
			if successes >= v.ConsecutiveSuccessesUntilUp {
				log.Printf("Setting server %s to HEALTHY", s.Id)
				healthyChan <- s.Id
			}
			time.Sleep(v.HealthcheckInterval)
		}

	}

}
