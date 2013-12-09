// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine

// A stand-alone HTTP server providing a web UI for task management.
package main

import (
	"net/http"

	"github.com/campoy/todo/server"
)

func main() {
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}
