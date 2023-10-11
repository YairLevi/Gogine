package game

type pawn struct {
	*Piece
}

func (p *pawn) GetPiece() *Piece {
	return p.Piece
}

func (p *pawn) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	board[dstRow][dstCol] = p
	board[p.Row][p.Col] = nil
	p.UpdatePiecePosition(dstRow, dstCol)
}
