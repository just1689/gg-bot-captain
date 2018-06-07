package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytesToDecoder(t *testing.T) {
	str := "Hello world"
	mySlice := []byte(str)
	decoder := BytesToDecoder(mySlice)

	assert.NotNil(t, decoder, "The decoder should build a builder")

}
