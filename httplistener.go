package gol

import (
	"fmt"
	"log"
	"net/http"
)

func ListenHttp(vip Vip) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		go serveHttp(vip, req.URL.Path, w, req)
	})

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", vip.Port),
			nil))

}
