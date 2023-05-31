package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
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

func (s *ScoreShield) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	return s.Stats
}

func (s *ScoreShield) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
