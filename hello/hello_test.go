package hello

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestShouldSayHello(t *testing.T) {
	assert.Equal(t, "Hello,GoLang!", SayHello("GoLang"))
}
