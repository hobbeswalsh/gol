package gol

import (
	"errors"
	"fmt"
	"math/rand"
)

func RandomServer(sm map[string]Server) (s Server, err error) {
	var ids []string
	for serverId, server := range sm {
		if server.Healthy {
			ids = append(ids, serverId)
		}
	}
	if len(ids) == 0 {
		err = errors.New(fmt.Sprintf("VIP has no healthy servers behind it"))
		return
	}

	chosenId := ids[rand.Intn(len(ids))]
	s = sm[chosenId]
	return
}

func (v *Vip) Select() (s Server, err error) {
	if len(v.Servers) == 0 {
		err = errors.New(fmt.Sprintf("VIP %s has no servers behind it", v.Id))
		return
	}

	return v.Algorithm(v.Servers)
}
