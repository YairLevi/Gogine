package game

type knight struct {
	*Piece
}

func (k *knight) GetPiece() *Piece {
	return k.Piece
}

func (k *knight) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	board[dstRow][dstCol] = k
	board[k.Row][k.Col] = nil
	k.UpdatePiecePosition(dstRow, dstCol)
}
