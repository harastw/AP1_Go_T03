package api

type GameRequest struct {
	Board [3][3]int `json:"board"`
}

type GameResponse struct {
	ID    string    `json:"id"`
	Board [3][3]int `json:"board"`
}
