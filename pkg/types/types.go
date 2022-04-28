package types

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"go/format"
	"io/ioutil"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

//go:embed types.go.tmpl
// Tmpl is the template for the generated code.
var Tmpl string

type TypeResolver = func(schemaType string, schemaRef string) (langType string)

var typeResolvers = map[string]TypeResolver{
	"golang": resolveGolangTypes,
}

// Property represents a property of a struct.
type Property struct {
	Name        string
	NameJSON    string
	Type        string
	Ref         string
	Description string
}

// Type represents a struct type.
type Type struct {
	Name        string
	Description string
	Ref         string
	Properties  []Property
}

// GeneratedFile represents a generated file.
type GeneratedFile struct {
	PackageName string
	Types       []Type
}

// Params for the Generate function.
type Params struct {
	PackageName string
	OutputFile  string
	SpecFile    string
	Lang        string
}

// Generate generates go source code with defined types from OpenAPI 3 spec file path in the given output file path.
func Generate(params Params) error {
	generatedFile := &GeneratedFile{
		PackageName: params.PackageName,
	}
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	doc, err := loader.LoadFromFile(params.SpecFile)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to load spec file: %s\n", err)
	}

	schemas := doc.Components.Schemas
	types := loadFromSchemas(schemas, params.Lang)
	generatedFile.Types = types
	if err != nil {
		return err
	}
	t, err := template.New("types").Parse(Tmpl)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, generatedFile)
	if err != nil {
		return err
	}
	source, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(params.OutputFile, source, 0644)
	return err
}

func loadFromSchemas(schemas openapi3.Schemas, lang string) []Type {
	var types []Type
	for k, v := range schemas {
		t := &Type{
			Name:        capitalize(k),
			Description: v.Value.Description,
			Ref:         v.Ref,
		}
		for pk, pv := range v.Value.Properties {
			resolveType, ok := typeResolvers[lang]
			if !ok {
				resolveType = resolveGolangTypes
			}
			pType := resolveType(pv.Value.Type, pv.Ref)
			t.Properties = append(t.Properties, Property{
				Name:        capitalize(pk),
				NameJSON:    pk,
				Type:        pType,
				Ref:         pv.Ref,
				Description: pv.Value.Description,
			})
		}
		types = append(types, *t)
	}
	return types
}

func resolveGolangTypes(t string, ref string) string {
	switch t {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		return fmt.Sprintf("[]%s", capitalize(strings.Split(ref, "/")[len(strings.Split(ref, "/"))-1]))
	case "object":
		return capitalize(strings.Split(ref, "/")[len(strings.Split(ref, "/"))-1])
	default:
		return t
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}
