// Code generated by "goki generate"; DO NOT EDIT.

package xyzv

import (
	"goki.dev/gi/v2/gi"
	"goki.dev/goosi/events"
	"goki.dev/gti"
	"goki.dev/ki/v2"
	"goki.dev/ordmap"
)

// Scene3DType is the [gti.Type] for [Scene3D]
var Scene3DType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gi/v2/xyzv.Scene3D",
	ShortName:  "xyzv.Scene3D",
	IDName:     "scene-3-d",
	Doc:        "Scene3D contains a svg.SVG element.\nThe rendered version is cached for a given size.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Scene", &gti.Field{Name: "Scene", Type: "*goki.dev/xyz.Scene", LocalType: "*xyz.Scene", Doc: "Scene is the 3D Scene", Directives: gti.Directives{}, Tag: "set:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"WidgetBase", &gti.Field{Name: "WidgetBase", Type: "goki.dev/gi/v2/gi.WidgetBase", LocalType: "gi.WidgetBase", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &Scene3D{},
})

// NewScene3D adds a new [Scene3D] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewScene3D(par ki.Ki, name ...string) *Scene3D {
	return par.NewChild(Scene3DType, name...).(*Scene3D)
}

// KiType returns the [*gti.Type] of [Scene3D]
func (t *Scene3D) KiType() *gti.Type {
	return Scene3DType
}

// New returns a new [*Scene3D] value
func (t *Scene3D) New() ki.Ki {
	return &Scene3D{}
}

// SetTooltip sets the [Scene3D.Tooltip]
func (t *Scene3D) SetTooltip(v string) *Scene3D {
	t.Tooltip = v
	return t
}

// SetClass sets the [Scene3D.Class]
func (t *Scene3D) SetClass(v string) *Scene3D {
	t.Class = v
	return t
}

// SetPriorityEvents sets the [Scene3D.PriorityEvents]
func (t *Scene3D) SetPriorityEvents(v []events.Types) *Scene3D {
	t.PriorityEvents = v
	return t
}

// SetCustomContextMenu sets the [Scene3D.CustomContextMenu]
func (t *Scene3D) SetCustomContextMenu(v func(m *gi.Scene)) *Scene3D {
	t.CustomContextMenu = v
	return t
}
