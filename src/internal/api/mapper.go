package api

import "game-service/internal/domain"

func ToDomainBoard(b [3][3]int) domain.Board {
	board := domain.Board{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = domain.Cell(b[i][j])
		}
	}
	return board
}

func ToResponse(game *domain.Game) GameResponse {
	boardInt := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boardInt[i][j] = int(game.Board[i][j])
		}
	}
	return GameResponse{
		ID:    game.ID.String(),
		Board: boardInt,
	}
}
