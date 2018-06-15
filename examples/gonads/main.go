package main

import (
	"flag"
	"net/http"
	"path/filepath"

	"github.com/codegangsta/negroni"
)

var (
	serverRoot string
)

func init() {
	flag.StringVar(&serverRoot, "root", "", "Root directory to serve")
	flag.Parse()
}

func main() {
	serverRoot, _ := filepath.Abs(serverRoot)

	script := NewScriptTransformer(serverRoot)
	style := NewSassTransformer(serverRoot)

	n := negroni.New(script, style)
	n.UseHandler(http.FileServer(http.Dir(serverRoot)))
	n.Run(":8080")
}
