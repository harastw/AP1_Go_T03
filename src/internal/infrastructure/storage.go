package infrastructure

import "sync"

type GameStorage struct {
	data sync.Map
}

func NewGameStorage() *GameStorage {
	return &GameStorage{}
}

func (s *GameStorage) Save(id string, game *StorageGame) {
	s.data.Store(id, game)
}

func (s *GameStorage) Get(id string) (*StorageGame, bool) {
	val, ok := s.data.Load(id)
	if !ok {
		return nil, false
	}
	return val.(*StorageGame), true
}
