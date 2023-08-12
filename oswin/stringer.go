// Code generated by "stringer -output stringer.go -type=Platforms,VirtualKeyboardTypes,EventType,ScreenOrientation,WindowFlags"; DO NOT EDIT.

package oswin

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MacOS-0]
	_ = x[LinuxX11-1]
	_ = x[Windows-2]
	_ = x[IOS-3]
	_ = x[Android-4]
	_ = x[PlatformsN-5]
}

const _Platforms_name = "MacOSLinuxX11WindowsIOSAndroidPlatformsN"

var _Platforms_index = [...]uint8{0, 5, 13, 20, 23, 30, 40}

func (i Platforms) String() string {
	if i < 0 || i >= Platforms(len(_Platforms_index)-1) {
		return "Platforms(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Platforms_name[_Platforms_index[i]:_Platforms_index[i+1]]
}

func (i *Platforms) FromString(s string) error {
	for j := 0; j < len(_Platforms_index)-1; j++ {
		if s == _Platforms_name[_Platforms_index[j]:_Platforms_index[j+1]] {
			*i = Platforms(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Platforms")
}

var _Platforms_descMap = map[Platforms]string{
	0: `MacOS is a mac desktop machine (aka Darwin)`,
	1: `LinuxX11 is a Linux OS machine running X11 window server`,
	2: `Windows is a Microsoft Windows machine`,
	3: `IOS is an Apple iOS or iPadOS mobile phone or iPad`,
	4: `Android is an Android mobile phone or tablet`,
	5: ``,
}

func (i Platforms) Desc() string {
	if str, ok := _Platforms_descMap[i]; ok {
		return str
	}
	return "Platforms(" + strconv.FormatInt(int64(i), 10) + ")"
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DefaultKeyboard-0]
	_ = x[SingleLineKeyboard-1]
	_ = x[NumberKeyboard-2]
	_ = x[VirtualKeyboardTypesN-3]
}

const _VirtualKeyboardTypes_name = "DefaultKeyboardSingleLineKeyboardNumberKeyboardVirtualKeyboardTypesN"

var _VirtualKeyboardTypes_index = [...]uint8{0, 15, 33, 47, 68}

func (i VirtualKeyboardTypes) String() string {
	if i < 0 || i >= VirtualKeyboardTypes(len(_VirtualKeyboardTypes_index)-1) {
		return "VirtualKeyboardTypes(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VirtualKeyboardTypes_name[_VirtualKeyboardTypes_index[i]:_VirtualKeyboardTypes_index[i+1]]
}

func (i *VirtualKeyboardTypes) FromString(s string) error {
	for j := 0; j < len(_VirtualKeyboardTypes_index)-1; j++ {
		if s == _VirtualKeyboardTypes_name[_VirtualKeyboardTypes_index[j]:_VirtualKeyboardTypes_index[j+1]] {
			*i = VirtualKeyboardTypes(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: VirtualKeyboardTypes")
}

var _VirtualKeyboardTypes_descMap = map[VirtualKeyboardTypes]string{
	0: `DefaultKeyboard is the keyboard with default input style and &#34;return&#34; return key`,
	1: `SingleLineKeyboard is the keyboard with default input style and &#34;Done&#34; return key`,
	2: `NumberKeyboard is the keyboard with number input style and &#34;Done&#34; return key`,
	3: ``,
}

func (i VirtualKeyboardTypes) Desc() string {
	if str, ok := _VirtualKeyboardTypes_descMap[i]; ok {
		return str
	}
	return "VirtualKeyboardTypes(" + strconv.FormatInt(int64(i), 10) + ")"
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MouseEvent-0]
	_ = x[MouseMoveEvent-1]
	_ = x[MouseDragEvent-2]
	_ = x[MouseScrollEvent-3]
	_ = x[MouseFocusEvent-4]
	_ = x[MouseHoverEvent-5]
	_ = x[KeyEvent-6]
	_ = x[KeyChordEvent-7]
	_ = x[TouchEvent-8]
	_ = x[MagnifyEvent-9]
	_ = x[RotateEvent-10]
	_ = x[WindowEvent-11]
	_ = x[WindowResizeEvent-12]
	_ = x[WindowPaintEvent-13]
	_ = x[WindowShowEvent-14]
	_ = x[WindowFocusEvent-15]
	_ = x[DNDEvent-16]
	_ = x[DNDMoveEvent-17]
	_ = x[DNDFocusEvent-18]
	_ = x[OSEvent-19]
	_ = x[OSOpenFilesEvent-20]
	_ = x[CustomEventType-21]
	_ = x[EventTypeN-22]
}

const _EventType_name = "MouseEventMouseMoveEventMouseDragEventMouseScrollEventMouseFocusEventMouseHoverEventKeyEventKeyChordEventTouchEventMagnifyEventRotateEventWindowEventWindowResizeEventWindowPaintEventWindowShowEventWindowFocusEventDNDEventDNDMoveEventDNDFocusEventOSEventOSOpenFilesEventCustomEventTypeEventTypeN"

var _EventType_index = [...]uint16{0, 10, 24, 38, 54, 69, 84, 92, 105, 115, 127, 138, 149, 166, 182, 197, 213, 221, 233, 246, 253, 269, 284, 294}

func (i EventType) String() string {
	if i < 0 || i >= EventType(len(_EventType_index)-1) {
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}

func (i *EventType) FromString(s string) error {
	for j := 0; j < len(_EventType_index)-1; j++ {
		if s == _EventType_name[_EventType_index[j]:_EventType_index[j+1]] {
			*i = EventType(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: EventType")
}

var _EventType_descMap = map[EventType]string{
	0:  `MouseEvent includes all mouse button actions, but not move or drag`,
	1:  `MouseMoveEvent is when the mouse is moving but no button is down`,
	2:  `MouseDragEvent is when the mouse is moving and there is a button down`,
	3:  `MouseScrollEvent is for mouse scroll wheel events`,
	4:  `MouseFocusEvent is for mouse focus (enter / exit of widget area) -- generated by gi.Window based on mouse move events`,
	5:  `MouseHoverEvent is for mouse hover -- generated by gi.Window based on mouse events`,
	6:  `KeyEvent for key pressed or released -- fine-grained data about each key as it happens`,
	7:  `KeyChordEvent is only generated when a non-modifier key is released, and it also contains a string representation of the full chord, suitable for translation into keyboard commands, emacs-style etc`,
	8:  `TouchEvent is a generic touch-based event`,
	9:  `MagnifyEvent is a touch-based magnify event (e.g., pinch)`,
	10: `RotateEvent is a touch-based rotate event`,
	11: `WindowEvent reports any changes in the window size, orientation, iconify, close, open, paint -- these are all &#34;internal&#34; events from OS to GUI system, and not sent to widgets`,
	12: `WindowResizeEvent is specifically for window resize events which need special treatment -- this is an internal event not sent to widgets`,
	13: `WindowPaintEvent is specifically for window paint events which need special treatment -- this is an internal event not sent to widgets, triggered right after window is opened for initial painting.`,
	14: `WindowShowEvent is a synthetic event sent to widget consumers, sent *only once* when window is shown for the very first time`,
	15: `WindowFocusEvent is a synthetic event sent to widget consumers, sent when window focus changes (action is Focus / DeFocus)`,
	16: `DNDEvent is for the Drag-n-Drop (DND) drop event`,
	17: `DNDMoveEvent is when the DND position has changed`,
	18: `DNDFocusEvent is for Enter / Exit events of the DND into / out of a given widget`,
	19: `OSEvent is an operating system generated event (app level typically)`,
	20: `OSOpenFilesEvent is an event telling app to open given files`,
	21: `CustomEventType is a user-defined event with a data any field`,
	22: `number of event types`,
}

func (i EventType) Desc() string {
	if str, ok := _EventType_descMap[i]; ok {
		return str
	}
	return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OrientationUnknown-0]
	_ = x[Portrait-1]
	_ = x[Landscape-2]
	_ = x[ScreenOrientationN-3]
}

const _ScreenOrientation_name = "OrientationUnknownPortraitLandscapeScreenOrientationN"

var _ScreenOrientation_index = [...]uint8{0, 18, 26, 35, 53}

func (i ScreenOrientation) String() string {
	if i < 0 || i >= ScreenOrientation(len(_ScreenOrientation_index)-1) {
		return "ScreenOrientation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ScreenOrientation_name[_ScreenOrientation_index[i]:_ScreenOrientation_index[i+1]]
}

func (i *ScreenOrientation) FromString(s string) error {
	for j := 0; j < len(_ScreenOrientation_index)-1; j++ {
		if s == _ScreenOrientation_name[_ScreenOrientation_index[j]:_ScreenOrientation_index[j+1]] {
			*i = ScreenOrientation(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: ScreenOrientation")
}

var _ScreenOrientation_descMap = map[ScreenOrientation]string{
	0: `OrientationUnknown means device orientation cannot be determined. Equivalent on Android to Configuration.ORIENTATION_UNKNOWN and on iOS to: UIDeviceOrientationUnknown UIDeviceOrientationFaceUp UIDeviceOrientationFaceDown`,
	1: `Portrait is a device oriented so it is tall and thin. Equivalent on Android to Configuration.ORIENTATION_PORTRAIT and on iOS to: UIDeviceOrientationPortrait UIDeviceOrientationPortraitUpsideDown`,
	2: `Landscape is a device oriented so it is short and wide. Equivalent on Android to Configuration.ORIENTATION_LANDSCAPE and on iOS to: UIDeviceOrientationLandscapeLeft UIDeviceOrientationLandscapeRight`,
	3: ``,
}

func (i ScreenOrientation) Desc() string {
	if str, ok := _ScreenOrientation_descMap[i]; ok {
		return str
	}
	return "ScreenOrientation(" + strconv.FormatInt(int64(i), 10) + ")"
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Dialog-0]
	_ = x[Modal-1]
	_ = x[Tool-2]
	_ = x[Fullscreen-3]
	_ = x[Minimized-4]
	_ = x[Focus-5]
	_ = x[WindowFlagsN-6]
}

const _WindowFlags_name = "DialogModalToolFullscreenMinimizedFocusWindowFlagsN"

var _WindowFlags_index = [...]uint8{0, 6, 11, 15, 25, 34, 39, 51}

func (i WindowFlags) String() string {
	if i < 0 || i >= WindowFlags(len(_WindowFlags_index)-1) {
		return "WindowFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _WindowFlags_name[_WindowFlags_index[i]:_WindowFlags_index[i+1]]
}

func (i *WindowFlags) FromString(s string) error {
	for j := 0; j < len(_WindowFlags_index)-1; j++ {
		if s == _WindowFlags_name[_WindowFlags_index[j]:_WindowFlags_index[j+1]] {
			*i = WindowFlags(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: WindowFlags")
}

var _WindowFlags_descMap = map[WindowFlags]string{
	0: `Dialog indicates that this is a temporary, pop-up window.`,
	1: `Modal indicates that this dialog window blocks events going to other windows until it is closed.`,
	2: `Tool indicates that this is a floating tool window that has minimized window decoration.`,
	3: `Fullscreen indicates a window that occupies the entire screen.`,
	4: `Minimized indicates a window reduced to an icon, or otherwise no longer visible or active. Otherwise, the window should be assumed to be visible.`,
	5: `Focus indicates that the window has the focus.`,
	6: ``,
}

func (i WindowFlags) Desc() string {
	if str, ok := _WindowFlags_descMap[i]; ok {
		return str
	}
	return "WindowFlags(" + strconv.FormatInt(int64(i), 10) + ")"
}
