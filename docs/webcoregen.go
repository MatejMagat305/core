// Code generated by "core generate -webcore content"; DO NOT EDIT.

package main

import (
	"errors"
	"fmt"
	"maps"
	"strings"

	"cogentcore.org/core/events"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/webcore"
)

func init() {
	maps.Copy(webcore.Examples, WebcoreExamples)
}

// WebcoreExamples are the compiled webcore examples for this app.
var WebcoreExamples = map[string]func(parent gi.Widget){
	"getting-started/hello-world-0": func(parent gi.Widget) {
		b := parent
		gi.NewButton(b).SetText("Hello, World!")
	},
	"basics/widgets-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").SetIcon(icons.Add)
	},
	"basics/widgets-1": func(parent gi.Widget) {
		sw := gi.NewSwitch(parent).SetText("Switch me!")
		// Later...
		gi.MessageSnackbar(parent, sw.Text)
	},
	"basics/events-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Button clicked")
		})
	},
	"basics/events-1": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprint("Button clicked at ", e.Pos()))
			e.SetHandled() // this event will not be handled by other event handlers now
		})
	},
	"basics/icons-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Send").SetIcon(icons.Send).OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Message sent")
		})
	},
	"widgets/buttons-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Download").SetIcon(icons.Download)
	},
	"widgets/buttons-1": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Send").SetIcon(icons.Send).OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Message sent")
		})
	},
	"widgets/buttons-2": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Share").SetIcon(icons.Share).SetMenu(func(m *gi.Scene) {
			gi.NewButton(m).SetText("Copy link")
			gi.NewButton(m).SetText("Send message")
		})
	},
	"widgets/buttons-3": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonFilled).SetText("Filled")
	},
	"widgets/buttons-4": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonTonal).SetText("Tonal")
	},
	"widgets/buttons-5": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonElevated).SetText("Elevated")
	},
	"widgets/buttons-6": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonOutlined).SetText("Outlined")
	},
	"widgets/buttons-7": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonText).SetText("Text")
	},
	"widgets/buttons-8": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonAction).SetText("Action")
	},
	"widgets/choosers-0": func(parent gi.Widget) {
		gi.NewChooser(parent).SetStrings("macOS", "Windows", "Linux")
	},
	"widgets/choosers-1": func(parent gi.Widget) {
		gi.NewChooser(parent).SetItems(
			gi.ChooserItem{Value: "Computer", Icon: icons.Computer, Tooltip: "Use a computer"},
			gi.ChooserItem{Value: "Phone", Icon: icons.Smartphone, Tooltip: "Use a phone"},
		)
	},
	"widgets/choosers-2": func(parent gi.Widget) {
		gi.NewChooser(parent).SetPlaceholder("Choose a platform").SetStrings("macOS", "Windows", "Linux")
	},
	"widgets/choosers-3": func(parent gi.Widget) {
		gi.NewChooser(parent).SetStrings("Apple", "Orange", "Strawberry").SetCurrentValue("Orange")
	},
	"widgets/choosers-4": func(parent gi.Widget) {
		gi.NewChooser(parent).SetType(gi.ChooserOutlined).SetStrings("Apple", "Orange", "Strawberry")
	},
	"widgets/choosers-5": func(parent gi.Widget) {
		gi.NewChooser(parent).SetIcon(icons.Sort).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-6": func(parent gi.Widget) {
		gi.NewChooser(parent).SetEditable(true).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-7": func(parent gi.Widget) {
		gi.NewChooser(parent).SetAllowNew(true).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-8": func(parent gi.Widget) {
		gi.NewChooser(parent).SetEditable(true).SetAllowNew(true).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-9": func(parent gi.Widget) {
		ch := gi.NewChooser(parent).SetStrings("Newest", "Oldest", "Popular")
		ch.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("Sorting by %v", ch.CurrentItem.Value))
		})
	},
	"widgets/spinners-0": func(parent gi.Widget) {
		gi.NewSpinner(parent)
	},
	"widgets/spinners-1": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetValue(12.7)
	},
	"widgets/spinners-2": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetMin(-0.5).SetMax(2.7)
	},
	"widgets/spinners-3": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetStep(6)
	},
	"widgets/spinners-4": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetStep(4).SetEnforceStep(true)
	},
	"widgets/spinners-5": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetType(gi.TextFieldOutlined)
	},
	"widgets/spinners-6": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetFormat("%X").SetStep(1).SetValue(44)
	},
	"widgets/spinners-7": func(parent gi.Widget) {
		sp := gi.NewSpinner(parent)
		sp.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("Value changed to %g", sp.Value))
		})
	},
	"widgets/text-fields-0": func(parent gi.Widget) {
		gi.NewTextField(parent)
	},
	"widgets/text-fields-1": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("Name:")
		gi.NewTextField(parent).SetPlaceholder("Jane Doe")
	},
	"widgets/text-fields-2": func(parent gi.Widget) {
		gi.NewTextField(parent).SetText("Hello, world!")
	},
	"widgets/text-fields-3": func(parent gi.Widget) {
		gi.NewTextField(parent).SetText("This is a long sentence that demonstrates how text field content can overflow onto multiple lines")
	},
	"widgets/text-fields-4": func(parent gi.Widget) {
		gi.NewTextField(parent).SetType(gi.TextFieldOutlined)
	},
	"widgets/text-fields-5": func(parent gi.Widget) {
		gi.NewTextField(parent).SetTypePassword()
	},
	"widgets/text-fields-6": func(parent gi.Widget) {
		gi.NewTextField(parent).AddClearButton()
	},
	"widgets/text-fields-7": func(parent gi.Widget) {
		gi.NewTextField(parent).SetLeadingIcon(icons.Euro).SetTrailingIcon(icons.OpenInNew, func(e events.Event) {
			gi.MessageSnackbar(parent, "Opening shopping cart")
		})
	},
	"widgets/text-fields-8": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.SetValidator(func() error {
			if !strings.Contains(tf.Text(), "Go") {
				return errors.New("Must contain Go")
			}
			return nil
		})
	},
	"widgets/text-fields-9": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, "OnChange: "+tf.Text())
		})
	},
	"widgets/text-fields-10": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.OnInput(func(e events.Event) {
			gi.MessageSnackbar(parent, "OnInput: "+tf.Text())
		})
	},
}
