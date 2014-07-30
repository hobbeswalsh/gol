package gol

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ListenHttp(vip Vip) {
	go startHealthchecks(vip, 10*time.Second)
	go watchHealthChecks()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		contentChan := make(chan *http.Response)
		doneChan := make(chan bool)
		go dispatchHttpRequest(vip, req, contentChan, doneChan)

		for {
			select {
			case resp := <-contentChan:
				h := w.Header()
				respHeader := resp.Header
				h.Add("Content-Type", respHeader.Get("Content-Type"))
				w.WriteHeader(resp.StatusCode)
				content, _ := ioutil.ReadAll(resp.Body)
				fmt.Fprint(w, string(content))
				return
			case <-doneChan:
				return
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	})

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", vip.Port),
			nil))

}
