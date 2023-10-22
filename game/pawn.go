package game

type pawn struct {
	*Piece
	MoveCount int
}

func (p *pawn) GetPiece() *Piece {
	return p.Piece
}

func (p *pawn) Move(dstRow int, dstCol int, board *[8][8]PieceInterface) {
	board[dstRow][dstCol] = p
	board[p.Row][p.Col] = nil
	p.UpdatePiecePosition(dstRow, dstCol)
	p.MoveCount++
}

func (p *pawn) isValidMove(dstRow int, dstCol int, board *[8][8]PieceInterface) bool {
	return true
}

func (p *pawn) isValidFirstMove(dstRow int, dstCol int, board *[8][8]PieceInterface) bool {
	if p.MoveCount != 0 {
		return false
	}
	return true
}
