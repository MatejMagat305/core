// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"fmt"
	"log"
	"sync"

	"github.com/iancoleman/strcase"
	"goki.dev/colors"
	"goki.dev/cursors"
	"goki.dev/girl/abilities"
	"goki.dev/girl/states"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/goosi/events"
	"goki.dev/gti"
	"goki.dev/icons"
	"goki.dev/ki/v2"
)

// TabView switches among child widgets via tabs.  The selected widget gets
// the full allocated space avail after the tabs are accounted for.  The
// TabView is just a Vertical layout that manages two child widgets: a
// HorizFlow Layout for the tabs (which can flow across multiple rows as
// needed) and a Stacked Frame that actually contains all the children, and
// provides scrollbars as needed to any content within.  Typically should have
// max stretch and a set preferred size, so it expands.
//
//goki:embedder
type TabView struct {
	Layout

	// maximum number of characters to include in tab label -- elides labels that are longer than that
	MaxChars int `desc:"maximum number of characters to include in tab label -- elides labels that are longer than that"`

	// signal for tab widget -- see TabViewSignals for the types
	// TabViewSig ki.Signal `copy:"-" json:"-" xml:"-" desc:"signal for tab widget -- see TabViewSignals for the types"`

	// show a new tab button at right of list of tabs
	NewTabButton bool `desc:"show a new tab button at right of list of tabs"`

	// if true, tabs are not user-deleteable
	NoDeleteTabs bool `desc:"if true, tabs are not user-deleteable"`

	// type of widget to create in a new tab via new tab button -- Frame by default
	NewTabType *gti.Type `desc:"type of widget to create in a new tab via new tab button -- Frame by default"`

	// [view: -] mutex protecting updates to tabs -- tabs can be driven programmatically and via user input so need extra protection
	Mu sync.Mutex `copy:"-" json:"-" xml:"-" view:"-" desc:"mutex protecting updates to tabs -- tabs can be driven programmatically and via user input so need extra protection"`
}

func (tv *TabView) CopyFieldsFrom(frm any) {
	fr := frm.(*TabView)
	tv.Layout.CopyFieldsFrom(&fr.Layout)
	tv.MaxChars = fr.MaxChars
	tv.NewTabButton = fr.NewTabButton
	tv.NewTabType = fr.NewTabType
}

func (tv *TabView) OnInit() {
	tv.TabViewHandlers()
	tv.TabViewStyles()
}

func (tv *TabView) TabViewHandlers() {
	tv.LayoutHandlers()
}

func (tv *TabView) TabViewStyles() {
	tv.AddStyles(func(s *styles.Style) {
		// need border for separators (see RenderTabSeps)
		// TODO: maybe better solution for tab sep styles?
		s.Border.Style.Set(styles.BorderSolid)
		s.Border.Width.Set(units.Dp(1))
		s.Border.Color.Set(colors.Scheme.OutlineVariant)
		s.BackgroundColor.SetSolid(colors.Scheme.Background)
		s.Color = colors.Scheme.OnBackground
		s.MaxWidth.SetDp(-1)
		s.MaxHeight.SetDp(-1)
	})
}

func (tv *TabView) OnChildAdded(child ki.Ki) {
	w, _ := AsWidget(child)
	switch w.Name() {
	case "tabs":
		w.AddStyles(func(s *styles.Style) {
			s.SetStretchMaxWidth()
			s.Height.SetEm(1.8)
			s.Overflow = styles.OverflowHidden // no scrollbars!
			s.Margin.Set()
			s.Padding.Set()
			// tabs.Spacing.SetDp(4 * Prefs.DensityMul())
			s.BackgroundColor.SetSolid(colors.Scheme.SurfaceContainerLow)

			// s.Border.Style.Set(styles.BorderNone)
			// s.Border.Style.Bottom = styles.BorderSolid
			// s.Border.Width.Bottom.SetDp(1)
			// s.Border.Color.Bottom = colors.Scheme.OutlineVariant
		})
	case "frame":
		frame := w.(*Frame)
		frame.StackTopOnly = true // key for allowing each tab to have its own size
		w.AddStyles(func(s *styles.Style) {
			s.SetMinPrefWidth(units.Em(10))
			s.SetMinPrefHeight(units.Em(6))
			s.SetStretchMax()
		})
	}
}

