package main

import (
	"fmt"
	"net/http"
)

func main() {
	InitSettings()
	http.HandleFunc("/", MainHandler)
	hostAndPort := fmt.Sprintf("%s:%s", Host(), Port())
	log("Starting on %s ", hostAndPort)
	if err := http.ListenAndServe(hostAndPort, nil); err != nil {
		fatal(err)
	}
}
