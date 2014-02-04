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

func TestHost(t *testing.T) {
	InitSettings()
	assert.Equal(t, DefaultHost, Host())

	os.Setenv("HOST", "1.2.3.4")
	InitSettings()
	assert.Equal(t, "1.2.3.4", Host())
}

func TestPort(t *testing.T) {
	InitSettings()
	assert.Equal(t, DefaultPort, Port())

	os.Setenv("PORT", "12345")
	InitSettings()
	assert.Equal(t, "12345", Port())
}
