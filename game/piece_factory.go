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
	case types.King:
		return &king{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
		}
	case types.Queen:
		return &queen{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
		}
	case types.Rook:
		return &rook{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
			hasMoved: false,
		}
	case types.Knight:
		return &knight{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
		}
	case types.Bishop:
		return &bishop{
			Piece: &Piece{
				Type:    pType,
				Color:   pColor,
				Texture: pf.tf.CreateTexture(pType, pColor),
			},
		}
	case types.Pawn:
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
