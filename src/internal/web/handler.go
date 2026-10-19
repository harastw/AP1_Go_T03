package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"game-service/internal/api"
	"game-service/internal/app"
	"game-service/internal/domain"

	"github.com/google/uuid"
)

type Handler struct {
	service app.GameAppService
}

func NewHandler(service app.GameAppService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gameIDStr := strings.TrimPrefix(r.URL.Path, "/game/")
	gameIDStr = strings.TrimSpace(gameIDStr)

	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		http.Error(w, "invalid game UUID", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req api.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	storedGame, err := h.service.GetGame(gameID.String())
	if err != nil {
		if errors.Is(err, domain.ErrGameNotFound) {
			storedGame = &domain.Game{
				ID:    gameID,
				Board: [3][3]domain.Cell{},
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	incomingGame := &domain.Game{
		ID:    storedGame.ID,
		Board: api.ToDomainBoard(req.Board),
	}

	if err := h.service.ValidateBoard(storedGame, incomingGame); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if h.service.CheckGameOver(incomingGame.Board) {
		_ = h.service.SaveGame(incomingGame)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(api.ToResponse(incomingGame))
		return
	}

	nextBoard, err := h.service.GetNextMove(incomingGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	incomingGame.Board = nextBoard

	_ = h.service.SaveGame(incomingGame)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api.ToResponse(incomingGame))
}
