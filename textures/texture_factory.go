package textures

import "main/types"

const (
	PieceSizeInTexture = 80
	OffsetInTexture    = 80
	TileSize           = 80
	TexturePath        = "textures/chess_pieces.png"
)

type TextureFactory struct {
	*SpriteSheet
}

func NewTextureFactory() *TextureFactory {
	return &TextureFactory{
		SpriteSheet: NewSpriteSheet(TexturePath),
	}
}

func (factory *TextureFactory) CreateTexture(pType types.PieceType, pColor types.PieceColor) *Texture {
	var colInSpriteSheet int
	switch pType {
	case types.King:
		colInSpriteSheet = 0
	case types.Queen:
		colInSpriteSheet = 1
	case types.Bishop:
		colInSpriteSheet = 2
	case types.Knight:
		colInSpriteSheet = 3
	case types.Rook:
		colInSpriteSheet = 4
	case types.Pawn:
		colInSpriteSheet = 5
	}

	rowInSpriteSheet := 0
	if pColor == types.Black {
		rowInSpriteSheet = 1
	}

	return factory.SpriteSheet.CreateTexture(
		float32(OffsetInTexture*colInSpriteSheet),
		float32(OffsetInTexture*rowInSpriteSheet),
		PieceSizeInTexture,
		PieceSizeInTexture,
		TileSize,
		TileSize,
	)
}
