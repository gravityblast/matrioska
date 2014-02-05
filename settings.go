package main

import (
	"os"
	"regexp"
)

const DefaultPublicPath = "./public"
const DefaultHost = "127.0.0.1"
const DefaultPort = "7000"

var settings map[string]string
var validGeometries map[string]bool


type IsValidGeometryFunc func(geometry string) bool

func getVar(name, defaultValue string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}

	return defaultValue
}

func InitSettings() {
	settings = map[string]string{
		"PublicPath": getVar("PUBLIC_PATH", DefaultPublicPath),
		"Host":       getVar("HOST", DefaultHost),
		"Port":       getVar("PORT", DefaultPort),
		"Geometries": getVar("GEOMETRIES", "*"),
	}

	validGeometries = make(map[string]bool)
	geometries := regexp.MustCompile(`[\s,]+`).Split(settings["Geometries"], -1)
	for _, g := range geometries {
		validGeometries[g] = true
	}
}

func PublicPath() string {
	return settings["PublicPath"]
}

func Host() string {
	return settings["Host"]
}

func Port() string {
	return settings["Port"]
}

func IsValidGeometry(geometry string) bool {
	return validGeometries["*"] || validGeometries[geometry]
}
