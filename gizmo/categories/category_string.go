// Code generated by "stringer -type=Category"; DO NOT EDIT.

package categories

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Game-1]
	_ = x[Ellipse-2]
	_ = x[Texture-3]
	_ = x[Tracker-4]
	_ = x[LifeMover-5]
	_ = x[LifeGrid-6]
	_ = x[Player-7]
	_ = x[CellsOrg-8]
	_ = x[COUNT-9]
}

const _Category_name = "UnknownGameEllipseTextureTrackerLifeMoverLifeGridPlayerCellsOrgCOUNT"

var _Category_index = [...]uint8{0, 7, 11, 18, 25, 32, 41, 49, 55, 63, 68}

func (i Category) String() string {
	if i < 0 || i >= Category(len(_Category_index)-1) {
		return "Category(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Category_name[_Category_index[i]:_Category_index[i+1]]
}
