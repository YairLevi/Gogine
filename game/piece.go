package game

import (
	"github.com/gen2brain/raylib-go/raylib"
	tx "main/textures"
	"main/types"
	"math"
)

type PieceInterface interface {
	GetPiece() *Piece
	Move(dstRow int, dstCol int, board *[8][8]PieceInterface)
}

type OnMoveCallback = func(newCol int, newRow int)

type Piece struct {
	// for dragging
	Rect            rl.Rectangle
	Dragging        bool
	DraggingManager *DragManager

	// game-specific, for Chess
	Row             int
	Col             int
	Type            types.PieceType
	Color           types.PieceColor
	Texture         *tx.Texture
	OnMoveCallbacks []OnMoveCallback
}

func (p *Piece) AddOnMoveCallback(callback OnMoveCallback) {
	p.OnMoveCallbacks = append(p.OnMoveCallbacks, callback)
}

func (p *Piece) onDragStop() {
	p.DraggingManager.StopDrag()

	mousePos := rl.GetMousePosition()
	newCol := int(math.Floor(float64(mousePos.X / tx.TileSize)))
	newRow := int(math.Floor(float64(mousePos.Y / tx.TileSize)))

	dragOutOfBounds := newCol > 7 || newRow > 7 || newCol < 0 || newRow < 0

	if dragOutOfBounds {
		p.Rect.X = float32(p.Col * tx.TileSize)
		p.Rect.Y = float32(p.Row * tx.TileSize)
		return
	}

	for _, callback := range p.OnMoveCallbacks {
		callback(newCol, newRow)
	}
}

func (p *Piece) UpdatePiecePosition(newRow int, newCol int) {
	p.Col = newCol
	p.Row = newRow
	p.Rect.X = float32(p.Col * tx.TileSize)
	p.Rect.Y = float32(p.Row * tx.TileSize)

}

func (p *Piece) Input() {
	if !rl.IsMouseButtonDown(rl.MouseLeftButton) {
		if p.Dragging {
			p.onDragStop()
		}
		return
	}

	mousePosition := rl.GetMousePosition()
	withinX := p.Rect.X < mousePosition.X && mousePosition.X < p.Rect.X+p.Rect.Width
	withinY := p.Rect.Y < mousePosition.Y && mousePosition.Y < p.Rect.Y+p.Rect.Height
	noPieceIsDragged := p.DraggingManager.DraggedPiece == nil

	if withinX && withinY && noPieceIsDragged {
		p.DraggingManager.StartDrag(p)
	}
}

func (p *Piece) Render() {
	if p.Dragging {
		mousePosition := rl.GetMousePosition()
		p.Rect.X = mousePosition.X - p.Rect.Width/2
		p.Rect.Y = mousePosition.Y - p.Rect.Height/2
	}
	p.Texture.Draw(p.Rect.X, p.Rect.Y)
}
