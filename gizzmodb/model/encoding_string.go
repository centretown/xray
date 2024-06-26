// Code generated by "stringer -type=Encoding"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RAW-0]
	_ = x[JSON-1]
	_ = x[YAML-2]
	_ = x[XML-3]
	_ = x[CSV-4]
}

const _Encoding_name = "RAWJSONYAMLXMLCSV"

var _Encoding_index = [...]uint8{0, 3, 7, 11, 14, 17}

func (i Encoding) String() string {
	if i < 0 || i >= Encoding(len(_Encoding_index)-1) {
		return "Encoding(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Encoding_name[_Encoding_index[i]:_Encoding_index[i+1]]
}
