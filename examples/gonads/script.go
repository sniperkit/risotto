package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	generator "github.com/sniperkit/risotto/pkg/generator"
)

type scriptTransformer struct {
	root string
}

func NewScriptTransformer(root string) *scriptTransformer {
	return &scriptTransformer{
		root: root,
	}
}

func (script *scriptTransformer) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestFile := filepath.Join(script.root, r.URL.Path)
	if filepath.Ext(requestFile) != ".jsx" {
		next(rw, r)
		return
	}

	fd, err := os.Open(requestFile)
	if err != nil {
		log.Error(err)
	}

	gen, err := generator.ParseAndGenerate(fd)
	if err != nil {
		log.Error(err)
	}

	if gen != nil {
		rw.Header().Add("Content-type", "text/javascript")
		io.Copy(rw, gen)
		return
	}

	next(rw, r)
}
