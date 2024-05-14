// Code generated by "stringer -type=Command"; DO NOT EDIT.

package notes

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NONE-0]
	_ = x[OPTIONS-1]
	_ = x[MORE-2]
	_ = x[NEXT-3]
	_ = x[PREVIOUS-4]
	_ = x[SET-5]
	_ = x[INCREMENT-6]
	_ = x[DECREMENT-7]
	_ = x[INCREMENT_MORE-8]
	_ = x[DECREMENT_MORE-9]
	_ = x[PAUSE_PLAY-10]
	_ = x[SHARE-11]
	_ = x[ACTION-12]
	_ = x[BACK-13]
	_ = x[CANCEL-14]
	_ = x[OUT-15]
	_ = x[COMMAND_COUNT-16]
}

const _Command_name = "NONEOPTIONSMORENEXTPREVIOUSSETINCREMENTDECREMENTINCREMENT_MOREDECREMENT_MOREPAUSE_PLAYSHAREACTIONBACKCANCELOUTCOMMAND_COUNT"

var _Command_index = [...]uint8{0, 4, 11, 15, 19, 27, 30, 39, 48, 62, 76, 86, 91, 97, 101, 107, 110, 123}

func (i Command) String() string {
	if i < 0 || i >= Command(len(_Command_index)-1) {
		return "Command(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Command_name[_Command_index[i]:_Command_index[i+1]]
}