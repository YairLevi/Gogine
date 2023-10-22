package game

import "math"

type king struct {
	*Piece
}

func (k *king) GetPiece() *Piece {
	return k.Piece
}

func (k *king) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	if !k.isValidMove(dstRow, dstCol, board) {
		k.UpdatePiecePosition(k.Row, k.Col)
		return
	}
	board[dstRow][dstCol] = k
	board[k.Row][k.Col] = nil
	k.UpdatePiecePosition(dstRow, dstCol)
}

func (k *king) isValidMove(dstRow int, dstCol int, board *[8][8]PieceInterface) bool {
	if dstRow == k.Row && dstCol == k.Col {
		return false
	}

	if board[dstRow][dstCol] != nil && board[dstRow][dstCol].GetPiece().Color == k.Color {
		return false
	}

	rowGap := math.Abs(float64(k.Row - dstRow))
	colGap := math.Abs(float64(k.Col - dstCol))

	if colGap <= 1 && rowGap <= 1 {
		return true
	}
	return false
}
