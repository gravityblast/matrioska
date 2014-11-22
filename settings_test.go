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

func TestIsValidGeometry(t *testing.T) {
	validGeometries = map[string]bool{}
	assert.False(t, IsValidGeometry("200x200"))

	validGeometries = map[string]bool{"*": true}
	assert.True(t, IsValidGeometry("200x200"))

	validGeometries = map[string]bool{"200x100": true, "200x": true, "x100": true}
	assert.False(t, IsValidGeometry("200x200"))
	assert.True(t, IsValidGeometry("200x100"))
	assert.True(t, IsValidGeometry("200x"))
	assert.True(t, IsValidGeometry("x100"))
}
