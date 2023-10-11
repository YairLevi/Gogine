package game

type DragManager struct {
	DraggedPiece *Piece
}

func (dm *DragManager) StartDrag(p *Piece) {
	if dm.DraggedPiece == nil {
		p.Dragging = true
		dm.DraggedPiece = p
	}
}

func (dm *DragManager) StopDrag() {
	dm.DraggedPiece.Dragging = false
	dm.DraggedPiece = nil
}
