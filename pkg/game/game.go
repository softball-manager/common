package game

import (
	"github.com/softball-manager/common/pkg/player"
)

type Game struct {
	TeamRuns       int  `json:"teamRuns"`
	OppositionRuns int  `json:"oppositionRuns"`
	Win            bool `json:"win"`
	Lineup         Lineup
}

type Lineup struct {
	BattingOrder     []player.Player
	FieldingByInning []DefensivePositions
}

type DefensivePositions struct {
	Inning    int        `json:"inning"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Player   player.Player `json:"player"`
	Position string        `json:"position"`
}
