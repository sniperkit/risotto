package parser

import (
	"testing"

	// external
	"github.com/stretchr/testify/assert"

	// internal
	ast "github.com/sniperkit/risotto/pkg/ast"
)

var validJSX = []string{
	"<div />",
	"<div param=\"value\"></div>",
	"<div><div /></div>",
	"<div prop={name} />",
}

func p(jsx string) (*ast.Program, error) {
	p := newParser("", jsx)
	return p.parse()
}

func TestJSX(t *testing.T) {
	for _, jsx := range validJSX {
		_, err := p(jsx)

		assert.NoError(t, err)

	}
}
