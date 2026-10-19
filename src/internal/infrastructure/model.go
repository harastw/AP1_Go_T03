package infrastructure

type StorageGame struct {
	ID    string    `json:"id"`
	Board [3][3]int `json:"board"`
}
