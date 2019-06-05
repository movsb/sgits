package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
)

func unauthorized(w http.ResponseWriter, reason string) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Login Please"`)
	w.WriteHeader(401)
	content := "401 Unauthorized\n"
	if reason != "" {
		content += "\n" + reason + "\n"
	}
	w.Write([]byte(content))
}

func auth(w http.ResponseWriter, req *http.Request) bool {
	if config.Username == "" {
		return true
	}

	username, password, ok := req.BasicAuth()
	if !ok {
		unauthorized(w, "")
		return false
	}

	if username != config.Username {
		unauthorized(w, "Unknown user")
		return false
	}

	if password != config.Password {
		unauthorized(w, "Wrong Password")
		return false
	}

	return true
}

func spawn(w http.ResponseWriter, req *http.Request) {
	ch := cgi.Handler{
		Path: config.Bin,
		Env: func() (env []string) {
			env = append(env, fmt.Sprintf("REMOTE_USER=%s", config.Username))
			env = append(env, "GIT_HTTP_EXPORT_ALL=")
			env = append(env, fmt.Sprintf("GIT_PROJECT_ROOT=%s", config.Root))
			return
		}(),
	}
	ch.ServeHTTP(w, req)
}
