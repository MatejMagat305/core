// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paint

import (
	"image"
	"io"
	"log/slog"

	"cogentcore.org/core/math32"
	"cogentcore.org/core/paint/raster"
	"cogentcore.org/core/paint/scan"
)

// The State holds all the current rendering state information used
// while painting -- a viewport just has one of these
type State struct {

	// current transform
	CurrentTransform math32.Matrix2

	// current path
	Path raster.Path

	// rasterizer -- stroke / fill rendering engine from raster
	Raster *raster.Dasher

	// scan scanner
	Scanner *scan.Scanner

	// scan spanner
	ImgSpanner *scan.ImgSpanner

	// starting point, for close path
	Start math32.Vector2

	// current point
	Current math32.Vector2

	// is current point current?
	HasCurrent bool

	// pointer to image to render into
	Image *image.RGBA

	// current mask
	Mask *image.Alpha

	// boundaries to restrict drawing to -- much faster than clip mask for basic square region exclusion -- used for restricting drawing
	Bounds image.Rectangle

	// bounding box of last object rendered -- computed by renderer during Fill or Stroke, grabbed by SVG objects
	LastRenderBBox image.Rectangle

	// stack of transforms
	TransformStack []math32.Matrix2

	// stack of bounds -- every render starts with a push onto this stack, and finishes with a pop
	BoundsStack []image.Rectangle

	// stack of clips, if needed
	ClipStack []*image.Alpha

	// if non-nil, SVG output of paint commands is sent here
	SVGOut io.Writer
}

// Init initializes State -- must be called whenever image size changes
func (rs *State) Init(width, height int, img *image.RGBA) {
	rs.CurrentTransform = math32.Identity2()
	rs.Image = img
	rs.ImgSpanner = scan.NewImgSpanner(img)
	rs.Scanner = scan.NewScanner(rs.ImgSpanner, width, height)
	rs.Raster = raster.NewDasher(width, height, rs.Scanner)
}

// PushTransform pushes current transform onto stack and apply new transform on top of it
// must protect within render mutex lock (see Lock version)
func (rs *State) PushTransform(tf math32.Matrix2) {
	if rs.TransformStack == nil {
		rs.TransformStack = make([]math32.Matrix2, 0)
	}
	rs.TransformStack = append(rs.TransformStack, rs.CurrentTransform)
	rs.CurrentTransform.SetMul(tf)
}

// PopTransform pops transform off the stack and set to current transform
// must protect within render mutex lock (see Lock version)
func (rs *State) PopTransform() {
	sz := len(rs.TransformStack)
	if sz == 0 {
		slog.Error("programmer error: paint.State.PopTransform: stack is empty")
		rs.CurrentTransform = math32.Identity2()
		return
	}
	rs.CurrentTransform = rs.TransformStack[sz-1]
	rs.TransformStack = rs.TransformStack[:sz-1]
}

// PushBounds pushes current bounds onto stack and sets new bounds.
// This is the essential first step in rendering.
// Any further actual rendering should always be surrounded
// by [State.Lock] and [State.Unlock] calls.
func (rs *State) PushBounds(b image.Rectangle) {
	if rs.BoundsStack == nil {
		rs.BoundsStack = make([]image.Rectangle, 0, 100)
	}
	if rs.Bounds.Empty() { // note: method name should be IsEmpty!
		rs.Bounds = rs.Image.Bounds()
	}
	rs.BoundsStack = append(rs.BoundsStack, rs.Bounds)
	rs.Bounds = b
}

// PopBounds pops the bounds off the stack and sets the current bounds.
// This must be equally balanced with corresponding [State.PushBounds] calls.
func (rs *State) PopBounds() {
	sz := len(rs.BoundsStack)
	if sz == 0 {
		slog.Error("programmer error: paint.State.PopBounds: stack is empty")
		rs.Bounds = rs.Image.Bounds()
		return
	}
	rs.Bounds = rs.BoundsStack[sz-1]
	rs.BoundsStack = rs.BoundsStack[:sz-1]
}

// PushClip pushes current Mask onto the clip stack
func (rs *State) PushClip() {
	if rs.Mask == nil {
		return
	}
	if rs.ClipStack == nil {
		rs.ClipStack = make([]*image.Alpha, 0, 10)
	}
	rs.ClipStack = append(rs.ClipStack, rs.Mask)
}

// PopClip pops Mask off the clip stack and set to current mask
func (rs *State) PopClip() {
	sz := len(rs.ClipStack)
	if sz == 0 {
		slog.Error("programmer error: paint.State.PopClip: stack is empty")
		rs.Mask = nil // implied
		return
	}
	rs.Mask = rs.ClipStack[sz-1]
	rs.ClipStack[sz-1] = nil
	rs.ClipStack = rs.ClipStack[:sz-1]
}

// Size returns the size of the underlying image as a [math32.Vector2].
func (rs *State) Size() math32.Vector2 {
	return math32.Vector2FromPoint(rs.Image.Rect.Size())
}
