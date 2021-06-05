package main

import (
	"log"
	"net/http"
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
	err := http.ListenAndServe(config.Listen, nil)
	if err != nil {
		log.Fatalf("error: cannot listen on %s: %v", config.Listen, err)
	}
}
