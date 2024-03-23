package domain

// Game Struct Domain
type Game struct {
	Name    string
	Type    TypeGame
	Ranking int
}

// Game Clarification
type TypeGame int64

const (
	Undefined TypeGame = iota
	Rpg
	Shooter
	Fight
	SurvivorHorror
)
