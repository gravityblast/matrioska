package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestThumbFromPath(t *testing.T) {
	path := "/path/to/image/foo-20x40.png"
	thumb, err := ThumbFromPath(path)

	assert.Nil(t, err)
	assert.Equal(t, path, thumb.Path)
	assert.Equal(t, "/path/to/image", thumb.Dir)
	assert.Equal(t, "foo", thumb.Name)
	assert.Equal(t, ".png", thumb.Ext)
	assert.Equal(t, "/path/to/image/foo.png", thumb.MainPath)
}

func TestParseThumbName(t *testing.T) {
	path := "foo-20x40.png"
	name, geometry, err := ParseThumbName(path)

	assert.Nil(t, err)
	assert.Equal(t, "foo", name)
	assert.Equal(t, 20, geometry.Width)
	assert.Equal(t, 40, geometry.Height)

	path = "foo20x40.png" //bad filename
	name, geometry, err = ParseThumbName(path)

	assert.NotNil(t, err)
	assert.Equal(t, "", name)
	assert.Equal(t, 0, geometry.Width)
	assert.Equal(t, 0, geometry.Height)
}
