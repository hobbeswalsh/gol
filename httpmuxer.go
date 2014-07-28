package gol

import (
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
)

func serveHttp(
	vip Vip,
	url string,
	w http.ResponseWriter,
	req *http.Request) {

	s, err := vip.Select()
	if err != nil {
		log.Println(err.Error())
		return
	}
	r, err1 := http.Get(
		fmt.Sprintf("http://%s:%d%s", s.Ip, s.Port, url))
	if err1 != nil {
		log.Println(err1.Error())
		return
	}

	h := w.Header()
	h.Add("Content-Length", string(r.ContentLength))
	h.Add("Status", "200")
	err2 := r.Write(w)
	fmt.Println(err2)
	return
}
