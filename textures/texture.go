package textures

import rl "github.com/gen2brain/raylib-go/raylib"

type SpriteSheet struct {
	Sheet rl.Texture2D
}

func NewSpriteSheet(file string) *SpriteSheet {
	return &SpriteSheet{
		Sheet: rl.LoadTexture(file),
	}
}

func (ss *SpriteSheet) CreateTexture(srcX, srcY, texWidth, texHeight, targetWidth, targetHeight float32) *Texture {
	return &Texture{
		Src:   rl.NewRectangle(srcX, srcY, texWidth, texHeight),
		Dest:  rl.NewRectangle(-1, -1, targetWidth, targetHeight),
		Sheet: ss.Sheet,
	}
}

type Texture struct {
	Sheet rl.Texture2D
	Src   rl.Rectangle
	Dest  rl.Rectangle
}

func (t *Texture) Draw(posX, posY float32) {
	t.Dest.X = posX
	t.Dest.Y = posY
	rl.DrawTexturePro(t.Sheet, t.Src, t.Dest, rl.Vector2{}, 0, rl.White)
}
