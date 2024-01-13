// Copyright (c) 2023, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtigen

import (
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"goki.dev/gti"
)

// TypeTmpl is the template for [gti.Type] declarations.
// It takes a [*Type] as its value.
var TypeTmpl = template.Must(template.New("Type").
	Funcs(template.FuncMap{
		"GtiTypeOf": GtiTypeOf,
	}).Parse(
	`
	{{if .Config.TypeVar}} // {{.LocalName}}Type is the [gti.Type] for [{{.LocalName}}]
	var {{.LocalName}}Type {{else}} var _ {{end}} = gti.AddType(&gti.Type
		{{- $typ := GtiTypeOf . -}}
		{{- printf "%#v" $typ -}}
	)
	`))

// GtiTypeOf converts the given [*Type] to a [*gti.Type]
func GtiTypeOf(typ *Type) *gti.Type {
	cp := typ.Type
	res := &cp
	res.Fields = typ.Fields.Fields
	res.Embeds = typ.Embeds.Fields
	if typ.Config.Instance {
		typ.Instance = "&" + typ.LocalName + "{}"
	}
	return res
}

// FuncTmpl is the template for [gti.Func] declarations.
// It takes a [*gti.Func] as its value.
var FuncTmpl = template.Must(template.New("Func").Parse(
	`
	var _ = gti.AddFunc(&gti.Func{
		Name: "{{.Name}}",
		Doc: {{printf "%q" .Doc}},
		Directives: {{printf "%#v" .Directives}},
		Args: {{printf "%#v" .Args}},
		Returns: {{printf "%#v" .Returns}},
	})
	`))

// SetterMethodsTmpl is the template for setter methods for a type.
// It takes a [*Type] as its value.
var SetterMethodsTmpl = template.Must(template.New("SetterMethods").
	Funcs(template.FuncMap{
		"SetterFields": SetterFields,
		"ToCamel":      strcase.ToCamel,
		"SetterType":   SetterType,
		"DocToComment": DocToComment,
	}).Parse(
	`
	{{$typ := .}}
	{{range (SetterFields .)}}
	// Set{{ToCamel .Name}} sets the [{{$typ.LocalName}}.{{.Name}}] {{- if ne .Doc ""}}:{{end}}
	{{DocToComment .Doc}}
	func (t *{{$typ.LocalName}}) Set{{ToCamel .Name}}(v {{SetterType .LocalType}}) *{{$typ.LocalName}} {
		t.{{.Name}} = v
		return t
	}
	{{end}}
`))

// SetterFields returns all of the fields and embedded fields of the given type
// that don't have a `set:"-"` struct tag.
func SetterFields(typ *Type) []gti.Field {
	res := []gti.Field{}
	do := func(fields Fields) {
		for _, f := range fields.Fields {
			// unspecified indicates to add a set method; only "-" means no set
			hasSetter := fields.Tags[f.Name].Get("set") != "-"
			if hasSetter {
				res = append(res, f)
			}
		}
	}
	do(typ.Fields)
	do(typ.EmbeddedFields)
	return res
}

// SetterType converts the given local type name to one that can be used
// in a setter by converting slices to variadic arguments.
func SetterType(typ string) string {
	if strings.HasPrefix(typ, "[]") {
		return "..." + strings.TrimPrefix(typ, "[]")
	}
	return typ
}

// DocToComment converts the given doc string to an appropriate comment string.
func DocToComment(doc string) string {
	return "// " + strings.ReplaceAll(doc, "\n", "\n// ")
}
