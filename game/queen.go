package game

import "math"

type Queen struct {
	*Piece
}

func (q *Queen) GetPiece() *Piece {
	return q.Piece
}

func (q *Queen) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	if !q.isValidMove(dstRow, dstCol, board) {
		q.UpdatePiecePosition(q.Row, q.Col)
		return
	}
	board[dstRow][dstCol] = q
	board[q.Row][q.Col] = nil
	q.UpdatePiecePosition(dstRow, dstCol)
}

func (q *Queen) isValidMove(dstRow int, dstCol int, board *[8][8]PieceInterface) bool {
	return q.isValidRookMove(q.Row, q.Col, dstRow, dstCol, board) ||
		q.isValidBishopMove(q.Row, q.Col, dstRow, dstCol, board)
}

func (q *Queen) isValidRookMove(srcRow int, srcCol int, dstRow int, dstCol int, board *[8][8]PieceInterface) bool {
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
	if q.Color == dstPiece.GetPiece().Color {
		return false
	}

	return true
}

func (q *Queen) isValidBishopMove(srcRow int, srcCol int, dstRow int, dstCol int, board *[8][8]PieceInterface) bool {
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
	if q.Color == dstPiece.GetPiece().Color {
		return false
	}

	return true
}
