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
	case types.KING:
		colInSpriteSheet = 0
	case types.QUEEN:
		colInSpriteSheet = 1
	case types.BISHOP:
		colInSpriteSheet = 2
	case types.KNIGHT:
		colInSpriteSheet = 3
	case types.ROOK:
		colInSpriteSheet = 4
	case types.PAWN:
		colInSpriteSheet = 5
	}

	rowInSpriteSheet := 0
	if pColor == types.BLACK {
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
