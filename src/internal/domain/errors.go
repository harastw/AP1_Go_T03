package domain

import "errors"

var (
	ErrGameNotFound     = errors.New("game not found")
	ErrBoardTampered    = errors.New("previous moves were changed")
	ErrInvalidMoveCount = errors.New("exactly one new move expected")
	ErrGameFinished     = errors.New("game already finished")
)
