package gol

import (
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
)

func dispatchHttpRequest(
	vip Vip,
	req *http.Request,
	c chan *http.Response,
	dc chan bool) {

	s, err := vip.Select()
	if err != nil {

		log.Println(err.Error())
		dc <- true
		return
	}
	r, fetchError := http.Get(
		fmt.Sprintf("http://%s:%d%s", s.Ip, s.Port, req.URL.Path))

	if fetchError != nil {
		fmt.Println(fetchError)
		dc <- true
		return
	}
	c <- r

	return
}
