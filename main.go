package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 320
	screenHeight = 200
	tileSize     = 16
)

type Tile uint8

const (
	Empty Tile = iota
	Dirt
	Wall
)

type Game struct {
	cave [][]Tile
}

func NewGame() *Game {
	cave := [][]Tile{
		{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},
		{Wall, Dirt, Dirt, Dirt, Dirt, Dirt, Dirt, Dirt, Dirt, Wall},
		{Wall, Dirt, Empty, Empty, Empty, Empty, Empty, Dirt, Dirt, Wall},
		{Wall, Dirt, Dirt, Dirt, Dirt, Dirt, Dirt, Dirt, Dirt, Wall},
		{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},
	}
	return &Game{
		cave: cave,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y, row := range g.cave {
		for x, tile := range row {
			px := float32(x * tileSize)
			py := float32(y * tileSize)

			switch tile {
			case Dirt:
				vector.FillRect(
					screen,
					px,
					py,
					tileSize,
					tileSize,
					colorRGB(120, 72, 32),
					false,
				)
			case Wall:
				vector.FillRect(
					screen,
					px,
					py,
					tileSize,
					tileSize,
					colorRGB(80, 80, 160),
					false,
				)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func colorRGB(r, g, b uint8) color.RGBA {
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

func main() {
	ebiten.SetWindowSize(960, 600)
	ebiten.SetWindowTitle("Boulder Dash")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
