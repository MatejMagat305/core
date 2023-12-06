// Copyright 2023 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on golang.org/x/exp/shiny:
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import "goki.dev/goosi"

// AppMulti contains the data and logic common to all implementations of [goosi.App]
// on multi-window platforms (desktop), as opposed to single-window
// platforms (mobile, web, and offscreen), for which you should use [AppSingle]. An AppMulti is associated
// with a corresponding type of [goosi.Window]. The [goosi.Window]
// type should embed [WindowMulti].
type AppMulti[W goosi.Window] struct {
	App

	// Windows are the windows associated with the app
	Windows []W

	// Screens are the screens associated with the app
	Screens []*goosi.Screen

	// AllScreens is a unique list of all screens ever seen, from which
	// information can be got if something is missing in [AppMulti.Screens]
	AllScreens []*goosi.Screen

	// CtxWindow is a dynamically set context window used for some operations
	CtxWindow W
}
