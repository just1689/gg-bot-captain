package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomPersonality(t *testing.T) {
	personality := RandomPersonality()
	isValid := personality == Hunter || personality == Wimp
	assert.True(t, isValid, "RandomPersonality should return Hunter or Wimp")

}
