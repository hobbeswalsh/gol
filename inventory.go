package gol

import (
	"errors"
	"fmt"
)

var vipMap = make(map[string]Vip)
var serverMap = make(map[string]Server)
var healthyChan = make(chan string)
var unhealthyChan = make(chan string)

func AddServer(vipId string, server string) error {
	vip, vipExists := vipMap[vipId]

	if vipExists == false {
		return errors.New(fmt.Sprintf("No such VIP: %s", vipId))
	}

	if serverAlreadyBound(vip, server) {
		return errors.New(fmt.Sprintf("Server %s is already bound to VIP %s", server, server, vipId))
	}

	vip.Servers = append(vip.Servers, server)
	return nil
}

func serverAlreadyBound(vip Vip, server string) bool {
	for _, boundServer := range vip.Servers {
		if server == boundServer {
			return true
		}
	}
	return false
}

func watchHealthChecks() {
	for {
		select {
		case unhealthyServer := <-unhealthyChan:
			server, ok := serverMap[unhealthyServer]
			if ok {
				server.Healthy = false
				serverMap[unhealthyServer] = server
			}
		case healthyServer := <-healthyChan:
			server, ok := serverMap[healthyServer]
			if ok {
				server.Healthy = true
				serverMap[healthyServer] = server
			}
		default:

		}
	}
}
