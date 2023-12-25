package main

import (
	"log"
	"net/http"
	"github.com/coreos/go-systemd/v22/activation"
)

func handle(w http.ResponseWriter, req *http.Request) {
	if !auth(w, req) {
		return
	}
	log.Println(req.Method, req.RequestURI)
	spawn(w, req)
}

func main() {
	mustParseConfig("sgits.yml")
	http.HandleFunc("/", handle)
	log.Println("Listen on", config.Listen)
	listeners, err := activation.Listeners()
	if err != nil {
		panic(err)
	}
	if len(listeners) == 0 {
		err := http.ListenAndServe(config.Listen, nil)
		if err != nil {
			log.Fatalf("error: cannot listen on %s: %v", config.Listen, err)
		}
	} else if len(listeners) == 1 {
		http.Serve(listeners[0], nil)
	} else {
		log.Fatalf("error: only one activated socket is expected.  got: %d", len(listeners))
	}
}
