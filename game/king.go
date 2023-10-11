package game

type king struct {
	*Piece
}

func (k *king) GetPiece() *Piece {
	return k.Piece
}

func (k *king) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	board[dstRow][dstCol] = k
	board[k.Row][k.Col] = nil
	k.UpdatePiecePosition(dstRow, dstCol)
}
