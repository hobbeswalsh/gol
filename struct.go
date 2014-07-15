package gol

import (
	"net"
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
	Algorithm   Algorithm
	Healthcheck Healthcheck
	Servers     []Server
}

type Algorithm func([]Server) (Server, error)
type Healthcheck func(Server) bool
