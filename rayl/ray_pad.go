package rayl

import (
	"fmt"

	"github.com/centretown/gpads/pad"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ pad.Pad = (*RayPad)(nil)

type RayPad struct{}

func (rp *RayPad) BeginPad() {
}

func (rp *RayPad) IsPadAvailable(Joystick int) bool {
	return rl.IsGamepadAvailable(int32(Joystick))
}

func (rp *RayPad) GetPadName(Joystick int) string {
	return rl.GetGamepadName(int32(Joystick))
}

func (rp *RayPad) IsPadButtonPressed(Joystick int, button int) bool {
	return rl.IsGamepadButtonPressed(int32(Joystick), int32(button))
}

func (rp *RayPad) IsPadButtonDown(Joystick int, button int) bool {
	return rl.IsGamepadButtonDown(int32(Joystick), int32(button))
}
func (rp *RayPad) IsPadButtonReleased(Joystick int, button int) bool {
	return rl.IsGamepadButtonReleased(int32(Joystick), int32(button))
}
func (rp *RayPad) IsPadButtonUp(Joystick int, button int) bool {
	return rl.IsGamepadButtonUp(int32(Joystick), int32(button))
}
func (rp *RayPad) GetPadButtonPressed() int {
	return int(rl.GetGamepadButtonPressed())
}
func (rp *RayPad) GetPadAxisCount(Joystick int) int {
	return int(rl.GetGamepadAxisCount(int32(Joystick)))
}
func (rp *RayPad) GetPadButtonCount(Joystick int) int {
	return 0
}
func (rp *RayPad) GetPadAxisMovement(Joystick int, axis int) float32 {
	return rl.GetGamepadAxisMovement(int32(Joystick), int32(axis))
}
func (rp *RayPad) GetPadAxisValue(Joystick int, axis int) int32 {
	return 0
}
func (rp *RayPad) SetPadMappings(mappings string) int {
	return int(rl.SetGamepadMappings(mappings))
}
func (rp *RayPad) DumpPad() {}
func (rp *RayPad) GetButtonName(Joystick int, button int) string {
	return fmt.Sprintf("button%d", button)
}
func (rp *RayPad) GetPadCount() int {
	return 0
}
