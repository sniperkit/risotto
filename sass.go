package main

import (
	log "github.com/Sirupsen/logrus"
	libsass "github.com/wellington/go-libsass"
	"net/http"
	"os"
	"path/filepath"
)

type sassTransformer struct {
	root    string
	context *libsass.Context
}

func NewSassTransformer(root string) *sassTransformer {
	return &sassTransformer{
		root:    root,
		context: libsass.NewContext(),
	}
}

func (sass *sassTransformer) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestFile := filepath.Join(sass.root, r.URL.Path)
	if filepath.Ext(requestFile) != ".scss" {
		next(rw, r)
		return
	}

	fd, err := os.Open(requestFile)
	if err != nil {
		log.Error(err)
	}
	defer fd.Close()

	if err := sass.context.Compile(fd, rw); err != nil {
		log.Error(err)
	} else {
		rw.Header().Add("Content-type", "text/css")
		return
	}
	next(rw, r)
}
