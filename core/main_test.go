package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirst(t *testing.T) {
	s := 2 * 2
	assert.Equal(t, 4, s)
}

func TestSecond(t *testing.T) {
	s := 2 + 3
	assert.Equal(t, 5, s)
}
