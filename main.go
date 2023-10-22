package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"main/animation"
)

const (
	WindowWidth  = 1000
	WindowHeight = 1000
	WindowTitle  = "Chess"
)

const (
	GroundSize = 400
)

func main() {
	rl.InitWindow(WindowWidth, WindowHeight, WindowTitle)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetWindowState(rl.FlagWindowResizable)

	walking := animation.ExampleDinoWalkingAnimation()
	running := animation.ExampleDinoRunningAnimation()
	sprinting := animation.ExampleDinoSprintingAnimation()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.SkyBlue)
			rl.DrawRectangle(0, WindowHeight-GroundSize, WindowWidth, GroundSize, rl.Brown)
			walking.Update()
			walking.Render()

			running.Update()
			running.Render()

			sprinting.Update()
			sprinting.Render()
		}
		rl.EndDrawing()
	}
}
