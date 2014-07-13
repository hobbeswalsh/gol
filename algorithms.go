package gol

import (
	"errors"
	"fmt"
	"math/rand"
)

func RandomServer(ss []Server) (s Server, err error) {
	if len(ss) == 0 {
		err = errors.New(fmt.Sprintf("VIP has no servers behind it"))
		return
	}
	s = ss[rand.Intn(len(ss))]
	return
}

func (v *Vip) Select() (s Server, err error) {
	if len(v.Servers) == 0 {
		err = errors.New(fmt.Sprintf("VIP %s has no servers behind it", v.Id))
		return
	}
	return v.Algorithm(v.Servers)
}
