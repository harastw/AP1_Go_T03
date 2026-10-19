package domain

import "github.com/google/uuid"

type Cell int

const (
	CellEmpty Cell = iota
	CellPlayer
	CellComputer
)

type Board [3][3]Cell

type Game struct {
	ID    uuid.UUID
	Board Board
}
