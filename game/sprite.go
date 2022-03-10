package game

import "github.com/hajimehoshi/ebiten/v2"

var (
	sprites        map[int]*ebiten.Image
	HexColor       map[int]string
	blockSpriteImg *ebiten.Image
)

func init() {
	sprites = map[int]*ebiten.Image{}
	HexColor = map[int]string{}
}

func SpriteById(id, tileSize int) *ebiten.Image {
	sprite, ok := sprites[id]

	if ok {
		return sprite
	}

	sprite = ebiten.NewImage(tileSize, tileSize)
	color, hex := RandomColor()
	HexColor[id] = hex
	sprite.Fill(color)

	sprites[id] = sprite

	return sprite
}

func BlockSprite(tileSize int) *ebiten.Image {
	if blockSpriteImg != nil {
		return blockSpriteImg
	}

	blockSpriteImg = ebiten.NewImage(tileSize, tileSize)
	color, _ := RandomColor()
	blockSpriteImg.Fill(color)

	return blockSpriteImg
}