// NTabs returns number of tabs
func (tv *TabView) NTabs() int {
	fr := tv.Frame()
	if fr == nil {
		return 0
	}
	return len(fr.Kids)
}

// CurTab returns currently-selected tab, and its index -- returns false none
func (tv *TabView) CurTab() (Widget, int, bool) {
	if tv.NTabs() == 0 {
		return nil, -1, false
	}
	tv.Mu.Lock()
	defer tv.Mu.Unlock()
	fr := tv.Frame()
	if fr.StackTop < 0 {
		return nil, -1, false
	}
	widg := fr.Child(fr.StackTop).(Widget)
	return widg, fr.StackTop, true
}

// TODO(kai): once subscenes are working, we should make tabs be subscenes

// NewTab adds a new tab with the given label and returns the resulting tab frame.
// It is the main end-user API for creating new tabs. If a name is also passed,
// the internal name (ID) of the tab will be set to that; otherwise, it will default
// to the kebab-case version of the label.
func (tv *TabView) NewTab(label string, name ...string) *Frame {
	fr := tv.Frame()
	idx := len(*fr.Children())
	frame := tv.InsertNewTab(label, idx, name...)
	return frame
}

// InsertNewTab inserts a new tab with the given label at the given index position
// within the list of tabs and returns the resulting tab frame. If a name is also
// passed, the internal name (ID) of the tab will be set to that; otherwise, it will default
// to the kebab-case version of the label.
func (tv *TabView) InsertNewTab(label string, idx int, name ...string) *Frame {
	updt := tv.UpdateStart()
	fr := tv.Frame()
	fr.SetChildAdded()
	nm := ""
	if len(name) > 0 {
		nm = name[0]
	} else {
		nm = strcase.ToKebab(label)
	}
	frame := fr.InsertNewChild(FrameType, idx, nm).(*Frame)
	tv.InsertTabOnlyAt(frame, label, idx, nm)
	tv.UpdateEndLayout(updt)
	return frame
}

// AddTab adds an already existing frame as a new tab with the given tab label
// and returns the index of that tab.
func (tv *TabView) AddTab(frame *Frame, label string) int {
	fr := tv.Frame()
	idx := len(*fr.Children())
	tv.InsertTab(frame, label, idx)
	return idx
}

// InsertTabOnlyAt inserts just the tab at given index, after the panel has
// already been added to the frame; assumed to be wrapped in update. Generally
// for internal use only. If a name is also passed, the internal name (ID) of the tab
// will be set to that; otherwise, it will default to the kebab-case version of the label.
func (tv *TabView) InsertTabOnlyAt(frame *Frame, label string, idx int, name ...string) {
	tb := tv.Tabs()
	tb.SetChildAdded()
	nm := ""
	if len(name) > 0 {
		nm = name[0]
	} else {
		nm = strcase.ToKebab(label)
	}
	tab := tb.InsertNewChild(TabButtonType, idx, nm).(*TabButton)
	tab.Data = idx
	tab.Tooltip = label
	tab.NoDelete = tv.NoDeleteTabs
	tab.SetText(label)
	tab.OnClick(func(e events.Event) {
		tv.SelectTabIndexAction(idx)
	})
	fr := tv.Frame()
	if len(fr.Kids) == 1 {
		fr.StackTop = 0
		tab.SetSelected(true)
	} else {
		frame.SetFlag(true, Invisible) // new tab is invisible until selected
	}
}

