package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var regexpGeometry = regexp.MustCompile(`^(\d+)?x(\d+)?$`)

type Geometry struct {
	Width  int
	Height int
}

func (g *Geometry) String() string {
	s := "x"
	if g.Width != 0 {
		s = strconv.Itoa(g.Width) + s
	}

	if g.Height != 0 {
		s += strconv.Itoa(g.Height)
	}

	return s
}

type ErrorBadGeometry struct {
	geometryString string
}

func (e ErrorBadGeometry) Error() string {
	return fmt.Sprintf("bad geometry %s", e.geometryString)
}

type ErrorGeometryNotAllowed struct {
	geometryString string
}

func (e ErrorGeometryNotAllowed) Error() string {
	return fmt.Sprintf("geometry not allowed %s", e.geometryString)
}

func ParseGeometry(s string) (Geometry, error) {
	geometry := Geometry{}
	matches := regexpGeometry.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return geometry, ErrorBadGeometry{s}
	}

	geometry.Width, _ = strconv.Atoi(matches[0][1])
	geometry.Height, _ = strconv.Atoi(matches[0][2])

	if geometry.Width == 0 && geometry.Height == 0 {
		return geometry, ErrorBadGeometry{s}
	}

	return geometry, nil
}
