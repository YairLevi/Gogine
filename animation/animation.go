package animation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
Animation by default goes to the right.
*/
type Animation struct {
	spriteSheet       rl.Texture2D
	frameRect         rl.Rectangle
	sourceX           float32
	sourceY           float32
	sourceWidth       float32
	sourceHeight      float32
	imagesInAnimation int
	frameSpeed        float32

	frameCounter int
	currentFrame int
}

func NewAnimation(
	spriteSheet rl.Texture2D,
	imagesInAnimation int,
	sourceWidth float32,
	sourceHeight float32,
	frameSpeed float32,
	frameRect rl.Rectangle,
) *Animation {
	return &Animation{
		spriteSheet:       spriteSheet,
		frameRect:         frameRect,
		frameSpeed:        frameSpeed,
		sourceWidth:       sourceWidth,
		sourceHeight:      sourceHeight,
		imagesInAnimation: imagesInAnimation,
		frameCounter:      0,
		currentFrame:      0,
	}
}

func (a *Animation) Update() {
	a.frameCounter++
	if float32(a.frameCounter) >= float32(rl.GetFPS())/a.frameSpeed {
		a.currentFrame++
		a.frameCounter = 0

		if a.currentFrame >= a.imagesInAnimation {
			a.currentFrame = 0
		}
	}
}

func (a *Animation) Render() {
	srcX := float32(a.currentFrame)*a.sourceWidth + a.sourceX
	srcY := a.sourceY
	rl.DrawTexturePro(
		a.spriteSheet,
		rl.NewRectangle(
			srcX,
			srcY,
			a.sourceWidth,
			a.sourceHeight,
		),
		a.frameRect,
		rl.Vector2{},
		0,
		rl.White,
	)
}

func ExampleDinoWalkingAnimation() *Animation {
	return NewAnimation(
		rl.LoadTexture("textures/dino.png"),
		4,
		24,
		24,
		6,
		rl.NewRectangle(100, 100, 100, 100),
	)
}

func ExampleDinoRunningAnimation() *Animation {
	anim := NewAnimation(
		rl.LoadTexture("textures/dino.png"),
		6,
		24,
		24,
		8,
		rl.NewRectangle(300, 100, 100, 100),
	)
	anim.sourceX = 24 * 4
	return anim
}

func ExampleDinoSprintingAnimation() *Animation {
	anim := NewAnimation(
		rl.LoadTexture("textures/dino.png"),
		6,
		24,
		24,
		10,
		rl.NewRectangle(500, 100, 100, 100),
	)
	anim.sourceX = 24 * 18
	return anim
}
