package services

import (
	"github.com/FrancoPersonal/go-examples/samples/hexa-hello/src/internal/adapters/repository"
	"github.com/FrancoPersonal/go-examples/samples/hexa-hello/src/internal/app/domain"
	"github.com/FrancoPersonal/go-examples/samples/hexa-hello/src/internal/app/ports"
)

type GameService struct {
	repository ports.GameRepository
}

func NewGameService() GameService {
	return GameService{repository.NewGameHardCore()}
}

func (service GameService) CreateGame(name string) (*domain.Game, error) {
	return service.repository.CreateGame(name)
}
func (service GameService) GetGame(name string) (*domain.Game, error) {
	return service.repository.GetGame(name)
}
