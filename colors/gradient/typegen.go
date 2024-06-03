// Code generated by "core generate"; DO NOT EDIT.

package gradient

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/colors/gradient.Base", IDName: "base", Doc: "Base contains the data and logic common to all gradient types.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Fields: []types.Field{{Name: "Stops", Doc: "the stops for the gradient; use AddStop to add stops"}, {Name: "Spread", Doc: "the spread method used for the gradient if it stops before the end"}, {Name: "Blend", Doc: "the colorspace algorithm to use for blending colors"}, {Name: "Units", Doc: "the units to use for the gradient"}, {Name: "Box", Doc: "the bounding box of the object with the gradient; this is used when rendering\ngradients with [Units] of [ObjectBoundingBox]."}, {Name: "Transform", Doc: "Transform is the gradient's own transformation matrix applied to the gradient's points.\nThis is a property of the Gradient itself."}, {Name: "Opacity", Doc: "Opacity is the overall object opacity multiplier, applied in conjunction with the\nstop-level opacity blending."}, {Name: "ApplyFuncs", Doc: "ApplyFuncs contains functions that are applied to the color after gradient color is generated.\nThis allows for efficient StateLayer and other post-processing effects\nto be applied.  The Applier handles other cases, but gradients always\nmust have the Update function called at render time, so they must\nremain Gradient types."}, {Name: "boxTransform", Doc: "boxTransform is the Transform applied to the bounding Box,\nonly for [Units] == [ObjectBoundingBox]."}, {Name: "stopsRGB", Doc: "stopsRGB are the computed RGB stops for blend types other than RGB"}, {Name: "stopsRGBSrc", Doc: "stopsRGBSrc are the source Stops when StopsRGB were last computed"}}})

// SetSpread sets the [Base.Spread]:
// the spread method used for the gradient if it stops before the end
func (t *Base) SetSpread(v Spreads) *Base { t.Spread = v; return t }

// SetBlend sets the [Base.Blend]:
// the colorspace algorithm to use for blending colors
func (t *Base) SetBlend(v colors.BlendTypes) *Base { t.Blend = v; return t }

// SetUnits sets the [Base.Units]:
// the units to use for the gradient
func (t *Base) SetUnits(v Units) *Base { t.Units = v; return t }

// SetBox sets the [Base.Box]:
// the bounding box of the object with the gradient; this is used when rendering
// gradients with [Units] of [ObjectBoundingBox].
func (t *Base) SetBox(v math32.Box2) *Base { t.Box = v; return t }

// SetTransform sets the [Base.Transform]:
// Transform is the gradient's own transformation matrix applied to the gradient's points.
// This is a property of the Gradient itself.
func (t *Base) SetTransform(v math32.Matrix2) *Base { t.Transform = v; return t }

// SetOpacity sets the [Base.Opacity]:
// Opacity is the overall object opacity multiplier, applied in conjunction with the
// stop-level opacity blending.
func (t *Base) SetOpacity(v float32) *Base { t.Opacity = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/colors/gradient.Linear", IDName: "linear", Doc: "Linear represents a linear gradient. It implements the [image.Image] interface.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "Base"}}, Fields: []types.Field{{Name: "Start", Doc: "the starting point of the gradient (x1 and y1 in SVG)"}, {Name: "End", Doc: "the ending point of the gradient (x2 and y2 in SVG)"}, {Name: "rStart", Doc: "computed current render versions transformed by object matrix"}, {Name: "rEnd"}, {Name: "distance"}, {Name: "distanceLengthSquared"}}})

// SetStart sets the [Linear.Start]:
// the starting point of the gradient (x1 and y1 in SVG)
func (t *Linear) SetStart(v math32.Vector2) *Linear { t.Start = v; return t }

// SetEnd sets the [Linear.End]:
// the ending point of the gradient (x2 and y2 in SVG)
func (t *Linear) SetEnd(v math32.Vector2) *Linear { t.End = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/colors/gradient.Radial", IDName: "radial", Doc: "Radial represents a radial gradient. It implements the [image.Image] interface.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "Base"}}, Fields: []types.Field{{Name: "Center", Doc: "the center point of the gradient (cx and cy in SVG)"}, {Name: "Focal", Doc: "the focal point of the gradient (fx and fy in SVG)"}, {Name: "Radius", Doc: "the radius of the gradient (rx and ry in SVG)"}, {Name: "rCenter", Doc: "current render version -- transformed by object matrix"}, {Name: "rFocal", Doc: "current render version -- transformed by object matrix"}, {Name: "rRadius", Doc: "current render version -- transformed by object matrix"}}})

// SetCenter sets the [Radial.Center]:
// the center point of the gradient (cx and cy in SVG)
func (t *Radial) SetCenter(v math32.Vector2) *Radial { t.Center = v; return t }

// SetFocal sets the [Radial.Focal]:
// the focal point of the gradient (fx and fy in SVG)
func (t *Radial) SetFocal(v math32.Vector2) *Radial { t.Focal = v; return t }

// SetRadius sets the [Radial.Radius]:
// the radius of the gradient (rx and ry in SVG)
func (t *Radial) SetRadius(v math32.Vector2) *Radial { t.Radius = v; return t }
