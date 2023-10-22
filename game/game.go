package game

import (
	"github.com/gen2brain/raylib-go/raylib"
	"main/types"
)

var (
	light = rl.NewColor(227, 193, 111, 255)
	dark  = rl.NewColor(184, 139, 74, 255)
)

type ChessBoard = [8][8]PieceInterface

type Game struct {
	X          int
	Y          int
	TileSize   int
	LightColor rl.Color
	DarkColor  rl.Color
	BoardState *ChessBoard
}

var (
	initPieces = [8][8]types.PieceType{
		{types.ROOK, types.KNIGHT, types.BISHOP, types.QUEEN, types.KING, types.BISHOP, types.KNIGHT, types.ROOK},
		{types.PAWN, types.PAWN, types.PAWN, types.PAWN, types.PAWN, types.PAWN, types.PAWN, types.PAWN},
		{},
		{},
		{},
		{},
		{types.PAWN, types.PAWN, types.PAWN, types.PAWN, types.PAWN, types.PAWN, types.PAWN, types.PAWN},
		{types.ROOK, types.KNIGHT, types.BISHOP, types.QUEEN, types.KING, types.BISHOP, types.KNIGHT, types.ROOK},
	}
	initColors = [8][8]types.PieceColor{
		{types.BLACK, types.BLACK, types.BLACK, types.BLACK, types.BLACK, types.BLACK, types.BLACK, types.BLACK},
		{types.BLACK, types.BLACK, types.BLACK, types.BLACK, types.BLACK, types.BLACK, types.BLACK, types.BLACK},
		{},
		{},
		{},
		{},
		{types.WHITE, types.WHITE, types.WHITE, types.WHITE, types.WHITE, types.WHITE, types.WHITE, types.WHITE},
		{types.WHITE, types.WHITE, types.WHITE, types.WHITE, types.WHITE, types.WHITE, types.WHITE, types.WHITE},
	}
)

func move(pieceI PieceInterface, boardState *ChessBoard) OnMoveCallback {
	return func(newCol int, newRow int) {
		pieceI.Move(newRow, newCol, boardState)
	}
}

func NewGame(posX, posY, tileSize int) *Game {
	factory := NewPieceFactory()
	dragManager := &DragManager{DraggedPiece: nil}
	boardState := &ChessBoard{}

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			pType, pColor := initPieces[row][col], initColors[row][col]
			pi := factory.Create(pType, pColor)
			if pi == nil {
				boardState[row][col] = nil
				continue
			}

			p := pi.GetPiece()
			p.AddOnMoveCallback(move(pi, boardState))

			p.Row = row
			p.Col = col
			p.DraggingManager = dragManager
			p.Rect = rl.NewRectangle(
				float32(col*tileSize),
				float32(row*tileSize),
				float32(tileSize),
				float32(tileSize),
			)
			boardState[row][col] = pi
		}
	}

	return &Game{
		X:          posX,
		Y:          posY,
		BoardState: boardState,
		TileSize:   tileSize,
		LightColor: light,
		DarkColor:  dark,
	}
}

func (g *Game) getPiecesArray() []*Piece {
	var pieces []*Piece
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			p := g.BoardState[row][col]
			if g.BoardState[row][col] != nil {
				pieces = append(pieces, p.GetPiece())
			}
		}
	}
	return pieces
}

func (g *Game) drawBoard() {
	for row := 0; row < 8; row++ {
		tileY := row*g.TileSize + g.Y
		for col := 0; col < 8; col++ {
			tileX := col*g.TileSize + g.X
			isLight := col%2 == 0 && row%2 == 0 || col%2 == 1 && row%2 == 1

			tileColor := g.DarkColor
			if isLight {
				tileColor = g.LightColor
			}
			rl.DrawRectangle(int32(tileX), int32(tileY), int32(g.TileSize), int32(g.TileSize), tileColor)
		}
	}
}

func (g *Game) drawPieces() {
	var draggedPiece PieceInterface = nil
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			p := g.BoardState[row][col]
			if p == nil {
				continue
			}
			if p.GetPiece().Dragging {
				draggedPiece = p
			} else {
				p.GetPiece().Render()
			}
		}
	}
	if draggedPiece != nil {
		draggedPiece.GetPiece().Render()
	}
}

func (g *Game) input() {
	pieces := g.getPiecesArray()
	for _, p := range pieces {
		p.Input()
	}
}

func (g *Game) render() {
	g.drawBoard()
	g.drawPieces()
}

func (g *Game) Run() {
	g.input()
	g.render()
}
