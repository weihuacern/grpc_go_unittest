package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthServer_Init(t *testing.T) {
	s, err := NewAuthServer()
	err = s.Init()
	assert.Nil(t, err)
}
