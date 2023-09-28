package main

import (
	"fmt"
	"io/ioutil"
	"os"

	// "./opcodes/opcodes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"main.go/opcodes"
)

type Game struct{
	cpu opcodes.Cpu
}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 320
}

func main() {
    ebiten.SetWindowSize(640, 320)
    ebiten.SetWindowTitle("Chip8 Emulator")
    ebiten.RunGame(&Game{})
    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }

}