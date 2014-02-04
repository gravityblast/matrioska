package main

import (
	assert "github.com/pilu/miniassert"
	"os"
	"testing"
)

func TestPublicPath(t *testing.T) {
	InitSettings()
	assert.Equal(t, DefaultPublicPath, PublicPath())

	os.Setenv("PUBLIC_PATH", "./foo")
	InitSettings()
	assert.Equal(t, "./foo", PublicPath())
}
