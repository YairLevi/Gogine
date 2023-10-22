package game

import (
	"main/textures"
	"main/types"
)

type PieceFactory struct {
	tf *textures.TextureFactory
}

func NewPieceFactory() *PieceFactory {
	return &PieceFactory{
		tf: textures.NewTextureFactory(),
	}
}

func (pf *PieceFactory) Create(pType types.PieceType, pColor types.PieceColor) PieceInterface {
	switch pType {
	case types.KING:
		return &king{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
		}
	case types.QUEEN:
		return &Queen{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
		}
	case types.ROOK:
		return &rook{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
			hasMoved: false,
		}
	case types.KNIGHT:
		return &knight{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
		}
	case types.BISHOP:
		return &Bishop{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
		}
	case types.PAWN:
		return &pawn{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
		}
	default:
		return nil
	}
}
