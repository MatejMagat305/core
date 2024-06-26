// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package texteditor

import (
	"unicode"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/parse/lexer"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/texteditor/textbuf"
)

///////////////////////////////////////////////////////////////////////////////
//    Search / Find

// FindMatches finds the matches with given search string (literal, not regex)
// and case sensitivity, updates highlights for all.  returns false if none
// found
func (ed *Editor) FindMatches(find string, useCase, lexItems bool) ([]textbuf.Match, bool) {
	fsz := len(find)
	if fsz == 0 {
		ed.Highlights = nil
		return nil, false
	}
	_, matches := ed.Buffer.Search([]byte(find), !useCase, lexItems)
	if len(matches) == 0 {
		ed.Highlights = nil
		return matches, false
	}
	hi := make([]textbuf.Region, len(matches))
	for i, m := range matches {
		hi[i] = m.Reg
		if i > ViewMaxFindHighlights {
			break
		}
	}
	ed.Highlights = hi
	return matches, true
}

// MatchFromPos finds the match at or after the given text position -- returns 0, false if none
func (ed *Editor) MatchFromPos(matches []textbuf.Match, cpos lexer.Pos) (int, bool) {
	for i, m := range matches {
		reg := ed.Buffer.AdjustReg(m.Reg)
		if reg.Start == cpos || cpos.IsLess(reg.Start) {
			return i, true
		}
	}
	return 0, false
}

// ISearch holds all the interactive search data
type ISearch struct {

	// if true, in interactive search mode
	On bool `json:"-" xml:"-"`

	// current interactive search string
	Find string `json:"-" xml:"-"`

	// pay attention to case in isearch -- triggered by typing an upper-case letter
	UseCase bool `json:"-" xml:"-"`

	// current search matches
	Matches []textbuf.Match `json:"-" xml:"-"`

	// position within isearch matches
	Pos int `json:"-" xml:"-"`

	// position in search list from previous search
	PrevPos int `json:"-" xml:"-"`

	// starting position for search -- returns there after on cancel
	StartPos lexer.Pos `json:"-" xml:"-"`
}

// ViewMaxFindHighlights is the maximum number of regions to highlight on find
var ViewMaxFindHighlights = 1000

// PrevISearchString is the previous ISearch string
var PrevISearchString string

// ISearchMatches finds ISearch matches -- returns true if there are any
func (ed *Editor) ISearchMatches() bool {
	got := false
	ed.ISearch.Matches, got = ed.FindMatches(ed.ISearch.Find, ed.ISearch.UseCase, false)
	return got
}

// ISearchNextMatch finds next match after given cursor position, and highlights
// it, etc
func (ed *Editor) ISearchNextMatch(cpos lexer.Pos) bool {
	if len(ed.ISearch.Matches) == 0 {
		ed.ISearchSig()
		return false
	}
	ed.ISearch.Pos, _ = ed.MatchFromPos(ed.ISearch.Matches, cpos)
	ed.ISearchSelectMatch(ed.ISearch.Pos)
	return true
}

// ISearchSelectMatch selects match at given match index (e.g., ed.ISearch.Pos)
func (ed *Editor) ISearchSelectMatch(midx int) {
	nm := len(ed.ISearch.Matches)
	if midx >= nm {
		ed.ISearchSig()
		return
	}
	m := ed.ISearch.Matches[midx]
	reg := ed.Buffer.AdjustReg(m.Reg)
	pos := reg.Start
	ed.SelectRegion = reg
	ed.SetCursor(pos)
	ed.SavePosHistory(ed.CursorPos)
	ed.ScrollCursorToCenterIfHidden()
	ed.ISearchSig()
}

// ISearchSig sends the signal that ISearch is updated
func (ed *Editor) ISearchSig() {
	ed.Send(events.Input, nil)
}

