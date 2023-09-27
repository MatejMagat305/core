// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package colors

import (
	"image"
	"image/color"
	"log/slog"

	"github.com/srwiley/rasterx"
	"goki.dev/mat32/v2"
)

// Full represents a fully specified color that can either be a solid color or
// a gradient. If Gradient is nil, it is a solid color; otherwise, it is a gradient.
// Solid should typically be set using the [Full.SetSolid] method to
// ensure that Gradient is nil and thus Solid will be taken into account.
type Full struct {
	Solid    color.RGBA
	Gradient *Gradient
}

// IsNil returns whether the color is nil, checking both the gradient
// and the solid color.
func (f *Full) IsNil() bool {
	return f.Gradient == nil && IsNil(f.Solid)
}

// SolidOrNil returns the solid color if it is not non-nil, or nil otherwise.
// It is should be used by consumers that explicitly handle nil colors.
func (f *Full) SolidOrNil() color.Color {
	if IsNil(f.Solid) {
		return nil
	}
	return f.Solid
}

// SetSolid sets the color to the given solid color,
// also setting the gradient to nil.
func (f *Full) SetSolid(solid color.RGBA) {
	f.Solid = solid
	f.Gradient = nil
}

// SetSolid sets the color to the given solid [color.Color],
// also setting the gradient to nil.
func (f *Full) SetColor(clr color.Color) {
	f.Solid = AsRGBA(clr)
	f.Gradient = nil
}

// SetSolid sets the color to the solid color with the given name,
// also setting the gradient to nil.
func (f *Full) SetName(name string) {
	f.Solid = LogFromName(name)
	f.Gradient = nil
}

// RenderColor returns the color or [rasterx.ColorFunc] for rendering, applying
// the given opacity and bounds.
func (f *Full) RenderColor(opacity float32, bounds image.Rectangle, xform mat32.Mat2) any {
	if f.Gradient == nil {
		return rasterx.ApplyOpacity(f.Solid, float64(opacity))
	}
	return f.Gradient.RenderColor(opacity, bounds, xform)
}

// SetAny sets the color from the given value of any type.
// It handles values of types [color.Color], [*Gradient],
// and string. It returns any error; see [MustFromAny] and
// [LogFromAny] for versions that do not return an error.
func (f *Full) SetAny(val any, base color.Color) error {
	switch valv := val.(type) {
	case color.Color:
		f.Solid = AsRGBA(valv)
	case *Gradient:
		*f.Gradient = *valv
	case string:
		f.SetString(valv, base)
	}
	return nil
}

// MustSetAny sets the color from the given value of any type.
// It panics on any resulting error; see [Full.SetAny]
// for more information and a version that returns an error.
func (f *Full) MustSetAny(val any, base color.Color) {
	err := f.SetAny(val, base)
	if err != nil {
		panic("colors.Full.MustSetAny: " + err.Error())
	}
}

// LogSetAny sets the color value from the given value of any type.
// It logs any resulting error; see [Full.SetAny]
// for more information and a version that returns an error.
func (f *Full) LogSetAny(val any, base color.Color) {
	err := f.SetAny(val, base)
	if err != nil {
		slog.Error("colors.Full.LogSetAny: " + err.Error())
	}
}
