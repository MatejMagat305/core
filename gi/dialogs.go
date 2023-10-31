// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"log/slog"

	"goki.dev/colors"
	"goki.dev/gi/v2/keyfun"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/goosi/events"
	"goki.dev/gti"
)

var (
	// standard vertical space between elements in a dialog, in Ex units
	StdDialogVSpace = float32(1)

	StdDialogVSpaceUnits = units.Ex(StdDialogVSpace)
)

// Dialog is a scene with methods for configuring a dialog
type Dialog struct {
	Scene

	// Stage is the main stage associated with the dialog
	Stage *MainStage

	// Data has arbitrary data for this dialog
	// Data any

	// RdOnly is whether the dialog is read only
	// RdOnly bool

	// a record of parent View names that have led up to this dialog,
	// which is displayed as extra contextual information in view dialog windows
	// VwPath string

	// Accepted means that the dialog was accepted -- else canceled
	Accepted bool

	// Buttons go here when added
	Buttons *Layout
}

// NewDialog returns a new [Dialog] in the context of the given widget,
// optionally with the given name.
func NewDialog(ctx Widget, name ...string) *Dialog {
	d := &Dialog{}
	nm := ""
	if len(name) > 0 {
		nm = name[0]
	} else {
		nm = ctx.Name() + "-dialog"
	}

	d.InitName(d, nm)
	d.EventMgr.Scene = &d.Scene
	d.BgColor.SetSolid(colors.Transparent)
	d.Lay = LayoutVert

	d.Stage = NewMainStage(DialogStage, &d.Scene, ctx)
	d.Modal(true)
	return d
}

// RecycleDialog looks for a dialog with the given data. If it
// finds it, it shows it and returns true. Otherwise, it returns false.
func RecycleDialog(data any) bool {
	rw, got := DialogRenderWins.FindData(data)
	if !got {
		return false
	}
	rw.Raise()
	return true
}

// Title adds the given title to the dialog
func (d *Dialog) Title(title string) *Dialog {
	d.Scene.Title = title
	NewLabel(d, "title").SetText(title).
		SetType(LabelHeadlineSmall).Style(func(s *styles.Style) {
		s.SetStretchMaxWidth()
		s.AlignH = styles.AlignCenter
		s.AlignV = styles.AlignTop
	})
	return d
}

// Prompt adds the given prompt to the dialog
func (d *Dialog) Prompt(prompt string) *Dialog {
	NewLabel(d, "prompt").SetText(prompt).
		SetType(LabelBodyMedium).Style(func(s *styles.Style) {
		s.Text.WhiteSpace = styles.WhiteSpaceNormal
		s.SetStretchMaxWidth()
		s.Width.Ch(30)
		s.Text.Align = styles.AlignLeft
		s.AlignV = styles.AlignTop
		s.Color = colors.Scheme.OnSurfaceVariant
	})
	return d
}

// ConfigButtons adds layout for holding buttons at bottom of dialog
// and saves as Buttons field, if not already done.
func (d *Dialog) ConfigButtons() *Layout {
	if d.Buttons != nil {
		return d.Buttons
	}
	bb := NewLayout(d, "buttons").
		SetLayout(LayoutHoriz)
	bb.Style(func(s *styles.Style) {
		bb.Spacing.Dp(8)
		s.SetStretchMaxWidth()
	})
	d.Buttons = bb
	NewStretch(bb)
	return bb
}

// Ok adds an OK button to the Buttons at bottom of dialog,
// connecting to Accept method the Ctrl+Enter keychord event.
// Also sends a Change event to the dialog for listeners there.
// If text is passed, that text is used for the text of the button
// instead of the standard "OK".
func (d *Dialog) Ok(text ...string) *Dialog {
	bb := d.ConfigButtons()
	txt := "OK"
	if len(text) > 0 {
		txt = text[0]
	}
	NewButton(bb, "ok").SetType(ButtonText).SetText(txt).OnClick(func(e events.Event) {
		e.SetHandled() // otherwise propagates to dead elements
		d.AcceptDialog()
	})
	d.OnKeyChord(func(e events.Event) {
		kf := keyfun.Of(e.KeyChord())
		if kf == keyfun.Accept {
			e.SetHandled()
			d.AcceptDialog()
		}
	})
	return d
}

