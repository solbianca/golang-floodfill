package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"
)

// rgbColor RBG HexColor Type
type rgbColor struct {
	Red   int
	Green int
	Blue  int
}

// RandomColor returns a random color in HEX format
func RandomColor() (color.Color, string) {
	rgbColor := getRandomColorInRgb()
	hex := getHex(rgbColor.Red) + getHex(rgbColor.Green) + getHex(rgbColor.Blue)
	return HexToColor(hex), hex
}

// getRandomColorInRgb Returns a random rgbColor
func getRandomColorInRgb() rgbColor {
	rand.Seed(time.Now().UnixNano())
	Red := rand.Intn(255)
	Green := rand.Intn(255)
	blue := rand.Intn(255)
	c := rgbColor{Red, Green, blue}
	return c
}

// GetHex Converts a decimal number to hex representations
func getHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

func HexToColor(h string) color.Color {
	u, err := strconv.ParseUint(h, 16, 0)
	if err != nil {
		panic(err)
	}

	return color.RGBA{
		R: uint8(u & 0xff0000 >> 16),
		G: uint8(u & 0xff00 >> 8),
		B: uint8(u & 0xff),
		A: 255,
	}
}
