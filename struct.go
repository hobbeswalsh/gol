package gol

import (
	"net"
	"time"
)

type Listener struct {
	Id   string
	Ip   net.IP
	Port int
}

type Server struct {
	Listener
	Healthy bool
}

type Vip struct {
	Listener
	HealthcheckInterval           time.Duration
	FailureInterval               time.Duration
	ConsecutiveFailuresBeforeDown int
	ConsecutiveSuccessesUntilUp   int
	Algorithm                     Algorithm
	Healthcheck                   Healthcheck
	Servers                       []string
}

type Algorithm func([]string) (Server, error)
type Healthcheck func(*Server) bool
type Healthcheckable interface {
	markDown()
	markUp()
}
