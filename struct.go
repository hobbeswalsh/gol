package gol

import (
	"net"
)

type Server struct {
	Id      string
	Ip      net.IP
	Port    int
	Healthy bool
}

type Vip struct {
	Id        string
	Ip        net.IP
	Port      int
	Algorithm Algorithm
	Servers   []Server
}

type Algorithm func([]Server) (Server, error)