// ISearchStart is an emacs-style interactive search mode -- this is called when
// the search command itself is entered
func (ed *Editor) ISearchStart() {
	if ed.ISearch.On {
		if ed.ISearch.Find != "" { // already searching -- find next
			sz := len(ed.ISearch.Matches)
			if sz > 0 {
				if ed.ISearch.Pos < sz-1 {
					ed.ISearch.Pos++
				} else {
					ed.ISearch.Pos = 0
				}
				ed.ISearchSelectMatch(ed.ISearch.Pos)
			}
		} else { // restore prev
			if PrevISearchString != "" {
				ed.ISearch.Find = PrevISearchString
				ed.ISearch.UseCase = lexer.HasUpperCase(ed.ISearch.Find)
				ed.ISearchMatches()
				ed.ISearchNextMatch(ed.CursorPos)
				ed.ISearch.StartPos = ed.CursorPos
			}
			// nothing..
		}
	} else {
		ed.ISearch.On = true
		ed.ISearch.Find = ""
		ed.ISearch.StartPos = ed.CursorPos
		ed.ISearch.UseCase = false
		ed.ISearch.Matches = nil
		ed.SelectReset()
		ed.ISearch.Pos = -1
		ed.ISearchSig()
	}
	ed.NeedsRender()
}

// ISearchKeyInput is an emacs-style interactive search mode -- this is called
// when keys are typed while in search mode
func (ed *Editor) ISearchKeyInput(kt events.Event) {
	kt.SetHandled()
	r := kt.KeyRune()
	// if ed.ISearch.Find == PrevISearchString { // undo starting point
	// 	ed.ISearch.Find = ""
	// }
	if unicode.IsUpper(r) { // todo: more complex
		ed.ISearch.UseCase = true
	}
	ed.ISearch.Find += string(r)
	ed.ISearchMatches()
	sz := len(ed.ISearch.Matches)
	if sz == 0 {
		ed.ISearch.Pos = -1
		ed.ISearchSig()
		return
	}
	ed.ISearchNextMatch(ed.CursorPos)
	ed.NeedsRender()
}

// ISearchBackspace gets rid of one item in search string
func (ed *Editor) ISearchBackspace() {
	if ed.ISearch.Find == PrevISearchString { // undo starting point
		ed.ISearch.Find = ""
		ed.ISearch.UseCase = false
		ed.ISearch.Matches = nil
		ed.SelectReset()
		ed.ISearch.Pos = -1
		ed.ISearchSig()
		return
	}
	if len(ed.ISearch.Find) <= 1 {
		ed.SelectReset()
		ed.ISearch.Find = ""
		ed.ISearch.UseCase = false
		return
	}
	ed.ISearch.Find = ed.ISearch.Find[:len(ed.ISearch.Find)-1]
	ed.ISearchMatches()
	sz := len(ed.ISearch.Matches)
	if sz == 0 {
		ed.ISearch.Pos = -1
		ed.ISearchSig()
		return
	}
	ed.ISearchNextMatch(ed.CursorPos)
	ed.NeedsRender()
}

// ISearchCancel cancels ISearch mode
func (ed *Editor) ISearchCancel() {
	if !ed.ISearch.On {
		return
	}
	if ed.ISearch.Find != "" {
		PrevISearchString = ed.ISearch.Find
	}
	ed.ISearch.PrevPos = ed.ISearch.Pos
	ed.ISearch.Find = ""
	ed.ISearch.UseCase = false
	ed.ISearch.On = false
	ed.ISearch.Pos = -1
	ed.ISearch.Matches = nil
	ed.Highlights = nil
	ed.SavePosHistory(ed.CursorPos)
	ed.SelectReset()
	ed.ISearchSig()
	ed.NeedsRender()
}

///////////////////////////////////////////////////////////////////////////////
//    Query-Replace

// QReplace holds all the query-replace data
type QReplace struct {

	// if true, in interactive search mode
	On bool `json:"-" xml:"-"`

	// current interactive search string
	Find string `json:"-" xml:"-"`

	// current interactive search string
	Replace string `json:"-" xml:"-"`

	// pay attention to case in isearch -- triggered by typing an upper-case letter
	UseCase bool `json:"-" xml:"-"`

	// search only as entire lexically tagged item boundaries -- key for replacing short local variables like i
	LexItems bool `json:"-" xml:"-"`

	// current search matches
	Matches []textbuf.Match `json:"-" xml:"-"`

	// position within isearch matches
	Pos int `json:"-" xml:"-"`

	// position in search list from previous search
	PrevPos int `json:"-" xml:"-"`

	// starting position for search -- returns there after on cancel
	StartPos lexer.Pos `json:"-" xml:"-"`
}

