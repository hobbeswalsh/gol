package gol

import (
	"net"
	"time"
)

// NewVip constructs a new Vip struct from arguments. A new Vip should have
// at least one server (to make logical sense), but it doesn't have to.
// The Vip must be instantiated with a selection algorithm as well.
func NewVip(
	id string,
	ip net.IP,
	port int,
	algorithm Algorithm,
	servers ...string) (v Vip) {

	v.Id = id
	v.Ip = ip
	v.Port = port
	v.Algorithm = algorithm
	v.Servers = servers
	v.Healthcheck = TCPHealthcheck
	v.HealthcheckInterval = time.Second * 15
	v.FailureInterval = time.Second * 1
	v.ConsecutiveFailuresBeforeDown = 3
	v.ConsecutiveSuccessesUntilUp = 2
	vipMap[id] = v
	return
}

// NewServer constructs a new Server struct from arguments. A new Server
// Needs a (globally-unique) (not yet enforced) id, an IP, and a port. Right now
// Servers start out healthy, but once health checks are implemented that will change.
func NewServer(id string, ip net.IP, port int) (s Server) {
	s.Id = id
	s.Ip = ip
	s.Port = port
	s.Healthy = false
	serverMap[id] = s
	return
}
