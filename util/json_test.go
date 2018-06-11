package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytesToDecoder(t *testing.T) {
	mySlice := []byte("Hello world")
	decoder := BytesToDecoder(mySlice)

	assert.NotNil(t, decoder, "The decoder should build a builder")

}
