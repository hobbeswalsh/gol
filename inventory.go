package gol

import (
	"errors"
	"fmt"
)

var vipMap = make(map[string]Vip)

func AddServer(vipId string, server Server) error {
	vip, vipExists := vipMap[vipId]

	if vipExists == false {
		return errors.New(fmt.Sprintf("No such VIP: %s", vipId))
	}

	if serverAlreadyBound(vip, server) {
		return errors.New(fmt.Sprintf("Server %s:%d is already bound to VIP %s", server.Ip, server.Port, vipId))
	}

	vip.Servers = append(vip.Servers, server)
	return nil
}

func serverAlreadyBound(vip Vip, server Server) bool {
	for _, boundServer := range vip.Servers {
		if server.Id == boundServer.Id {
			return true
		}
	}
	return false
}

func fail(s *Server) {
	s.Healthy = false
}

func succeed(s *Server) {
	s.Healthy = true
}
