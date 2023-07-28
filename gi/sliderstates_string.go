// Code generated by "stringer -type=SliderStates"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SliderActive-0]
	_ = x[SliderInactive-1]
	_ = x[SliderHover-2]
	_ = x[SliderFocus-3]
	_ = x[SliderDown-4]
	_ = x[SliderSelected-5]
	_ = x[SliderValue-6]
	_ = x[SliderBox-7]
	_ = x[SliderStatesN-8]
}

const _SliderStates_name = "SliderActiveSliderInactiveSliderHoverSliderFocusSliderDownSliderSelectedSliderValueSliderBoxSliderStatesN"

var _SliderStates_index = [...]uint8{0, 12, 26, 37, 48, 58, 72, 83, 92, 105}

func (i SliderStates) String() string {
	if i < 0 || i >= SliderStates(len(_SliderStates_index)-1) {
		return "SliderStates(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SliderStates_name[_SliderStates_index[i]:_SliderStates_index[i+1]]
}

func (i *SliderStates) FromString(s string) error {
	for j := 0; j < len(_SliderStates_index)-1; j++ {
		if s == _SliderStates_name[_SliderStates_index[j]:_SliderStates_index[j+1]] {
			*i = SliderStates(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: SliderStates")
}

var _SliderStates_descMap = map[SliderStates]string{
	0: `normal state -- there but not being interacted with`,
	1: `inactive -- not responsive`,
	2: `mouse is hovering over the slider`,
	3: `slider is the focus -- will respond to keyboard input`,
	4: `slider is currently being pressed down`,
	5: `slider has been selected`,
	6: `use background-color here to fill in selected value of slider`,
	7: `these styles define the overall box around slider -- typically no border and a white background -- needs a background to allow local re-rendering`,
	8: `total number of slider states`,
}

func (i SliderStates) Desc() string {
	if str, ok := _SliderStates_descMap[i]; ok {
		return str
	}
	return "SliderStates(" + strconv.FormatInt(int64(i), 10) + ")"
}
