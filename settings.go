package main

import (
	"os"
)

const DefaultPublicPath = "./public"
const DefaultHost = "127.0.0.1"
const DefaultPort = "7000"

var settings map[string]string

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
