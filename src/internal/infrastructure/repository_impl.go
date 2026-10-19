package infrastructure

import "game-service/internal/domain"

type repository struct {
	storage *GameStorage
}

func NewRepository(storage *GameStorage) GameRepository {
	return &repository{storage: storage}
}

func (r *repository) Save(game *StorageGame) error {
	r.storage.Save(game.ID, game)
	return nil
}

func (r *repository) Get(id string) (*StorageGame, error) {
	game, ok := r.storage.Get(id)
	if !ok {
		return nil, domain.ErrGameNotFound
	}
	return game, nil
}
