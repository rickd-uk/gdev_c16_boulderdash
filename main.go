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
	screenWidth        = 320
	screenHeight       = 200
	tileSize           = 16
	simulationInterval = 12
)

type Tile uint8

const (
	Empty Tile = iota
	Dirt
	Wall
	Player
	Boulder
	Diamond
)

type Game struct {
	cave [][]Tile

	playerX int
	playerY int

	simulationTicks  int
	diamondCollected int
	score            int
}

func newGame() *Game {
	cave := [][]Tile{
		{Wall, Wall, Wall, Wall, Wall, Wall, Wall},
		{Wall, Dirt, Dirt, Dirt, Dirt, Dirt, Wall},
		{Wall, Player, Diamond, Empty, Diamond, Dirt, Wall},
		{Wall, Wall, Wall, Wall, Wall, Wall, Wall},
	}
	return &Game{
		cave:    cave,
		playerX: 1,
		playerY: 2,
	}
}

func (g *Game) inBounds(x, y int) bool {
	return y >= 0 &&
		y < len(g.cave) &&
		x >= 0 &&
		x < len(g.cave[y])
}

func (g *Game) tileAt(x, y int) Tile {
	if !g.inBounds(x, y) {
		return Wall
	}
	return g.cave[y][x]
}

func (g *Game) setTile(x, y int, tile Tile) {
	if !g.inBounds(x, y) {
		return
	}
	g.cave[y][x] = tile
}

func (g *Game) movePlayer(dx, dy int) {
	newX := g.playerX + dx
	newY := g.playerY + dy

	target := g.tileAt(newX, newY)

	if target == Boulder {
		// allow horizontal pushing
		if dy != 0 {
			return
		}
		beyondX := newX + dx
		beyondY := newY

		// if space to push boulder into is not empty you cannot push
		if g.tileAt(beyondX, beyondY) != Empty {
			return
		}
		g.setTile(beyondX, beyondY, Boulder)
		g.setTile(newX, newY, Empty)

		// If the space for player to move is NOT Empty / Dirt, can't move
		// We already handled the Boulder
	} else if target != Empty && target != Dirt && target != Diamond {
		return
	}

	if target == Diamond {
		g.diamondCollected++
		g.score += 10
	}

	g.setTile(g.playerX, g.playerY, Empty)
	g.setTile(newX, newY, Player)

	g.playerX = newX
	g.playerY = newY

	g.printCave()
	fmt.Println()
}

func (g *Game) updateGravity() {
	changed := false

	for y := len(g.cave) - 1; y >= 0; y-- {
		for x := range g.cave[y] {
			// It must be a Boulder and space below must be empty to
			// for Boulder to drop down into it
			if g.tileAt(x, y) != Boulder {
				continue
			}
			below := g.tileAt(x, y+1)

			if below == Empty {
				g.setTile(x, y, Empty)
				g.setTile(x, y+1, Boulder)
				changed = true
				continue
			}
			if below != Boulder {
				continue
			}
			if g.tileAt(x-1, y) == Empty &&
				g.tileAt(x-1, y+1) == Empty {
				g.setTile(x, y, Empty)
				g.setTile(x-1, y+1, Boulder)
				changed = true
			} else if g.tileAt(x+1, y) == Empty &&
				g.tileAt(x+1, y+1) == Empty {
				g.setTile(x, y, Empty)
				g.setTile(x+1, y+1, Boulder)
				changed = true
			}
		}
	}
	if changed {
		g.printCave()
		fmt.Println()
	}
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

	g.simulationTicks++

	if g.simulationTicks >= simulationInterval {
		g.simulationTicks = 0
		g.updateGravity()
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
			case Boulder:
				vector.FillRect(
					screen,
					px,
					py,
					tileSize,
					tileSize,
					colorRGB(160, 160, 160),
					false,
				)
			case Diamond:
				vector.FillRect(
					screen,
					px,
					py,
					tileSize,
					tileSize,
					colorRGB(80, 220, 240),
					false,
				)
			}
		}
	}
}

func (g *Game) printCave() {
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
			case Boulder:
				fmt.Print("O")
			case Diamond:
				fmt.Print("D")
			}
		}
		fmt.Println()
	}
	fmt.Printf("Diamonds:   %d  Score:  %d\n", g.diamondCollected, g.score)
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

	game := newGame()
	game.printCave()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
