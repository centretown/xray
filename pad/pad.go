package pad

type Pad interface {
	BeginPad()
	IsPadAvailable(Joystick int) bool
	GetPadName(Joystick int) string
	IsPadButtonPressed(Joystick int, button int) bool
	IsPadButtonDown(Joystick int, button int) bool
	IsPadButtonReleased(Joystick int, button int) bool
	IsPadButtonUp(Joystick int, button int) bool
	GetPadButtonPressed() int
	GetPadAxisCount(Joystick int) int
	GetPadButtonCount(Joystick int) int
	GetPadAxisMovement(Joystick int, axis int) float32
	GetPadAxisValue(Joystick int, axis int) int32
	SetPadMappings(mappings string) int
	DumpPad()
	GetButtonName(Joystick int, button int) string
}
