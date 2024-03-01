// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package giv

import (
	"fmt"

	"cogentcore.org/core/gi"
	"cogentcore.org/core/ki"
	"cogentcore.org/core/laser"
	"cogentcore.org/core/strcase"
	"cogentcore.org/core/styles"
)

// ArgView represents a slice of reflect.Value's and associated names, for the
// purpose of supplying arguments to methods called via the MethodView
// framework.
type ArgView struct {
	gi.Frame

	// Args are the args that we are a view onto
	Args []Value

	// a record of parent View names that have led up to this view -- displayed as extra contextual information in view dialog windows
	ViewPath string
}

func (av *ArgView) OnInit() {
	av.Frame.OnInit()
	av.SetStyles()
}

func (av *ArgView) SetStyles() {
	av.Style(func(s *styles.Style) {
		s.Direction = styles.Column
		s.Grow.Set(1, 1)
	})
	av.OnWidgetAdded(func(w gi.Widget) {
		switch w.PathFrom(av) {
		case "title":
			title := w.(*gi.Label)
			title.Type = gi.LabelTitleLarge
			title.Style(func(s *styles.Style) {
				s.Grow.Set(1, 0)
				s.Text.Align = styles.Center
			})
		case "args-grid":
			w.Style(func(s *styles.Style) {
				s.Display = styles.Grid
				s.Columns = 2
				s.Min.X.Ch(60)
				s.Min.Y.Em(10)
				s.Grow.Set(1, 1)
				s.Overflow.Set(styles.OverflowAuto)
			})
		}
		if w.Parent().Name() == "args-grid" {
			w.Style(func(s *styles.Style) {
			})
		}
	})
}

// ConfigWidget configures the view
func (av *ArgView) ConfigWidget() {
	config := ki.Config{}
	config.Add(gi.LabelType, "title")
	config.Add(gi.FrameType, "args-grid")
	mods, updt := av.ConfigChildren(config)
	av.ConfigArgsGrid()
	if mods {
		av.UpdateEnd(updt)
	}
}

// TitleWidget returns the title label widget
func (av *ArgView) TitleWidget() *gi.Label {
	return av.ChildByName("title", 0).(*gi.Label)
}

// ArgsGrid returns the grid layout widget, which contains all the fields
// and values
func (av *ArgView) ArgsGrid() *gi.Frame {
	return av.ChildByName("args-grid", 0).(*gi.Frame)
}

// ConfigArgsGrid configures the ArgsGrid for the current struct
func (av *ArgView) ConfigArgsGrid() {
	if laser.AnyIsNil(av.Args) {
		return
	}
	sg := av.ArgsGrid()
	config := ki.Config{}
	argnms := make(map[string]bool)
	for i := range av.Args {
		arg := av.Args[i]
		if view, _ := arg.Tag("view"); view == "-" {
			continue
		}
		vtyp := arg.WidgetType()
		knm := strcase.ToKebab(arg.Name())
		if _, has := argnms[knm]; has {
			knm += fmt.Sprintf("%d", i)
		}
		argnms[knm] = true
		labnm := "label-" + knm
		valnm := "value-" + knm
		config.Add(gi.LabelType, labnm)
		config.Add(vtyp, valnm)
	}
	mods, updt := sg.ConfigChildren(config) // not sure if always unique?
	if mods {
		av.SetNeedsLayout(updt)
	} else {
		updt = sg.UpdateStart()
	}
	idx := 0
	for i := range av.Args {
		arg := av.Args[i]
		if view, _ := arg.Tag("view"); view == "-" {
			continue
		}
		arg.SetTag("grow", "1")
		lbl := sg.Child(idx * 2).(*gi.Label)
		lbl.Text = arg.Label()
		lbl.Tooltip = arg.Doc()
		w, wb := gi.AsWidget(sg.Child((idx * 2) + 1))
		if wb.Prop("configured") == nil {
			wb.SetProp("configured", true)
			arg.ConfigWidget(w)
		} else {
			arg.AsValueBase().Widget = w
			arg.UpdateWidget()
		}
		idx++
	}
	sg.UpdateEnd(updt)
}

// UpdateArgs updates each of the value-view widgets for the args
func (av *ArgView) UpdateArgs() {
	updt := av.UpdateStart()
	for i := range av.Args {
		ad := av.Args[i]
		ad.UpdateWidget()
	}
	av.UpdateEnd(updt)
}
