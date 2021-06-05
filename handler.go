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
		content += reason + "\n"
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
		Dir:  `.`,
		Env: func() (env []string) {
			env = append(env, fmt.Sprintf("REMOTE_USER=%s", config.Username))
			env = append(env, "GIT_HTTP_EXPORT_ALL=")
			env = append(env, fmt.Sprintf("GIT_PROJECT_ROOT=%s", config.Root))
			return
		}(),
	}

	// net/http/cgi/host.go:122
	// Chunked request bodies are not supported by CGI.
	//
	// error: RPC failed; HTTP 400 curl 22 The requested URL returned error: 400
	// fatal: the remote end hung up unexpectedly
	//
	// https://github.com/git/git/blob/master/Documentation/config/http.txt#L216
	// https://gitlab.com/gitlab-org/gitlab/-/issues/17649
	// https://github.com/golang/go/issues/5613
	fixChunked(req)

	ch.ServeHTTP(w, req)
}

func fixChunked(req *http.Request) {
	if len(req.TransferEncoding) > 0 && req.TransferEncoding[0] == `chunked` {
		// hacking!
		req.TransferEncoding = nil
		req.Header.Set(`Transfer-Encoding`, `chunked`)

		// let cgi use Body.
		req.ContentLength = -1
	}
}
