package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestParseGeometry(t *testing.T) {
	geometry, err := ParseGeometry("200x300")
	assert.Nil(t, err)
	assert.Equal(t, 200, geometry.Width)
	assert.Equal(t, 300, geometry.Height)

	geometry, err = ParseGeometry("200x")
	assert.Nil(t, err)
	assert.Equal(t, 200, geometry.Width)
	assert.Equal(t, 0, geometry.Height)

	geometry, err = ParseGeometry("x300")
	assert.Nil(t, err)
	assert.Equal(t, 0, geometry.Width)
	assert.Equal(t, 300, geometry.Height)

	geometry, err = ParseGeometry("x")
	assert.NotNil(t, err)
	assert.Type(t, "main.ErrorBadGeometry", err)

	geometry, err = ParseGeometry("bad-geometry")
	assert.NotNil(t, err)
	assert.Type(t, "main.ErrorBadGeometry", err)
}

func TestGeometry_String(t *testing.T) {
	g := Geometry{200, 300}
	assert.Equal(t, "200x300", g.String())

	g = Geometry{200, 0}
	assert.Equal(t, "200x", g.String())
}
