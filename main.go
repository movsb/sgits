package main

import "net/http"

func handle(w http.ResponseWriter, req *http.Request) {
	if !auth(w, req) {
		return
	}
	spawn(w, req)
}

func main() {
	mustParseConfig("sgits.yml")
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(config.Listen, nil)
	if err != nil {
		panic(err)
	}
}
