package game

type queen struct {
	*Piece
}

func (q *queen) GetPiece() *Piece {
	return q.Piece
}

func (q *queen) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	board[dstRow][dstCol] = q
	board[q.Row][q.Col] = nil
	q.UpdatePiecePosition(dstRow, dstCol)
}
