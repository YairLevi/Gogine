package game

//type QueenMover struct {
//	RookMover
//	BishopMover
//}
//
//func (m *QueenMover) Move(srcRow int, srcCol int, dstRow int, dstCol int, board *ChessBoard) {
//	p := board[srcRow][srcCol]
//
//	if m.RookMover.IsValidMove(srcRow, srcCol, dstRow, dstCol, board) ||
//		m.BishopMover.IsValidMove(srcRow, srcCol, dstRow, dstCol, board) {
//
//		board[dstRow][dstCol] = p
//		board[srcRow][srcCol] = nil
//		p.UpdatePiecePosition(dstRow, dstCol)
//		return
//	}
//	p.UpdatePiecePosition(srcRow, srcCol)
//}
//
//func (m *QueenMover) IsValidMove(srcRow int, srcCol int, dstRow int, dstCol int, board *ChessBoard) bool {
//	return m.RookMover.IsValidMove(srcRow, srcCol, dstRow, dstCol, board) ||
//		m.BishopMover.IsValidMove(srcRow, srcCol, dstRow, dstCol, board)
//}
