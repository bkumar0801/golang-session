package integer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWhichIsBetter(t *testing.T) {
	WhichIsBetter()
}

func TestWhatIsExpected(t *testing.T) {
	assert.Equal(t, uint8(98), WhatIsExpected(6754))
}

func TestUpCast(t *testing.T) {
	assert.Equal(t, 10, UpCast(10))
}