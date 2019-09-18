// Code generated by "stringer -type=SliderSignals"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SliderValueChanged-0]
	_ = x[SliderPressed-1]
	_ = x[SliderReleased-2]
	_ = x[SliderMoved-3]
	_ = x[SliderSignalsN-4]
}

const _SliderSignals_name = "SliderValueChangedSliderPressedSliderReleasedSliderMovedSliderSignalsN"

var _SliderSignals_index = [...]uint8{0, 18, 31, 45, 56, 70}

func (i SliderSignals) String() string {
	if i < 0 || i >= SliderSignals(len(_SliderSignals_index)-1) {
		return "SliderSignals(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SliderSignals_name[_SliderSignals_index[i]:_SliderSignals_index[i+1]]
}

func (i *SliderSignals) FromString(s string) error {
	for j := 0; j < len(_SliderSignals_index)-1; j++ {
		if s == _SliderSignals_name[_SliderSignals_index[j]:_SliderSignals_index[j+1]] {
			*i = SliderSignals(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: SliderSignals")
}
