// Code generated by "stringer -type=BorderDrawStyle"; DO NOT EDIT.

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
	_ = x[BorderSolid-0]
	_ = x[BorderDotted-1]
	_ = x[BorderDashed-2]
	_ = x[BorderDouble-3]
	_ = x[BorderGroove-4]
	_ = x[BorderRidge-5]
	_ = x[BorderInset-6]
	_ = x[BorderOutset-7]
	_ = x[BorderNone-8]
	_ = x[BorderHidden-9]
	_ = x[BorderN-10]
}

const _BorderDrawStyle_name = "BorderSolidBorderDottedBorderDashedBorderDoubleBorderGrooveBorderRidgeBorderInsetBorderOutsetBorderNoneBorderHiddenBorderN"

var _BorderDrawStyle_index = [...]uint8{0, 11, 23, 35, 47, 59, 70, 81, 93, 103, 115, 122}

func (i BorderDrawStyle) String() string {
	if i < 0 || i >= BorderDrawStyle(len(_BorderDrawStyle_index)-1) {
		return "BorderDrawStyle(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BorderDrawStyle_name[_BorderDrawStyle_index[i]:_BorderDrawStyle_index[i+1]]
}

func (i *BorderDrawStyle) FromString(s string) error {
	for j := 0; j < len(_BorderDrawStyle_index)-1; j++ {
		if s == _BorderDrawStyle_name[_BorderDrawStyle_index[j]:_BorderDrawStyle_index[j+1]] {
			*i = BorderDrawStyle(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: BorderDrawStyle")
}
