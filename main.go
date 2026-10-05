package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	Player
)

type Game struct {
	cave [][]Tile

	playerX int
	playerY int
}

func NewGame() *Game {
	cave := [][]Tile{
		{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},
		{Wall, Dirt, Dirt, Dirt, Dirt, Dirt, Dirt, Dirt, Dirt, Wall},
		{Wall, Dirt, Empty, Empty, Empty, Empty, Empty, Dirt, Dirt, Wall},
		{Wall, Dirt, Dirt, Dirt, Player, Dirt, Dirt, Dirt, Dirt, Wall},
		{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},
	}
	return &Game{
		cave:    cave,
		playerX: 4,
		playerY: 3,
	}
}

func (g *Game) movePlayer(dx, dy int) {
	newX := g.playerX + dx
	newY := g.playerY + dy

	target := g.cave[newY][newX]

	if target != Empty && target != Dirt {
		return
	}
	g.cave[g.playerY][g.playerX] = Empty
	g.cave[newY][newX] = Player

	g.playerX = newX
	g.playerY = newY

	g.PrintCave()
	fmt.Println()
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.movePlayer(-1, 0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.movePlayer(1, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.movePlayer(0, -1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.movePlayer(0, 1)
	}

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
			case Player:
				vector.FillRect(
					screen,
					px,
					py,
					tileSize,
					tileSize,
					colorRGB(240, 220, 80),
					false,
				)
			}
		}
	}
}

func (g *Game) PrintCave() {
	for _, row := range g.cave {
		for _, tile := range row {
			switch tile {
			case Empty:
				fmt.Print(" ")
			case Dirt:
				fmt.Print("#")
			case Wall:
				fmt.Print("*")
			case Player:
				fmt.Print("|")
			}
		}
		fmt.Println()
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
	game.PrintCave()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
