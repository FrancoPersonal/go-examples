package repository

import "github.com/FrancoPersonal/samples/hexa-hello/src/internal/app/ports"

type GameHardCore struct {
}

func NewGameHardCore() (ports.GameRepository, error) {
	return GameHardCore{}, nil
}
