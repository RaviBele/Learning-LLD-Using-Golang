package main

type PieceType int

const (
	X PieceType = iota + 1
	O
)

func (p PieceType) String() string {
	return [...]string{"X", "O"}[p-1]
}

type IPiece interface {
	getPieceType() PieceType
}

type Piece struct {
	pieceType PieceType
}

func NewPiece(pieceType PieceType) *Piece {
	return &Piece{pieceType: pieceType}
}

func (p *Piece) getPieceType() PieceType {
	return p.pieceType
}

type PieceX struct {
	Piece
}

func NewPieceX() *PieceX {
	return &PieceX{Piece: *NewPiece(X)}
}

type PieceO struct {
	Piece
}

func NewPieceO() *PieceO {
	return &PieceO{Piece: *NewPiece(O)}
}
