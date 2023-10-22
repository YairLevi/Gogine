package game

import (
	"math"
)

type knight struct {
	*Piece
}

func (k *knight) GetPiece() *Piece {
	return k.Piece
}

func (k *knight) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	if !k.isValidMove(dstRow, dstCol, board) {
		k.UpdatePiecePosition(k.Row, k.Col)
		return
	}

	board[dstRow][dstCol] = k
	board[k.Row][k.Col] = nil
	k.UpdatePiecePosition(dstRow, dstCol)
}

func (k *knight) isValidMove(dstRow int, dstCol int, board *[8][8]PieceInterface) bool {
	if dstRow == k.Row && dstCol == k.Col {
		return false
	}

	rowGap := math.Abs(float64(dstRow - k.Row))
	colGap := math.Abs(float64(dstCol - k.Col))

	if !(rowGap == 1 && colGap == 2 || rowGap == 2 && colGap == 1) {
		return false
	}

	if board[dstRow][dstCol] == nil {
		return true
	}

	if board[dstRow][dstCol].GetPiece().Color == k.Color {
		return false
	}

	return true
}
