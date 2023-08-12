// Code generated by "stringer -output stringer.go -type=Cases"; DO NOT EDIT.

package textbuf

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LowerCase-0]
	_ = x[UpperCase-1]
	_ = x[CamelCase-2]
	_ = x[LowerCamelCase-3]
	_ = x[SnakeCase-4]
	_ = x[UpperSnakeCase-5]
	_ = x[KebabCase-6]
	_ = x[CasesN-7]
}

const _Cases_name = "LowerCaseUpperCaseCamelCaseLowerCamelCaseSnakeCaseUpperSnakeCaseKebabCaseCasesN"

var _Cases_index = [...]uint8{0, 9, 18, 27, 41, 50, 64, 73, 79}

func (i Cases) String() string {
	if i < 0 || i >= Cases(len(_Cases_index)-1) {
		return "Cases(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Cases_name[_Cases_index[i]:_Cases_index[i+1]]
}

func (i *Cases) FromString(s string) error {
	for j := 0; j < len(_Cases_index)-1; j++ {
		if s == _Cases_name[_Cases_index[j]:_Cases_index[j+1]] {
			*i = Cases(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Cases")
}

var _Cases_descMap = map[Cases]string{
	0: ``,
	1: ``,
	2: `CamelCase is init-caps`,
	3: `LowerCamelCase has first letter lower-case`,
	4: `SnakeCase is snake_case -- lower with underbars`,
	5: `UpperSnakeCase is SNAKE_CASE -- upper with underbars`,
	6: `KebabCase is kebab-case -- lower with -&#39;s`,
	7: `CasesN is the number of textview states`,
}

func (i Cases) Desc() string {
	if str, ok := _Cases_descMap[i]; ok {
		return str
	}
	return "Cases(" + strconv.FormatInt(int64(i), 10) + ")"
}
