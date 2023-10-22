package game

import (
	"math"
)

type Bishop struct {
	*Piece
}

func (b *Bishop) GetPiece() *Piece {
	return b.Piece
}

func (b *Bishop) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	if !b.isValidMove(b.Row, b.Col, dstRow, dstCol, board) {
		b.UpdatePiecePosition(b.Row, b.Col)
		return
	}

	board[dstRow][dstCol] = b
	board[b.Row][b.Col] = nil
	b.UpdatePiecePosition(dstRow, dstCol)
}

func (b *Bishop) isValidMove(srcRow int, srcCol int, dstRow int, dstCol int, board *[8][8]PieceInterface) bool {
	dstPiece := board[dstRow][dstCol]

	// Check valid step
	if dstCol == srcCol || dstRow == srcRow {
		return false
	}

	if math.Abs(float64(dstCol-srcCol)) != math.Abs(float64(dstRow-srcRow)) {
		return false
	}

	// Check clear path to destination
	deltaCol := int(float64(dstCol-srcCol) / math.Abs(float64(dstCol-srcCol)))
	deltaRow := int(float64(dstRow-srcRow) / math.Abs(float64(dstRow-srcRow)))
	i := 1
	for srcCol+deltaCol*i != dstCol && srcRow+deltaRow*i != dstRow {
		if board[srcRow+deltaRow*i][srcCol+deltaCol*i] != nil {
			return false
		}
		i++
	}

	// Ok if empty.
	if board[dstRow][dstCol] == nil {
		return true
	}

	// Check if destination doesn't contain ally piece
	if b.GetPiece().Color == dstPiece.GetPiece().Color {
		return false
	}

	return true
}
