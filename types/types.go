package types

type PieceColor int
type PieceType int

const (
	White PieceColor = iota
	Black
)
const (
	Pawn PieceType = iota + 1
	Bishop
	Knight
	Rook
	Queen
	King
)
