// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filetree

import (
	"strings"

	"cogentcore.org/core/base/vcs"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/keymap"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/states"
)

// VCSLabelFunc gets the appropriate label for removing from version control
func VCSLabelFunc(fn *Node, label string) string {
	repo, _ := fn.Repo()
	if repo != nil {
		label = strings.Replace(label, "VCS", string(repo.Vcs()), 1)
	}
	return label
}

func (fn *Node) VCSContextMenu(m *core.Scene) {
	core.NewFuncButton(m).SetFunc(fn.AddToVCSSel).SetText(VCSLabelFunc(fn, "Add to VCS")).SetIcon(icons.Add).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || fn.Info.VCS != vcs.Untracked, states.Disabled)
		})
	core.NewFuncButton(m).SetFunc(fn.DeleteFromVCSSel).SetText(VCSLabelFunc(fn, "Delete from VCS")).SetIcon(icons.Delete).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || fn.Info.VCS == vcs.Untracked, states.Disabled)
		})
	core.NewFuncButton(m).SetFunc(fn.CommitToVCSSel).SetText(VCSLabelFunc(fn, "Commit to VCS")).SetIcon(icons.Star).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || fn.Info.VCS == vcs.Untracked, states.Disabled)
		})
	core.NewFuncButton(m).SetFunc(fn.RevertVCSSel).SetText(VCSLabelFunc(fn, "Revert from VCS")).SetIcon(icons.Undo).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || fn.Info.VCS == vcs.Untracked, states.Disabled)
		})
	core.NewSeparator(m)

	core.NewFuncButton(m).SetFunc(fn.DiffVCSSel).SetText(VCSLabelFunc(fn, "Diff VCS")).SetIcon(icons.Add).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || fn.Info.VCS == vcs.Untracked, states.Disabled)
		})
	core.NewFuncButton(m).SetFunc(fn.LogVCSSel).SetText(VCSLabelFunc(fn, "Log VCS")).SetIcon(icons.List).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || fn.Info.VCS == vcs.Untracked, states.Disabled)
		})
	core.NewFuncButton(m).SetFunc(fn.BlameVCSSel).SetText(VCSLabelFunc(fn, "Blame VC S")).SetIcon(icons.CreditScore).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || fn.Info.VCS == vcs.Untracked, states.Disabled)
		})
}

func (fn *Node) ContextMenu(m *core.Scene) {
	core.NewButton(m).SetText("Info").SetIcon(icons.Info).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection(), states.Disabled)
		}).OnClick(func(e events.Event) {
		fn.This.(Filer).ShowFileInfo()
	})

	core.NewButton(m).SetText("Open").SetIcon(icons.Open).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection(), states.Disabled)
		}).OnClick(func(e events.Event) {
		fn.This.(Filer).OpenFilesDefault()
	})
	core.NewSeparator(m)

	core.NewButton(m).SetText("Duplicate").SetIcon(icons.Copy).
		SetKey(keymap.Duplicate).Styler(func(s *styles.Style) {
		s.SetState(!fn.HasSelection(), states.Disabled)
	}).OnClick(func(e events.Event) {
		fn.This.(Filer).DuplicateFiles()
	})
	core.NewButton(m).SetText("Delete").SetIcon(icons.Delete).
		SetKey(keymap.Delete).Styler(func(s *styles.Style) {
		s.SetState(!fn.HasSelection(), states.Disabled)
	}).OnClick(func(e events.Event) {
		fn.This.(Filer).DeleteFiles()
	})
	core.NewButton(m).SetText("Rename").SetIcon(icons.NewLabel).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection(), states.Disabled)
		}).OnClick(func(e events.Event) {
		fn.This.(Filer).RenameFiles()
	})
	core.NewSeparator(m)

	core.NewFuncButton(m).SetFunc(fn.OpenAll).SetText("Open all").SetIcon(icons.KeyboardArrowDown).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || !fn.IsDir(), states.Disabled)
		})
	core.NewFuncButton(m).SetFunc(fn.CloseAll).SetIcon(icons.KeyboardArrowRight).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || !fn.IsDir(), states.Disabled)
		})
	core.NewFuncButton(m).SetFunc(fn.SortBys).SetText("Sort by").SetIcon(icons.Sort).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection() || !fn.IsDir(), states.Disabled)
		})
	core.NewSeparator(m)

	core.NewFuncButton(m).SetFunc(fn.NewFiles).SetText("New file").SetIcon(icons.OpenInNew).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection(), states.Disabled)
		})
	core.NewFuncButton(m).SetFunc(fn.NewFolders).SetText("New folder").SetIcon(icons.CreateNewFolder).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection(), states.Disabled)
		})
	core.NewSeparator(m)

	fn.VCSContextMenu(m)
	core.NewSeparator(m)

	core.NewFuncButton(m).SetFunc(fn.RemoveFromExterns).SetIcon(icons.Delete).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection(), states.Disabled)
		})

	core.NewSeparator(m)
	core.NewButton(m).SetText("Copy").SetIcon(icons.ContentCopy).SetKey(keymap.Copy).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection(), states.Disabled)
		}).
		OnClick(func(e events.Event) {
			fn.Copy(true)
		})
	core.NewButton(m).SetText("Cut").SetIcon(icons.ContentCut).SetKey(keymap.Cut).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection(), states.Disabled)
		}).
		OnClick(func(e events.Event) {
			fn.Cut()
		})
	pbt := core.NewButton(m).SetText("Paste").SetIcon(icons.ContentPaste).SetKey(keymap.Paste).
		Styler(func(s *styles.Style) {
			s.SetState(!fn.HasSelection(), states.Disabled)
		}).
		OnClick(func(e events.Event) {
			fn.Paste()
		})
	cb := fn.Events().Clipboard()
	if cb != nil {
		pbt.SetState(cb.IsEmpty(), states.Disabled)
	}
}
