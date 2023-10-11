package game

import "math"

type rook struct {
	*Piece
	hasMoved bool
}

func (r *rook) GetPiece() *Piece {
	return r.Piece
}

func (r *rook) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	if !r.isValidMove(r.Row, r.Col, dstRow, dstCol, board) {
		r.UpdatePiecePosition(r.Row, r.Col)
		return
	}
	r.hasMoved = true
	board[dstRow][dstCol] = r
	board[r.Row][r.Col] = nil
	r.UpdatePiecePosition(dstRow, dstCol)
}

func (r *rook) isValidMove(srcRow int, srcCol int, dstRow int, dstCol int, board *[8][8]PieceInterface) bool {
	dstPiece := board[dstRow][dstCol]

	// didn't move
	if dstCol == srcCol && dstRow == srcRow {
		return false
	}

	// Check valid rook step
	if dstCol != srcCol && dstRow != srcRow {
		return false
	}

	// Check clear path to destination
	if dstRow == srcRow {
		deltaCol := int(float64(dstCol-srcCol) / math.Abs(float64(dstCol-srcCol)))
		i := 1
		for srcCol+deltaCol*i != dstCol {
			if board[srcRow][srcCol+deltaCol*i] != nil {
				return false
			}
			i++
		}
	} else {
		deltaRow := int(float64(dstRow-srcRow) / math.Abs(float64(dstRow-srcRow)))
		i := 1
		for srcRow+deltaRow*i != dstRow {
			if board[srcRow+deltaRow*i][srcCol] != nil {
				return false
			}
			i++
		}
	}

	// Ok if empty.
	if board[dstRow][dstCol] == nil {
		return true
	}

	// Check if destination doesn't contain ally piece
	if r.Color == dstPiece.GetPiece().Color {
		return false
	}

	return true
}
