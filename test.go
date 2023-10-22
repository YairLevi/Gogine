package main

//
//type Piece interface {
//	Move()
//}
//
//type GenericPiece struct {
//	Row int
//	Col int
//}
//
//func (s GenericPiece) Move() {
//	return
//}
//
//type Rook struct {
//	GenericPiece
//	HasMoved bool
//}
//
//func (c Rook) Move() {
//	c.HasMoved = true
//}
//
//type Pawn struct {
//	GenericPiece
//	MoveCount int
//}
//
//func (p Pawn) Move() {
//	p.MoveCount++
//}
//
//func main() {
//	shapes := []Piece{Rook{}, Pawn{}}
//
//	piece := shapes[0]
//	gp := piece.(GenericPiece)
//
//	rook, isRook := piece.(Rook)
//	if isRook {
//		rook.HasMoved
//	}
//}
