package gol

import (
	"errors"
	"fmt"
	"math/rand"
)

func RandomServer(servers []string) (s Server, err error) {
	var ids []Server
	for _, server := range servers {
		realServer := serverMap[server]
		if realServer.Healthy {
			ids = append(ids, realServer)
		}
	}
	if len(ids) == 0 {
		err = errors.New(fmt.Sprintf("VIP has no healthy servers behind it"))
		return
	}

	s = ids[rand.Intn(len(ids))]
	return
}

func (v *Vip) Select() (s Server, err error) {
	if len(v.Servers) == 0 {
		err = errors.New(fmt.Sprintf("VIP %s has no servers behind it", v.Id))
		return
	}

	return v.Algorithm(v.Servers)
}
