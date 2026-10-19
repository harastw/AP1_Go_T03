package domain

type GameService interface {
	GetNextMove(game *Game) (Board, error)
	ValidateBoard(stored *Game, incoming *Game) error
	CheckGameOver(board Board) bool
}