var (
	// PrevQReplaceFinds are the previous QReplace strings
	PrevQReplaceFinds []string

	// PrevQReplaceRepls are the previous QReplace strings
	PrevQReplaceRepls []string
)

// QReplaceSig sends the signal that QReplace is updated
func (ed *Editor) QReplaceSig() {
	ed.Send(events.Input, nil)
}

// QReplacePrompt is an emacs-style query-replace mode -- this starts the process, prompting
// user for items to search etc
func (ed *Editor) QReplacePrompt() {
	find := ""
	if ed.HasSelection() {
		find = string(ed.Selection().ToBytes())
	}
	d := core.NewBody().AddTitle("Query-Replace").
		AddText("Enter strings for find and replace, then select Query-Replace -- with dialog dismissed press <b>y</b> to replace current match, <b>n</b> to skip, <b>Enter</b> or <b>q</b> to quit, <b>!</b> to replace-all remaining")
	fc := core.NewChooser(d).SetEditable(true).SetDefaultNew(true)
	fc.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 0)
		s.Min.X.Ch(80)
	})
	fc.SetStrings(PrevQReplaceFinds...).SetCurrentIndex(0)
	if find != "" {
		fc.SetCurrentValue(find)
	}

	rc := core.NewChooser(d).SetEditable(true).SetDefaultNew(true)
	rc.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 0)
		s.Min.X.Ch(80)
	})
	rc.SetStrings(PrevQReplaceRepls...).SetCurrentIndex(0)

	lexitems := ed.QReplace.LexItems
	lxi := core.NewSwitch(d).SetText("Lexical Items").SetChecked(lexitems)
	lxi.SetTooltip("search matches entire lexically tagged items -- good for finding local variable names like 'i' and not matching everything")

	d.AddBottomBar(func(parent core.Widget) {
		d.AddCancel(parent)
		d.AddOK(parent).SetText("Query-Replace").OnClick(func(e events.Event) {
			var find, repl string
			if s, ok := fc.CurrentItem.Value.(string); ok {
				find = s
			}
			if s, ok := rc.CurrentItem.Value.(string); ok {
				repl = s
			}
			lexItems := lxi.IsChecked()
			ed.QReplaceStart(find, repl, lexItems)
		})
	})
	d.RunDialog(ed)
}

// QReplaceStart starts query-replace using given find, replace strings
func (ed *Editor) QReplaceStart(find, repl string, lexItems bool) {
	ed.QReplace.On = true
	ed.QReplace.Find = find
	ed.QReplace.Replace = repl
	ed.QReplace.LexItems = lexItems
	ed.QReplace.StartPos = ed.CursorPos
	ed.QReplace.UseCase = lexer.HasUpperCase(find)
	ed.QReplace.Matches = nil
	ed.QReplace.Pos = -1

	core.StringsInsertFirstUnique(&PrevQReplaceFinds, find, core.SystemSettings.SavedPathsMax)
	core.StringsInsertFirstUnique(&PrevQReplaceRepls, repl, core.SystemSettings.SavedPathsMax)

	ed.QReplaceMatches()
	ed.QReplace.Pos, _ = ed.MatchFromPos(ed.QReplace.Matches, ed.CursorPos)
	ed.QReplaceSelectMatch(ed.QReplace.Pos)
	ed.QReplaceSig()
}

// QReplaceMatches finds QReplace matches -- returns true if there are any
func (ed *Editor) QReplaceMatches() bool {
	got := false
	ed.QReplace.Matches, got = ed.FindMatches(ed.QReplace.Find, ed.QReplace.UseCase, ed.QReplace.LexItems)
	return got
}

