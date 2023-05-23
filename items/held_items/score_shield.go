package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// ScoreShield is a held item
type ScoreShield struct {
	HeldItemData
}

// NewScoreShield creates a new ScoreShield held item
func NewScoreShield() (scoreShield *ScoreShield, err error) {
	scoreShieldData, err := FetchHeldItemData(ScoreShieldName)
	if err != nil {
		fmt.Println("Error fetching held item data for Score Shield")
		return
	}

	scoreShield = &ScoreShield{
		HeldItemData: scoreShieldData,
	}
	return
}

func (s *ScoreShield) GetStatBoosts() stats.Stats {
	return s.Stats
}

func (s *ScoreShield) Activate(originalStats stats.Stats) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
