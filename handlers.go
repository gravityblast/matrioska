package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func SkipFavicon(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusNotFound)
		return true
	}

	return false
}

func FileExist(w http.ResponseWriter, r *http.Request) bool {
	path := filepath.Join(PublicPath(), r.URL.Path)
	if info, err := os.Stat(path); err == nil && !info.IsDir() {
		log("Found: %s", r.URL.Path)
		http.ServeFile(w, r, path)
		return true
	}

	return false
}

func MainFileNotFound(path string, w http.ResponseWriter, r *http.Request) bool {
	if _, err := os.Stat(path); err != nil {
		log("Error: %s", err.Error())
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "file not found")
		return true
	}

	return false
}

func Generate(thumb *Thumb, w http.ResponseWriter) {
	image, err := thumb.Generate()

	if image != nil {
		defer image.Destroy()
	}

	if err != nil {
		log("Error: %s", err.Error())

		switch err.(type) {
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		case ErrorGeometryNotAllowed:
			http.Error(w, "geometry not allowed", http.StatusBadRequest)
		}

		return
	}

	blob, err := image.ToBlob("jpg")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(blob)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if SkipFavicon(w, r) {
		return
	}

	if FileExist(w, r) {
		return
	}

	thumb, err := ThumbFromPath(r.URL.Path)

	if err != nil {
		log("Error: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if MainFileNotFound(thumb.MainFullPath(), w, r) {
		return
	}

	Generate(thumb, w)
}
