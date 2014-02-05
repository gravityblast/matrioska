package main

import (
	"fmt"
	assert "github.com/pilu/miniassert"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestMainHandler_Favicon(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/favicon.ico", nil)
	MainHandler(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestMainHandler_OriginalFileExists(t *testing.T) {
	recorder := httptest.NewRecorder()
	InitSettings()
	settings["PublicPath"] = "./test_fixtures"
	request, _ := http.NewRequest("GET", "/test.png", nil)
	MainHandler(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestMainHandler_ThumbExists(t *testing.T) {
	recorder := httptest.NewRecorder()
	InitSettings()
	settings["PublicPath"] = "./test_fixtures"
	request, _ := http.NewRequest("GET", "/test-50x50.png", nil)
	MainHandler(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestMainHandler_GenerateThumb(t *testing.T) {
	InitSettings()
	fixturesPath := "./test_fixtures"
	settings["PublicPath"] = fixturesPath
	fileName := "test-100x100.png"
	filePath := path.Join(fixturesPath, fileName)

	// we don't have a thumb yet
	_, err := os.Stat(filePath)
	assert.NotNil(t, err)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", fmt.Sprintf("/%s", fileName), nil)
	MainHandler(recorder, request)

	// check if the thumb has been generated
	assert.Equal(t, http.StatusOK, recorder.Code)
	_, err = os.Stat(filePath)
	assert.Nil(t, err)

	// remove the generated thumb
	os.Remove(filePath)
}

func TestMainHandler_BadGeometry(t *testing.T) {
	InitSettings()
	settings["PublicPath"] = "./test_fixtures"
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "test-axb.png", nil)
	MainHandler(recorder, request)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestMainHandler_MainFileNotFound(t *testing.T) {
	InitSettings()
	settings["PublicPath"] = "./test_fixtures"
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "foo-100x100.png", nil)
	MainHandler(recorder, request)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestMainHandler_InvalidGeomtries(t *testing.T) {
	InitSettings()
	settings["PublicPath"] = "./test_fixtures"
	validGeometries = map[string]bool{ "200x200": true}

	// only 200x200 is allowed
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "test-10x10.png", nil)
	MainHandler(recorder, request)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	validGeometries = map[string]bool{ "*": true}
}
