// Code generated by "stringer -type=MessageItem"; DO NOT EDIT.

package message

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MajorUsage-0]
	_ = x[MinorUsage-1]
	_ = x[KeyUsage-2]
	_ = x[InstallUsage-3]
	_ = x[QuickUsage-4]
	_ = x[Monitor-5]
	_ = x[View-6]
	_ = x[Capture-7]
	_ = x[Duration-8]
	_ = x[Frames-9]
	_ = x[Capturing-10]
	_ = x[Mhz-11]
	_ = x[Counter-12]
	_ = x[FPS-13]
	_ = x[LastTextItem-14]
}

const _MessageItem_name = "MajorUsageMinorUsageKeyUsageInstallUsageQuickUsageMonitorViewCaptureDurationFramesCapturingMhzCounterFPSLastTextItem"

var _MessageItem_index = [...]uint8{0, 10, 20, 28, 40, 50, 57, 61, 68, 76, 82, 91, 94, 101, 104, 116}

func (i MessageItem) String() string {
	if i < 0 || i >= MessageItem(len(_MessageItem_index)-1) {
		return "MessageItem(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MessageItem_name[_MessageItem_index[i]:_MessageItem_index[i+1]]
}
