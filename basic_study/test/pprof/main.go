package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()

	time.Sleep(1000 * time.Second)
}
