package game

type bishop struct {
	*Piece
}

func (b *bishop) GetPiece() *Piece {
	return b.Piece
}

func (b *bishop) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	board[dstRow][dstCol] = b
	board[b.Row][b.Col] = nil
	b.UpdatePiecePosition(dstRow, dstCol)
}
