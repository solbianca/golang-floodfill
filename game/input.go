package game

import "github.com/hajimehoshi/ebiten/v2"

type Input struct {
	IsLeftButtonPressed bool
	IsLeftButtonDown    bool
	IsLeftButtonUp      bool

	IsRightButtonPressed bool
	IsRightButtonDown    bool
	IsRightButtonUp      bool
}

func NewInput() *Input {
	return &Input{
		IsLeftButtonPressed:  false,
		IsLeftButtonDown:     false,
		IsLeftButtonUp:       false,
		IsRightButtonPressed: false,
		IsRightButtonDown:    false,
		IsRightButtonUp:      false,
	}
}

func (i *Input) Process() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && i.IsLeftButtonPressed == false {
		i.IsLeftButtonPressed = true
		i.IsLeftButtonDown = true

		return
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && i.IsLeftButtonPressed == true {
		i.IsLeftButtonPressed = false
		i.IsLeftButtonDown = false
		i.IsLeftButtonUp = true

		return
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && i.IsRightButtonPressed == false {
		i.IsRightButtonPressed = true
		i.IsRightButtonDown = true

		return
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && i.IsRightButtonPressed == true {
		i.IsRightButtonPressed = false
		i.IsRightButtonDown = false
		i.IsRightButtonUp = true

		return
	}

	if i.IsLeftButtonPressed == true && i.IsLeftButtonDown == true {
		i.IsLeftButtonDown = false
		return
	}

	if i.IsRightButtonPressed == true && i.IsRightButtonDown == true {
		i.IsRightButtonDown = false
		return
	}

	if i.IsLeftButtonPressed == false && i.IsLeftButtonUp == true {
		i.IsLeftButtonUp = false
	}

	if i.IsRightButtonPressed == false && i.IsRightButtonUp == true {
		i.IsRightButtonUp = false
	}
}
