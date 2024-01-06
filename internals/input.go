
package internals

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Input manages the input state including gamepads and keyboards.
type Input struct {
	gamepadIDs                 []ebiten.GamepadID
	virtualGamepadButtonStates map[virtualGamepadButton]int
	gamepadConfig              gamepadConfig
}

// GamepadIDButtonPressed returns a gamepad ID where at least one button is pressed.
// If no button is pressed, GamepadIDButtonPressed returns -1.
func (i *Input) GamepadIDButtonPressed() ebiten.GamepadID {
	i.gamepadIDs = ebiten.AppendGamepadIDs(i.gamepadIDs[:0])
	for _, id := range i.gamepadIDs {
		for b := ebiten.GamepadButton(0); b <= ebiten.GamepadButtonMax; b++ {
			if ebiten.IsGamepadButtonPressed(id, b) {
				return id
			}
		}
	}
	return -1
}

func (i *Input) stateForVirtualGamepadButton(b virtualGamepadButton) int {
	if i.virtualGamepadButtonStates == nil {
		return 0
	}
	return i.virtualGamepadButtonStates[b]
}

func (i *Input) Update() {
	if !i.gamepadConfig.IsGamepadIDInitialized() {
		return
	}

	if i.virtualGamepadButtonStates == nil {
		i.virtualGamepadButtonStates = map[virtualGamepadButton]int{}
	}
	for _, b := range virtualGamepadButtons {
		if !i.gamepadConfig.IsButtonPressed(b) {
			i.virtualGamepadButtonStates[b] = 0
			continue
		}
		i.virtualGamepadButtonStates[b]++
	}
}

func (i *Input) IsRotateRightJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyX) {
		return true
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonButtonB) == 1
}

func (i *Input) IsRotateLeftJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		return true
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonButtonA) == 1
}

func (i *Input) StateForLeft() int {
	if v := inpututil.KeyPressDuration(ebiten.KeyArrowLeft); 0 < v {
		return v
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonLeft)
}

func (i *Input) StateForRight() int {
	if v := inpututil.KeyPressDuration(ebiten.KeyArrowRight); 0 < v {
		return v
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonRight)
}

func (i *Input) StateForDown() int {
	if v := inpututil.KeyPressDuration(ebiten.KeyArrowDown); 0 < v {
		return v
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonDown)
}