// InsertTab inserts a frame into given index position within list of tabs.
// If a name is also passed, the internal name (ID) of the tab will be set
// to that; otherwise, it will default to the kebab-case version of the label.
func (tv *TabView) InsertTab(frame *Frame, label string, idx int, name ...string) {
	tv.Mu.Lock()
	fr := tv.Frame()
	updt := tv.UpdateStart()
	fr.SetChildAdded()
	fr.InsertChild(frame, idx)
	tv.InsertTabOnlyAt(frame, label, idx, name...)
	tv.Mu.Unlock()
	tv.UpdateEndLayout(updt)
}

// TabAtIndex returns content widget and tab button at given index, false if
// index out of range (emits log message)
func (tv *TabView) TabAtIndex(idx int) (Widget, *TabButton, bool) {
	tv.Mu.Lock()
	defer tv.Mu.Unlock()

	fr := tv.Frame()
	tb := tv.Tabs()
	sz := len(*fr.Children())
	if idx < 0 || idx >= sz {
		log.Printf("giv.TabView: index %v out of range for number of tabs: %v\n", idx, sz)
		return nil, nil, false
	}
	tab := tb.Child(idx).(*TabButton)
	widg := fr.Child(idx).(Widget)
	return widg, tab, true
}

// SelectTabIndex selects tab at given index, returning it -- returns false if
// index is invalid
func (tv *TabView) SelectTabIndex(idx int) (Widget, bool) {
	widg, tab, ok := tv.TabAtIndex(idx)
	if !ok {
		return nil, false
	}
	fr := tv.Frame()
	if fr.StackTop == idx {
		return widg, true
	}
	tv.Mu.Lock()
	updt := tv.UpdateStart()
	tv.UnselectOtherTabs(idx)
	tab.SetSelected(true)
	fr.StackTop = idx
	tv.Mu.Unlock()
	tv.UpdateEndLayout(updt)
	return widg, true
}

// SelectTabIndexAction selects tab at given index and emits selected signal,
// with the index of the selected tab -- this is what is called when a tab is
// clicked
func (tv *TabView) SelectTabIndexAction(idx int) {
	_, ok := tv.SelectTabIndex(idx)
	if ok {
		// tv.TabViewSig.Emit(tv.This(), int64(TabSelected), idx)
	}
}

// TabByName returns tab with given name (nil if not found -- see TabByNameTry)
func (tv *TabView) TabByName(label string) Widget {
	t, _ := tv.TabByNameTry(label)
	return t
}

// TabByNameTry returns tab with given name, and an error if not found.
func (tv *TabView) TabByNameTry(label string) (Widget, error) {
	tv.Mu.Lock()
	defer tv.Mu.Unlock()

	tb := tv.Tabs()
	idx, ok := tb.Children().IndexByName(label, 0)
	if !ok {
		return nil, fmt.Errorf("gi.TabView: Tab named %v not found in %v", label, tv.Path())
	}
	fr := tv.Frame()
	widg := fr.Child(idx).(Widget)
	return widg, nil
}

// TabIndexByName returns tab index for given tab name, and an error if not found.
func (tv *TabView) TabIndexByName(label string) (int, error) {
	tv.Mu.Lock()
	defer tv.Mu.Unlock()

	tb := tv.Tabs()
	idx, ok := tb.Children().IndexByName(label, 0)
	if !ok {
		return -1, fmt.Errorf("gi.TabView: Tab named %v not found in %v", label, tv.Path())
	}
	return idx, nil
}

// TabName returns tab name at given index
func (tv *TabView) TabName(idx int) string {
	tv.Mu.Lock()
	defer tv.Mu.Unlock()

	tb := tv.Tabs()
	tbut, err := tb.ChildTry(idx)
	if err != nil {
		return ""
	}
	return tbut.Name()
}

// SelectTabByName selects tab by name, returning it.
func (tv *TabView) SelectTabByName(label string) Widget {
	idx, err := tv.TabIndexByName(label)
	if err == nil {
		tv.SelectTabIndex(idx)
		fr := tv.Frame()
		return fr.Child(idx).(Widget)
	}
	return nil
}

