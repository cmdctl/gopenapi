package test

import (
	"embed"
	"github.com/getkin/kin-openapi/openapi3"
	"net/url"
)

//go:embed data/v3.0/* data/v3.1/*
var FS embed.FS

// LoadTestSchema loads the test schema from the embedded file system.
// it takes a path to OpenAPI 3 the schema file in json or yaml format.
func LoadTestSchema(path string) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, uri *url.URL) ([]byte, error) {
		return FS.ReadFile(uri.Path)
	}
	return loader.LoadFromFile(path)
}
