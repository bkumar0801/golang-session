package integer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWhichIsBetter(t *testing.T) {
	WhichIsBetter()
}

func TestWhatIsExpected(t *testing.T) {
	assert.Equal(t, 20, WhatIsExpected(6754))
}
