# Matrioska

A thumbnails generator server.

## Installation

    go get github.com/pilu/matrioska

## Usage

    PUBLIC_PATH=/path/to/images matrioska

If you have an image like:

    http://localhost:7000/foo.png

You can request a thumbnail with:

    http://localhost:7000/foo-300x250.png

### Geometry format

* WIDTHxHEIGHT
* WIDTHx
* xHEIGHT

### Options

Options are taken from env variables.
Defaults are:

* PUBLIC_PATH (./public)
* GEOMETRIES ("\*")
* PORT (7000)
* HOST (127.0.0.1)

### Restricting valid geometries

    GEOMETRIES="200x300, 500x, x200" matrioska

## Author

Andrea Franz (http://gravityblast.com)


