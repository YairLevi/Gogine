package types

type PieceColor int
type PieceType int

const (
	WHITE PieceColor = iota
	BLACK
)
const (
	PAWN PieceType = iota + 1
	BISHOP
	KNIGHT
	ROOK
	QUEEN
	KING
)
