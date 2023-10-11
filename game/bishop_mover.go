package game

//import (
//	"math"
//)
//
//type BishopMover struct{}
//
//func (m *BishopMover) Move(srcRow int, srcCol int, dstRow int, dstCol int, board *ChessBoard) {
//	p := board[srcRow][srcCol]
//
//	if !m.IsValidMove(srcRow, srcCol, dstRow, dstCol, board) {
//		p.UpdatePiecePosition(srcRow, srcCol)
//		return
//	}
//	board[dstRow][dstCol] = p
//	board[srcRow][srcCol] = nil
//	p.UpdatePiecePosition(dstRow, dstCol)
//}
//
//func (m *BishopMover) IsValidMove(srcRow int, srcCol int, dstRow int, dstCol int, board *ChessBoard) bool {
//	piece := board[srcRow][srcCol]
//
//	// Check valid step
//	if dstCol == srcCol || dstRow == srcRow {
//		return false
//	}
//
//	if math.Abs(float64(dstCol-srcCol)) != math.Abs(float64(dstRow-srcRow)) {
//		return false
//	}
//
//	// Check clear path to destination
//	deltaCol := int(float64(dstCol-srcCol) / math.Abs(float64(dstCol-srcCol)))
//	deltaRow := int(float64(dstRow-srcRow) / math.Abs(float64(dstRow-srcRow)))
//	i := 1
//	for srcCol+deltaCol*i != dstCol && srcRow+deltaRow*i != dstRow {
//		if board[srcRow+deltaRow*i][srcCol+deltaCol*i] != nil {
//			return false
//		}
//		i++
//	}
//
//	// Ok if empty.
//	if board[dstRow][dstCol] == nil {
//		return true
//	}
//
//	// Check if destination doesn't contain ally piece
//	if piece.IsBlack() && board[dstRow][dstCol].IsBlack() || piece.IsWhite() && board[dstRow][dstCol].IsWhite() {
//		return false
//	}
//
//	return true
//}
