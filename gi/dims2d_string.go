// Code generated by "stringer -type=Dims2D"; DO NOT EDIT.

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
	_ = x[X-0]
	_ = x[Y-1]
	_ = x[Dims2DN-2]
}

const _Dims2D_name = "XYDims2DN"

var _Dims2D_index = [...]uint8{0, 1, 2, 9}

func (i Dims2D) String() string {
	if i < 0 || i >= Dims2D(len(_Dims2D_index)-1) {
		return "Dims2D(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Dims2D_name[_Dims2D_index[i]:_Dims2D_index[i+1]]
}

func (i *Dims2D) FromString(s string) error {
	for j := 0; j < len(_Dims2D_index)-1; j++ {
		if s == _Dims2D_name[_Dims2D_index[j]:_Dims2D_index[j+1]] {
			*i = Dims2D(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Dims2D")
}
