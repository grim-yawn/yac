package yac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Returns handler without any changes
func emptyWrapper(h Handler) Handler {
	return h
}

func TestWrapper_Add(t *testing.T) {
	ws := Wrappers{}
	ws = ws.Add(emptyWrapper)

	assert.Equal(t, len(ws), 1)
}

func TestWrappers_Wrap(t *testing.T) {
	ws := Wrappers{}

	ws = ws.Add(
		func(h Handler) Handler {
			return h
		},
		func(h Handler) Handler {
			return h
		},
	)
}

type counterHandler struct {
	calls int
}