// SelectTabByNameTry selects tab by name, returning it.  Returns error if not found.
func (tv *TabView) SelectTabByNameTry(label string) (Widget, error) {
	idx, err := tv.TabIndexByName(label)
	if err == nil {
		tv.SelectTabIndex(idx)
		fr := tv.Frame()
		return fr.Child(idx).(Widget), nil
	}
	return nil, err
}

// RecycleTab returns a tab with given name, first by looking for an existing one,
// and if not found, making a new one. If sel, then select it. It returns the
// frame for the tab.
func (tv *TabView) RecycleTab(label string, sel bool) *Frame {
	frame, err := tv.TabByNameTry(label)
	if err == nil {
		if sel {
			tv.SelectTabByName(label)
		}
		return frame
	}
	frame = tv.NewTab(typ, label)
	if sel {
		tv.SelectTabByName(label)
	}
	return frame
}

// DeleteTabIndex deletes tab at given index, optionally calling destroy on
// tab contents -- returns widget if destroy == false, tab name, and bool success
func (tv *TabView) DeleteTabIndex(idx int, destroy bool) (Widget, string, bool) {
	widg, _, ok := tv.TabAtIndex(idx)
	if !ok {
		return nil, "", false
	}

	tnm := tv.TabName(idx)
	tv.Mu.Lock()
	fr := tv.Frame()
	sz := len(*fr.Children())
	tb := tv.Tabs()
	updt := tv.UpdateStart()
	nxtidx := -1
	if fr.StackTop == idx {
		if idx > 0 {
			nxtidx = idx - 1
		} else if idx < sz-1 {
			nxtidx = idx
		}
	}
	fr.DeleteChildAtIndex(idx, destroy)
	tb.DeleteChildAtIndex(idx, ki.DestroyKids) // always destroy -- we manage
	tv.RenumberTabs()
	tv.Mu.Unlock()
	if nxtidx >= 0 {
		tv.SelectTabIndex(nxtidx)
	}
	tv.UpdateEndLayout(updt)
	if destroy {
		return nil, tnm, true
	} else {
		return widg, tnm, true
	}
}

// DeleteTabIndexAction deletes tab at given index using destroy flag, and
// emits TabDeleted signal with name of deleted tab
// this is called by the delete button on the tab
func (tv *TabView) DeleteTabIndexAction(idx int) {
	_, tnm, ok := tv.DeleteTabIndex(idx, true)
	_ = tnm
	if ok {
		// todo: needed?
		// tv.TabViewSig.Emit(tv.This(), int64(TabDeleted), tnm)
	}
}

// ConfigNewTabButton configures the new tab + button at end of list of tabs
func (tv *TabView) ConfigNewTabButton(sc *Scene) bool {
	sz := tv.NTabs()
	tb := tv.Tabs()
	ntb := len(tb.Kids)
	if tv.NewTabButton {
		if ntb == sz+1 {
			return false
		}
		if tv.NewTabType == nil {
			tv.NewTabType = FrameType
		}
		tab := tb.InsertNewChild(ButtonType, ntb, "new-tab").(*Button)
		tab.Data = -1
		tab.SetIcon(icons.Add).SetType(ButtonAction)
		tab.OnClick(func(e events.Event) {
			tv.NewTabAction(tv.NewTabType, "New Tab")
			tv.SelectTabIndex(len(*tv.Frame().Children()) - 1)
		})
		return true
	} else {
		if ntb == sz {
			return false
		}
		tb.DeleteChildAtIndex(ntb-1, ki.DestroyKids) // always destroy -- we manage
		return true
	}
}

// TabViewSignals are signals that the TabView can send
type TabViewSignals int64

const (
	// TabSelected indicates tab was selected -- data is the tab index
	TabSelected TabViewSignals = iota

	// TabAdded indicates tab was added -- data is the tab index
	TabAdded

	// TabDeleted indicates tab was deleted -- data is the tab name
	TabDeleted

	TabViewSignalsN
)

