package app

import (
	"game-service/internal/domain"
	"game-service/internal/infrastructure"
)

type GameAppService interface {
	domain.GameService
	GetGame(id string) (*domain.Game, error)
	SaveGame(game *domain.Game) error
}

type Service struct {
	repo  infrastructure.GameRepository
	logic domain.GameService
}

func NewService(repo infrastructure.GameRepository, logic domain.GameService) GameAppService {
	return &Service{
		repo:  repo,
		logic: logic,
	}
}

func (s *Service) GetNextMove(game *domain.Game) (domain.Board, error) {
	return s.logic.GetNextMove(game)
}

func (s *Service) ValidateBoard(stored *domain.Game, incoming *domain.Game) error {
	return s.logic.ValidateBoard(stored, incoming)
}

func (s *Service) CheckGameOver(board domain.Board) bool {
	return s.logic.CheckGameOver(board)
}

func (s *Service) GetGame(id string) (*domain.Game, error) {
	stg, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return infrastructure.ToDomain(stg)
}

func (s *Service) SaveGame(game *domain.Game) error {
	return s.repo.Save(infrastructure.ToStorage(game))
}