// Cancel adds Cancel button to the Buttons at bottom of dialog,
// connecting to Cancel method and the Esc keychord event.
// Also sends a Change event to the dialog scene for listeners there.
// If text is passed, that text is used for the text of the button
// instead of the standard "Cancel".
func (d *Dialog) Cancel(text ...string) *Dialog {
	bb := d.ConfigButtons()
	txt := "Cancel"
	if len(text) > 0 {
		txt = text[0]
	}
	NewButton(bb, "cancel").SetType(ButtonText).SetText(txt).OnClick(func(e events.Event) {
		e.SetHandled() // otherwise propagates to dead elements
		d.CancelDialog()
	})
	d.OnKeyChord(func(e events.Event) {
		kf := keyfun.Of(e.KeyChord())
		if kf == keyfun.Abort {
			e.SetHandled()
			d.CancelDialog()
		}
	})
	return d
}

func (d *Dialog) Modal(modal bool) *Dialog {
	d.Stage.Modal = modal
	return d
}

func (d *Dialog) NewWindow(newWindow bool) *Dialog {
	d.Stage.NewWindow = newWindow
	return d
}

func (d *Dialog) FullWindow(fullWindow bool) *Dialog {
	d.Stage.FullWindow = fullWindow
	return d
}

// Run runs (shows) the dialog.
func (d *Dialog) Run() {
	d.DialogStyles()
	d.Stage.Run()
}

// AcceptDialog accepts the dialog, activated by the default Ok button
func (d *Dialog) AcceptDialog() {
	d.Accepted = true
	d.Send(events.Change)
	d.Close()
}

// CancelDialog cancels the dialog, activated by the default Cancel button
func (d *Dialog) CancelDialog() {
	d.Accepted = false
	d.Send(events.Change)
	d.Close()
}

// OnAccept adds an event listener for when the dialog is accepted
// (closed in a positive or neutral way)
func (d *Dialog) OnAccept(fun func(e events.Event)) *Dialog {
	d.OnChange(func(e events.Event) {
		if d.Accepted {
			fun(e)
		}
	})
	return d
}

// OnCancel adds an event listener for when the dialog is canceled
// (closed in a negative way)
func (d *Dialog) OnCancel(fun func(e events.Event)) *Dialog {
	d.OnChange(func(e events.Event) {
		if !d.Accepted {
			fun(e)
		}
	})
	return d
}

// Close closes the stage associated with this dialog
func (d *Dialog) Close() {
	mm := d.Stage.StageMgr
	if mm == nil {
		slog.Error("dialog has no MainMgr")
		return
	}
	if d.Stage.NewWindow {
		mm.RenderWin.CloseReq()
		return
	}
	mm.PopDeleteType(DialogStage)
}

// DefaultStyle sets default style functions for dialog Scene
func (d *Dialog) DialogStyles() {
	d.Style(func(s *styles.Style) {
		// s.Border.Radius = styles.BorderRadiusExtraLarge
		s.Color = colors.Scheme.OnSurface
		d.Spacing = StdDialogVSpaceUnits
		s.Padding.Left.Dp(8)
		if !d.Stage.NewWindow && !d.Stage.FullWindow {
			s.Padding.Set(units.Dp(24))
			s.Border.Radius = styles.BorderRadiusLarge
			s.BoxShadow = styles.BoxShadow3()
			// material likes SurfaceContainerHigh here, but that seems like too much; STYTODO: maybe figure out a better background color setup for dialogs?
			s.BackgroundColor.SetSolid(colors.Scheme.SurfaceContainerLow)
		}
	})
}

// NewItems adds to the dialog a prompt for creating new item(s) of the given type,
// showing registered gti types that embed given type.
func (dlg *Dialog) NewItems(typ *gti.Type) *Dialog {
	nrow := NewLayout(dlg, "n-row")
	nrow.Lay = LayoutHoriz

	NewLabel(nrow, "n-label").SetText("Number:  ")

	nsb := NewSpinner(nrow, "n-field")
	nsb.SetMin(1)
	nsb.Value = 1
	nsb.Step = 1

	tspc := NewSpace(dlg, "type-space")
	tspc.SetFixedHeight(units.Em(0.5))

	trow := NewLayout(dlg, "t-row")
	trow.Lay = LayoutHoriz

	NewLabel(trow, "t-label").SetText("Type:    ")

	typs := NewChooser(trow, "types")
	typs.ItemsFromTypes(gti.AllEmbeddersOf(typ), true, true, 50)

	dlg.Data = typ

	typs.OnChange(func(e events.Event) {
		dlg.Data = typs.CurVal
	})
	return dlg
}
