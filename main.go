package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	game2 "main/game"
	"main/textures"
)

const (
	WindowWidth  = 1000
	WindowHeight = 1000
	WindowTitle  = "Chess"
)

var (
	LightSquareColor = rl.NewColor(227, 193, 111, 255)
	DarkSquareColor  = rl.NewColor(184, 139, 74, 255)
)

func main() {
	rl.InitWindow(WindowWidth, WindowHeight, WindowTitle)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	game := game2.NewGame(0, 0, textures.TileSize, LightSquareColor, DarkSquareColor)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)

		game.Run()

		rl.EndDrawing()
	}
}