// ConfigWidget initializes the tab widget children if it hasn't been done yet.
// only the 2 primary children (Frames) need to be configured.
// no re-config needed when adding / deleting tabs -- just new layout.
func (tv *TabView) ConfigWidget(sc *Scene) {
	if len(tv.Kids) != 0 {
		return
	}
	tv.Lay = LayoutVert

	frame := NewFrame(tv, "tabs")
	frame.Lay = LayoutHorizFlow

	frame = NewFrame(tv, "frame")
	frame.Lay = LayoutStacked

	tv.ConfigNewTabButton(sc)
}

// Tabs returns the layout containing the tabs -- the first element within us
func (tv *TabView) Tabs() *Frame {
	// TODO(kai): come up with a better structure for this?
	tv.ConfigWidget(tv.Sc)
	return tv.Child(0).(*Frame)
}

// Frame returns the stacked frame layout -- the second element
func (tv *TabView) Frame() *Frame {
	tv.ConfigWidget(tv.Sc)
	return tv.Child(1).(*Frame)
}

// UnselectOtherTabs turns off all the tabs except given one
func (tv *TabView) UnselectOtherTabs(idx int) {
	sz := tv.NTabs()
	tbs := tv.Tabs()
	for i := 0; i < sz; i++ {
		if i == idx {
			continue
		}
		tb := tbs.Child(i).(*TabButton)
		if tb.StateIs(states.Selected) {
			tb.SetSelected(false)
		}
	}
}

// RenumberTabs assigns proper index numbers to each tab
func (tv *TabView) RenumberTabs() {
	sz := tv.NTabs()
	tbs := tv.Tabs()
	for i := 0; i < sz; i++ {
		tb := tbs.Child(i).(*TabButton)
		tb.Data = i
	}
}

// RenderTabSeps renders the separators between tabs
func (tv *TabView) RenderTabSeps(sc *Scene) {
	rs, pc, st := tv.RenderLock(sc)
	defer tv.RenderUnlock(rs)

	// just like with standard separator, use top width like CSS
	// (see https://www.w3schools.com/howto/howto_css_dividers.asp)
	pc.StrokeStyle.Width = st.Border.Width.Top
	pc.StrokeStyle.SetColor(&st.Border.Color.Top)
	bw := st.Border.Width.Dots()

	tbs := tv.Tabs()
	sz := len(tbs.Kids)
	for i := 1; i < sz; i++ {
		tb := tbs.Child(i).(Widget)
		ni := tb.AsWidget()

		pos := ni.LayState.Alloc.Pos
		sz := ni.LayState.Alloc.Size.Sub(st.TotalMargin().Size())
		pc.DrawLine(rs, pos.X-bw.Pos().X, pos.Y, pos.X-bw.Pos().X, pos.Y+sz.Y)
	}
	pc.FillStrokeClear(rs)
}