// QReplaceNextMatch finds next match using, QReplace.Pos and highlights it, etc
func (ed *Editor) QReplaceNextMatch() bool {
	nm := len(ed.QReplace.Matches)
	if nm == 0 {
		return false
	}
	ed.QReplace.Pos++
	if ed.QReplace.Pos >= nm {
		return false
	}
	ed.QReplaceSelectMatch(ed.QReplace.Pos)
	return true
}

// QReplaceSelectMatch selects match at given match index (e.g., ed.QReplace.Pos)
func (ed *Editor) QReplaceSelectMatch(midx int) {
	nm := len(ed.QReplace.Matches)
	if midx >= nm {
		return
	}
	m := ed.QReplace.Matches[midx]
	reg := ed.Buffer.AdjustReg(m.Reg)
	pos := reg.Start
	ed.SelectRegion = reg
	ed.SetCursor(pos)
	ed.SavePosHistory(ed.CursorPos)
	ed.ScrollCursorToCenterIfHidden()
	ed.QReplaceSig()
}

// QReplaceReplace replaces at given match index (e.g., ed.QReplace.Pos)
func (ed *Editor) QReplaceReplace(midx int) {
	nm := len(ed.QReplace.Matches)
	if midx >= nm {
		return
	}
	m := ed.QReplace.Matches[midx]
	rep := ed.QReplace.Replace
	reg := ed.Buffer.AdjustReg(m.Reg)
	pos := reg.Start
	// last arg is matchCase, only if not using case to match and rep is also lower case
	matchCase := !ed.QReplace.UseCase && !lexer.HasUpperCase(rep)
	ed.Buffer.ReplaceText(reg.Start, reg.End, pos, rep, EditSignal, matchCase)
	ed.Highlights[midx] = textbuf.RegionNil
	ed.SetCursor(pos)
	ed.SavePosHistory(ed.CursorPos)
	ed.ScrollCursorToCenterIfHidden()
	ed.QReplaceSig()
}

// QReplaceReplaceAll replaces all remaining from index
func (ed *Editor) QReplaceReplaceAll(midx int) {
	nm := len(ed.QReplace.Matches)
	if midx >= nm {
		return
	}
	for mi := midx; mi < nm; mi++ {
		ed.QReplaceReplace(mi)
	}
}

// QReplaceKeyInput is an emacs-style interactive search mode -- this is called
// when keys are typed while in search mode
func (ed *Editor) QReplaceKeyInput(kt events.Event) {
	kt.SetHandled()
	switch {
	case kt.KeyRune() == 'y':
		ed.QReplaceReplace(ed.QReplace.Pos)
		if !ed.QReplaceNextMatch() {
			ed.QReplaceCancel()
		}
	case kt.KeyRune() == 'n':
		if !ed.QReplaceNextMatch() {
			ed.QReplaceCancel()
		}
	case kt.KeyRune() == 'q' || kt.KeyChord() == "ReturnEnter":
		ed.QReplaceCancel()
	case kt.KeyRune() == '!':
		ed.QReplaceReplaceAll(ed.QReplace.Pos)
		ed.QReplaceCancel()
	}
	ed.NeedsRender()
}

// QReplaceCancel cancels QReplace mode
func (ed *Editor) QReplaceCancel() {
	if !ed.QReplace.On {
		return
	}
	ed.QReplace.On = false
	ed.QReplace.Pos = -1
	ed.QReplace.Matches = nil
	ed.Highlights = nil
	ed.SavePosHistory(ed.CursorPos)
	ed.SelectReset()
	ed.QReplaceSig()
	ed.NeedsRender()
}

// EscPressed emitted for [keymap.Abort] or [keymap.CancelSelect];
// effect depends on state.
func (ed *Editor) EscPressed() {
	switch {
	case ed.ISearch.On:
		ed.ISearchCancel()
		ed.SetCursorShow(ed.ISearch.StartPos)
	case ed.QReplace.On:
		ed.QReplaceCancel()
		ed.SetCursorShow(ed.ISearch.StartPos)
	case ed.HasSelection():
		ed.SelectReset()
	default:
		ed.Highlights = nil
	}
	ed.NeedsRender()
}
