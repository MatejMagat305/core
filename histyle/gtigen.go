// Code generated by "goki generate -add-types"; DO NOT EDIT.

package histyle

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:       "goki.dev/gi/v2/histyle.Trilean",
	Doc:        "Trilean value for StyleEntry value inheritance.",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "goki.dev/gi/v2/histyle.StyleEntry",
	Doc:        "StyleEntry is one value in the map of highlight style values",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Color", &gti.Field{Name: "Color", Type: "color.RGBA", Doc: "text color", Directives: gti.Directives{}}},
		{"Background", &gti.Field{Name: "Background", Type: "color.RGBA", Doc: "background color", Directives: gti.Directives{}}},
		{"Border", &gti.Field{Name: "Border", Type: "color.RGBA", Doc: "[view: -] border color? not sure what this is -- not really used", Directives: gti.Directives{}}},
		{"Bold", &gti.Field{Name: "Bold", Type: "Trilean", Doc: "bold font", Directives: gti.Directives{}}},
		{"Italic", &gti.Field{Name: "Italic", Type: "Trilean", Doc: "italic font", Directives: gti.Directives{}}},
		{"Underline", &gti.Field{Name: "Underline", Type: "Trilean", Doc: "underline", Directives: gti.Directives{}}},
		{"NoInherit", &gti.Field{Name: "NoInherit", Type: "bool", Doc: "don't inherit these settings from sub-category or category levels -- otherwise everything with a Pass is inherited", Directives: gti.Directives{}}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "goki.dev/gi/v2/histyle.Style",
	Doc:        "Style is a full style map of styles for different token.Tokens tag values",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "goki.dev/gi/v2/histyle.Styles",
	Doc:        "Styles is a collection of styles",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
