package team

import (
	"github.com/softball-manager/common/pkg/game"
	"github.com/softball-manager/common/pkg/player"
)

type Team struct {
	TeamName string          `json:"teamName"`
	Players  []player.Player `json:"players"`
	Games    []game.Game     `json:"games"`
}
