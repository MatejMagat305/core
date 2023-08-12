// Code generated by "stringer -output stringer.go -type=Units"; DO NOT EDIT.

package units

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnitPx-0]
	_ = x[UnitDp-1]
	_ = x[UnitPct-2]
	_ = x[UnitRem-3]
	_ = x[UnitEm-4]
	_ = x[UnitEx-5]
	_ = x[UnitCh-6]
	_ = x[UnitVw-7]
	_ = x[UnitVh-8]
	_ = x[UnitVmin-9]
	_ = x[UnitVmax-10]
	_ = x[UnitCm-11]
	_ = x[UnitMm-12]
	_ = x[UnitQ-13]
	_ = x[UnitIn-14]
	_ = x[UnitPc-15]
	_ = x[UnitPt-16]
	_ = x[UnitDot-17]
	_ = x[UnitsN-18]
}

const _Units_name = "UnitPxUnitDpUnitPctUnitRemUnitEmUnitExUnitChUnitVwUnitVhUnitVminUnitVmaxUnitCmUnitMmUnitQUnitInUnitPcUnitPtUnitDotUnitsN"

var _Units_index = [...]uint8{0, 6, 12, 19, 26, 32, 38, 44, 50, 56, 64, 72, 78, 84, 89, 95, 101, 107, 114, 120}

func (i Units) String() string {
	if i < 0 || i >= Units(len(_Units_index)-1) {
		return "Units(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Units_name[_Units_index[i]:_Units_index[i+1]]
}

func (i *Units) FromString(s string) error {
	for j := 0; j < len(_Units_index)-1; j++ {
		if s == _Units_name[_Units_index[j]:_Units_index[j+1]] {
			*i = Units(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Units")
}

var _Units_descMap = map[Units]string{
	0:  `UnitPx = pixels -- 1px = 1/96th of 1in -- these are NOT raw display pixels`,
	1:  `UnitDp = density-independent pixels -- 1dp = 1/160th of 1in`,
	2:  `UnitPct = percentage of surrounding contextual element`,
	3:  `UnitRem = font size of the root element -- defaults to 12pt scaled by DPI factor`,
	4:  `UnitEm = font size of the element -- fallback to 12pt by default`,
	5:  `UnitEx = x-height of the element&#39;s font (size of &#39;x&#39; glyph) -- fallback to 0.5em by default`,
	6:  `UnitCh = width of the &#39;0&#39; glyph in the element&#39;s font -- fallback to 0.5em by default`,
	7:  `UnitVw = 1% of the viewport&#39;s width`,
	8:  `UnitVh = 1% of the viewport&#39;s height`,
	9:  `UnitVmin = 1% of the viewport&#39;s smaller dimension`,
	10: `UnitVmax = 1% of the viewport&#39;s larger dimension`,
	11: `UnitCm = centimeters -- 1cm = 96px/2.54`,
	12: `UnitMm = millimeters -- 1mm = 1/10th of cm`,
	13: `UnitQ = quarter-millimeters -- 1q = 1/40th of cm`,
	14: `UnitIn = inches -- 1in = 2.54cm = 96px`,
	15: `UnitPc = picas -- 1pc = 1/6th of 1in`,
	16: `UnitPt = points -- 1pt = 1/72th of 1in`,
	17: `UnitDot = actual real display pixels -- generally only use internally`,
	18: ``,
}

func (i Units) Desc() string {
	if str, ok := _Units_descMap[i]; ok {
		return str
	}
	return "Units(" + strconv.FormatInt(int64(i), 10) + ")"
}
