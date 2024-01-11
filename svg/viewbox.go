// Copyright (c) 2018, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"goki.dev/mat32"
)

////////////////////////////////////////////////////////////////////////////////////////
// ViewBox defines the SVG viewbox

// ViewBox is used in SVG to define the coordinate system
type ViewBox struct {

	// offset or starting point in parent Viewport2D
	Min mat32.Vec2

	// size of viewbox within parent Viewport2D
	Size mat32.Vec2

	// how to scale the view box within parent
	PreserveAspectRatio ViewBoxPreserveAspectRatio
}

// todo: need to implement the viewbox preserve aspect ratio logic!

// Defaults returns viewbox to defaults
func (vb *ViewBox) Defaults() {
	vb.Min = mat32.Vec2{}
	vb.Size = mat32.Vec2{}
	vb.PreserveAspectRatio.Align = NoAlign
	vb.PreserveAspectRatio.MeetOrSlice = Meet
}

// todo: these should be regular ints and use bitflag etc.

// ViewBoxAlign defines values for the PreserveAspectRatio alignment factor
type ViewBoxAlign int32

const (
	NoAlign ViewBoxAlign = 1 << iota          // do not preserve uniform scaling
	XMin                                      // align ViewBox.Min with smallest values of Viewport
	XMid                                      // align ViewBox.Min with midpoint values of Viewport
	XMax                                      // align ViewBox.Min+Size with maximum values of Viewport
	XMask   ViewBoxAlign = XMin + XMid + XMax // mask for X values -- clear all X before setting new one
	YMin    ViewBoxAlign = 1 << iota          // align ViewBox.Min with smallest values of Viewport
	YMid                                      // align ViewBox.Min with midpoint values of Viewport
	YMax                                      // align ViewBox.Min+Size with maximum values of Viewport
	YMask   ViewBoxAlign = YMin + YMid + YMax // mask for Y values -- clear all Y before setting new one
)

// ViewBoxMeetOrSlice defines values for the PreserveAspectRatio meet or slice factor
type ViewBoxMeetOrSlice int32 //enums:enum

const (
	// Meet means the entire ViewBox is visible within Viewport, and it is
	// scaled up as much as possible to meet the align constraints
	Meet ViewBoxMeetOrSlice = iota

	// Slice means the entire ViewBox is covered by the ViewBox, and the
	// ViewBox is scaled down as much as possible, while still meeting the
	// align constraints
	Slice
)

// ViewBoxPreserveAspectRatio determines how to scale the view box within parent Viewport2D
type ViewBoxPreserveAspectRatio struct {

	// how to align x,y coordinates within viewbox
	Align ViewBoxAlign `svg:"align"`

	// how to scale the view box relative to the viewport
	MeetOrSlice ViewBoxMeetOrSlice `svg:"meetOrSlice"`
}
