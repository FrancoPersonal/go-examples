package ports

import (
	"github.com/FrancoPersonal/go-examples/samples/hexa-hello/src/internal/app/domain"
)

type GameRepository interface {
	CreateGame(name string) (*domain.Game, error)
	GetGame(name string) (*domain.Game, error)
}
