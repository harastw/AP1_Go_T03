package infrastructure

type GameRepository interface {
	Save(game *StorageGame) error
	Get(id string) (*StorageGame, error)
}
