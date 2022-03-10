package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 320
)

type GameObject interface {
	update() error
	draw(screen *ebiten.Image)
}

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