func (tv *TabView) Render(sc *Scene) {
	if tv.PushBounds(sc) {
		tv.RenderScrolls(sc)
		tv.RenderChildren(sc)
		tv.RenderTabSeps(sc)
		tv.PopBounds(sc)
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// TabButton

// TabButton contains a larger select button and a small close button. Indicator
// icon is used for close icon.
type TabButton struct {
	Button

	// if true, this tab does not have the delete button avail
	NoDelete bool `desc:"if true, this tab does not have the delete button avail"`
}

func (tb *TabButton) OnInit() {
	tb.ButtonHandlers()
	tb.TabButtonStyles()
}

func (tb *TabButton) TabButtonStyles() {
	tb.AddStyles(func(s *styles.Style) {
		s.Cursor = cursors.Pointer
		s.MinWidth.SetCh(8)
		s.MaxWidth.SetDp(500)
		s.MinHeight.SetEm(1.6)

		// s.Border.Style.Right = styles.BorderSolid
		// s.Border.Width.Right.SetDp(1)

		s.BackgroundColor.SetSolid(colors.Scheme.SurfaceContainerLow)
		s.Color = colors.Scheme.OnSurface

		s.Border.Radius.Set()
		s.Text.Align = styles.AlignCenter
		s.Margin.Set()
		s.Padding.Set(units.Dp(8 * Prefs.DensityMul()))

		// s.Border.Style.Set(styles.BorderNone)
		// if tb.StateIs(states.Selected) {
		// 	s.Border.Style.Bottom = styles.BorderSolid
		// 	s.Border.Width.Bottom.SetDp(2)
		// 	s.Border.Color.Bottom = colors.Scheme.Primary
		// }
	})
}

func (tb *TabButton) OnChildAdded(child ki.Ki) {
	w, _ := AsWidget(child)
	switch w.Name() {
	case "Parts":
		w.AddStyles(func(s *styles.Style) {
			s.Overflow = styles.OverflowHidden // no scrollbars!
		})
	case "icon":
		w.AddStyles(func(s *styles.Style) {
			s.Width.SetEm(1)
			s.Height.SetEm(1)
			s.Margin.Set()
			s.Padding.Set()
		})
	case "label":
		label := w.(*Label)
		label.Type = LabelTitleSmall
		w.AddStyles(func(s *styles.Style) {
			s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
			s.Cursor = cursors.None
			s.Margin.Set()
			s.Padding.Set()
		})
	case "close-stretch":
		w.AddStyles(func(s *styles.Style) {
			s.Width.SetCh(1)
		})
	case "close":
		w.AddStyles(func(s *styles.Style) {
			s.Width.SetEx(0.5)
			s.Height.SetEx(0.5)
			s.Margin.Set()
			s.Padding.Set()
			s.AlignV = styles.AlignMiddle
			s.Border.Radius = styles.BorderRadiusFull
			s.BackgroundColor.SetSolid(colors.Transparent)
		})
	case "sc-stretch":
		w.AddStyles(func(s *styles.Style) {
			s.MinWidth.SetCh(2)
		})
	case "shortcut":
		w.AddStyles(func(s *styles.Style) {
			s.Margin.Set()
			s.Padding.Set()
		})
	}
}

func (tb *TabButton) TabView() *TabView {
	tv := tb.ParentByType(TabViewType, ki.Embeds)
	if tv == nil {
		return nil
	}
	return AsTabView(tv)
}

func (tb *TabButton) ConfigParts(sc *Scene) {
	if !tb.NoDelete {
		tb.ConfigPartsDeleteButton(sc)
		return
	}
	tb.Button.ConfigParts(sc) // regular
}

func (tb *TabButton) ConfigPartsDeleteButton(sc *Scene) {
	config := ki.Config{}
	icIdx, lbIdx := tb.ConfigPartsIconLabel(&config, tb.Icon, tb.Text)
	config.Add(StretchType, "close-stretch")
	clsIdx := len(config)
	config.Add(ButtonType, "close")
	mods, updt := tb.Parts.ConfigChildren(config)
	tb.ConfigPartsSetIconLabel(tb.Icon, tb.Text, icIdx, lbIdx)
	if mods {
		cls := tb.Parts.Child(clsIdx).(*Button)
		if tb.Indicator.IsNil() {
			tb.Indicator = icons.Close
		}

		icnm := tb.Indicator
		cls.SetIcon(icnm)
		cls.SetProp("no-focus", true)
		cls.OnClick(func(e events.Event) {
			tabIdx := tb.Data.(int)
			tvv := tb.TabView()
			if tvv != nil {
				if !Prefs.Params.OnlyCloseActiveTab || tb.StateIs(states.Selected) { // only process delete when already selected if OnlyCloseActiveTab is on
					tvv.DeleteTabIndexAction(tabIdx)
				} else {
					tvv.SelectTabIndexAction(tabIdx) // otherwise select
				}
			}
		})
		tb.UpdateEnd(updt)
	}
}

func (tb *TabButton) ConfigWidget(sc *Scene) {
	tb.ConfigParts(sc)
}
