// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// adapted from: https://github.com/material-foundation/material-color-utilities
// Copyright 2022 Google LLC
// Licensed under the Apache License, Version 2.0 (the "License")

package hct

import (
	"github.com/goki/cam/cam16"
	"github.com/goki/cam/cie"
)

// HCT, hue, chroma, and tone. A color system that provides a perceptually
// accurate color measurement system that can also accurately render what
// colors will appear as in different lighting environments.
type HCT struct {

	// [min: 0] [max: 360] hue (h) is the spectral identity of the color (red, green, blue etc) in degrees (0-360)
	Hue float32 `min:"0" max:"360" desc:"hue (h) is the spectral identity of the color (red, green, blue etc) in degrees (0-360)"`

	// chroma (C) is the colorfulness or saturation of the color -- greyscale colors have no chroma, and fully saturated ones have high chroma.  The maximum varies as a function of hue and tone.
	Chroma float32 `desc:"chroma (C) is the colorfulness or saturation of the color -- greyscale colors have no chroma, and fully saturated ones have high chroma.  The maximum varies as a function of hue and tone."`

	// [min: 0] [max: 100] tone is the L* component from the LAB (L*a*b*) color system, which is linear in human perception of lightness
	Tone float32 `min:"0" max:"100" desc:"tone is the L* component from the LAB (L*a*b*) color system, which is linear in human perception of lightness"`

	// sRGB standard gamma-corrected 0-1 normalized RGB representation of the color
	R, G, B float32 `desc:"sRGB standard gamma-corrected 0-1 normalized RGB representation of the color"`
}

// NewHCT returns a new HCT representation for given parameters:
// hue = 0..360
// chroma = 0..? depends on other params
// tone = 0..100
// also computes and sets the sRGB normalized, gamma corrected R,G,B values
// while keeping the sRGB representation within its gamut,
// which may cause the chroma to decrease until it is inside the gamut.
func NewHCT(hue, chroma, tone float32) HCT {
	r, g, b := SolveToRGB(hue, chroma, tone)
	return SRGBToHCT(r, g, b)
}

// SetHue sets the hue of this color. Chroma may decrease because chroma has a
// different maximum for any given hue and tone.
// 0 <= hue < 360; invalid values are corrected.
func (h *HCT) SetHue(hue float32) {
	r, g, b := SolveToRGB(hue, h.Chroma, h.Tone)
	*h = SRGBToHCT(r, g, b)
}

// SetChroma sets the chroma of this color (0 to max that depends on other params),
// while keeping the sRGB representation within its gamut,
// which may cause the chroma to decrease until it is inside the gamut.
func (h *HCT) SetChroma(chroma float32) {
	r, g, b := SolveToRGB(h.Hue, chroma, h.Tone)
	*h = SRGBToHCT(r, g, b)
}

// SetTone sets the tone of this color (0 < tone < 100),
// while keeping the sRGB representation within its gamut,
// which may cause the chroma to decrease until it is inside the gamut.
func (h *HCT) SetTone(tone float32) {
	r, g, b := SolveToRGB(h.Hue, h.Chroma, tone)
	*h = SRGBToHCT(r, g, b)
}

// SRGBToCAM returns CAM values from given SRGB color coordinates,
// under standard viewing conditions.  The RGB value range is 0-1,
// and RGB values have gamma correction.
func SRGBToHCT(r, g, b float32) HCT {
	x, y, z := cie.SRGBToXYZ(r, g, b)
	cam := cam16.XYZToCAM(100*x, 100*y, 100*z)
	l, _, _ := cie.XYZToLAB(x, y, z)
	return HCT{Hue: cam.Hue, Chroma: cam.Chroma, Tone: l, R: r, G: g, B: b}
}

/*
  // Translate a color into different [ViewingConditions].
  //
  // Colors change appearance. They look different with lights on versus off,
  // the same color, as in hex code, on white looks different when on black.
  // This is called color relativity, most famously explicated by Josef Albers
  // in Interaction of Color.
  //
  // In color science, color appearance models can account for this and
  // calculate the appearance of a color in different settings. HCT is based on
  // CAM16, a color appearance model, and uses it to make these calculations.
  //
  // See [ViewingConditions.make] for parameters affecting color appearance.
  Hct inViewingConditions(ViewingConditions vc) {
    // 1. Use CAM16 to find XYZ coordinates of color in specified VC.
    final cam16 = Cam16.fromInt(toInt());
    final viewedInVc = cam16.xyzInViewingConditions(vc);

    // 2. Create CAM16 of those XYZ coordinates in default VC.
    final recastInVc = Cam16.fromXyzInViewingConditions(
      viewedInVc[0],
      viewedInVc[1],
      viewedInVc[2],
      ViewingConditions.make(),
    );

    // 3. Create HCT from:
    // - CAM16 using default VC with XYZ coordinates in specified VC.
    // - L* converted from Y in XYZ coordinates in specified VC.
    final recastHct = Hct.from(
      recastInVc.hue,
      recastInVc.chroma,
      ColorUtils.lstarFromY(viewedInVc[1]),
    );
    return recastHct;
  }
}

*/
