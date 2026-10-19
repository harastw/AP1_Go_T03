package infrastructure

import (
	"game-service/internal/domain"

	"github.com/google/uuid"
)

func ToStorage(g *domain.Game) *StorageGame {
	boardInt := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boardInt[i][j] = int(g.Board[i][j])
		}
	}
	return &StorageGame{
		ID:    g.ID.String(),
		Board: boardInt,
	}
}

func ToDomain(s *StorageGame) (*domain.Game, error) {
	id, err := parseUUID(s.ID)
	if err != nil {
		return nil, err
	}
	board := domain.Board{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = domain.Cell(s.Board[i][j])
		}
	}
	return &domain.Game{ID: id, Board: board}, nil
}

func parseUUID(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
