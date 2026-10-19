package domain

type gameLogic struct{}

func NewGameLogic() GameService {
	return &gameLogic{}
}

func (s *gameLogic) CheckGameOver(board Board) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] != CellEmpty && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		}
		if board[0][i] != CellEmpty && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return true
		}
	}
	if board[0][0] != CellEmpty && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	if board[0][2] != CellEmpty && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == CellEmpty {
				return false
			}
		}
	}
	return true
}

func (s *gameLogic) ValidateBoard(stored *Game, incoming *Game) error {
	diffCount := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if stored.Board[i][j] != incoming.Board[i][j] {
				if stored.Board[i][j] != CellEmpty {
					return ErrBoardTampered
				}
				if incoming.Board[i][j] != CellPlayer {
					return ErrInvalidMoveCount
				}
				diffCount++
			}
		}
	}
	if diffCount != 1 {
		return ErrInvalidMoveCount
	}
	return nil
}

func (s *gameLogic) GetNextMove(game *Game) (Board, error) {
	board := game.Board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == CellEmpty {
				board[i][j] = CellComputer
				return board, nil
			}
		}
	}
	return board, nil
}
