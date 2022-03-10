package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"

	"floodfill/floodfill"
	"floodfill/game"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

var (
	screenWidth  = 320
	screenHeight = 320
	tileSize     = 32

	input *game.Input
	tiles *game.TileCollection
	rooms *floodfill.RoomCollection

	fontFace font.Face
)

type Game struct {
}

func (g *Game) Update() error {
	input.Process()

	if input.IsLeftButtonDown {
		println("left")
	}
	if input.IsRightButtonDown {
		println("right")
	}

	cursorCol, cursorRow := convertCoordinatesToAddress(ebiten.CursorPosition())
	if input.IsLeftButtonPressed {
		if tiles.Has(cursorCol, cursorRow) {
			tile := tiles.Get(cursorCol, cursorRow)
			tile.IsBlock = true
		}
	}

	if input.IsRightButtonPressed {
		if tiles.Has(cursorCol, cursorRow) {
			tile := tiles.Get(cursorCol, cursorRow)
			tile.IsBlock = false
		}
	}

	if input.IsLeftButtonUp {
		rooms = floodfill.FloodFill(tiles)

		for _, room := range rooms.Rooms {
			game.SpriteById(room.RoomId, 32)
			fmt.Printf("room len: [%d],room id:[%d] hex:[%s]\n", room.Len(), room.RoomId, game.HexColor[room.RoomId])
		}
	}

	if input.IsRightButtonUp {
		rooms = floodfill.FloodFill(tiles)

		for _, room := range rooms.Rooms {
			game.SpriteById(room.RoomId, 32)
			fmt.Printf("room len: [%d],room id:[%d] hex:[%s]\n", room.Len(), room.RoomId, game.HexColor[room.RoomId])
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, tile := range tiles.All() {
		x, y := tile.Column*tileSize, tile.Row*tileSize
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))

		var sprite *ebiten.Image

		if tile.IsBlock {
			sprite = game.BlockSprite(tileSize)
		} else {
			sprite = game.SpriteById(tile.RoomId, 32)
		}

		screen.DrawImage(sprite, op)

		if !tile.IsBlock {
			text.Draw(
				screen,
				fmt.Sprintf("%d", tile.RoomId),
				fontFace,
				x+6,
				y+tileSize-6,
				game.HexToColor("ffffff"),
			)
		}
	}

	for line := 0; line <= screenWidth; line += tileSize {
		ebitenutil.DrawLine(
			screen,
			float64(line),
			float64(0),
			float64(line),
			float64(screenHeight),
			color.RGBA{255, 255, 255, 150},
		)
	}

	for line := 0; line <= screenHeight; line += tileSize {
		ebitenutil.DrawLine(
			screen,
			float64(0),
			float64(line),
			float64(screenWidth),
			float64(line),
			color.RGBA{255, 255, 255, 150},
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Flood Fill")
	ebiten.SetWindowResizable(false)
	ebiten.SetScreenClearedEveryFrame(false)

	columns, rows := screenWidth/tileSize, screenHeight/tileSize
	tiles = game.NewTileCollection(columns, rows)
	rooms = floodfill.FloodFill(tiles)
	input = game.NewInput()

	fontFace = loadFont("assets/NotoSans-Bold.ttf", 14)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func convertCoordinatesToAddress(x, y int) (column, row int) {
	column = x / tileSize
	row = y / tileSize

	return column, row
}

func loadFont(path string, size float64) font.Face {
	fontData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("fontFace not founded by path [%s], %v", path, err))
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		panic(fmt.Errorf("fontFace not parsed by path [%s], %v", path, err))
	}

	return truetype.NewFace(
		ttfFont, &truetype.Options{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingFull,
		},
	)
}
