package main

import (
	"errors"
	"fmt"
	"os/exec"
	"path"
	"regexp"
)

type Thumb struct {
	Path     string
	Dir      string
	Name     string
	Ext      string
	Geometry Geometry
	MainPath string
}

var regexpFileSize = regexp.MustCompile(`([^/]+)\-([^\.]+)?.\w+$`)

func (thumb Thumb) MainFullPath() string {
	return path.Join(PublicPath(), thumb.MainPath)
}

func (thumb Thumb) FullPath() string {
	return path.Join(PublicPath(), thumb.Path)
}

func (thumb Thumb) Generate() error {
	if g := thumb.Geometry.String(); !IsValidGeometry(g) {
		return ErrorGeometryNotAllowed{ g }
	}
	sourcePath := thumb.MainFullPath()
	outputPath := thumb.FullPath()

	log("Generating: %s", outputPath)
	cmd := exec.Command("convert", sourcePath, "-geometry", thumb.Geometry.String(), outputPath)

	return cmd.Run()
}

func ParseThumbName(filePath string) (name string, geometry Geometry, err error) {
	fileName := path.Base(filePath)

	if matches := regexpFileSize.FindAllStringSubmatch(fileName, -1); len(matches) > 0 {
		name = matches[0][1]
		geometry, err = ParseGeometry(matches[0][2])
		return
	}

	err = errors.New(fmt.Sprintf("Bad filename %s", filePath))

	return
}

func ThumbFromPath(filePath string) (Thumb, error) {
	thumb := Thumb{
		Path: filePath,
	}

	name, geometry, err := ParseThumbName(filePath)
	if err != nil {
		return thumb, err
	}

	thumb.Geometry = geometry
	thumb.Dir = path.Dir(filePath)
	thumb.Name = name
	thumb.Ext = path.Ext(filePath)
	thumb.MainPath = path.Join(thumb.Dir, fmt.Sprintf("%s%s", thumb.Name, thumb.Ext))

	return thumb, nil
}
