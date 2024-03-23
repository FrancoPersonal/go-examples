package repository

import (
	"errors"

	"github.com/FrancoPersonal/samples/hexa-hello/src/internal/app/domain"
	"github.com/FrancoPersonal/samples/hexa-hello/src/internal/app/ports"
)

var (
	games map[string]domain.Game
)

type GameHardCore struct {
}

func NewGameHardCore() (ports.GameRepository, error) {
	games = make(map[string]domain.Game)
	return GameHardCore{}, nil
}

func (repo GameHardCore) CreateGame(name string) (*domain.Game, error) {
	if _, ok := games[name]; ok {
		return nil, errors.New("already Exist")
	}
	return &domain.Game{
		Name: name,
	}, nil
}

func (repo GameHardCore) GetGame(name string) (*domain.Game, error) {
	if value, ok := games[name]; !ok {
		return nil, errors.New("not found")
	} else {
		return &value, nil
	}
}
