package helditems

import (
	"fmt"
	"github.com/samber/lo"
)

// GetHeldItem fetches a held item
func GetHeldItem(heldItemName string, numStacksPtr *int) (heldItem HeldItem, err error) {
	if !IsHeldItemPlayable(heldItemName) {
		err = fmt.Errorf("held item %s is not playable", heldItemName)
		return
	}

	switch heldItemName {
	case AeosCookieName:
		numStacks := lo.FromPtr(numStacksPtr)
		heldItem, err = NewAeosCookie(numStacks)
	case BuddyBarrierName:
		heldItem, err = NewBuddyBarrier()
	case EnergyAmplifierName:
		heldItem, err = NewEnergyAmplifier()
	case FocusBandName:
		heldItem, err = NewFocusBand()
	case RapidFireScarfName:
		heldItem, err = NewRapidFireScarf()
	case RockyHelmetName:
		heldItem, err = NewRockyHelmet()
	case ScoreShieldName:
		heldItem, err = NewScoreShield()
	case SpecialAttackSpecsName:
		numStacks := lo.FromPtr(numStacksPtr)
		heldItem, err = NewSpecialAttackSpecs(numStacks)
	case AssaultVestName:
		heldItem, err = NewAssaultVest()
	case ChoiceSpecsName:
		heldItem, err = NewChoiceSpecs()
	case ExpShareName:
		heldItem, err = NewExpShare()
	case LeftoversName:
		heldItem, err = NewLeftovers()
	case RazorClawName:
		heldItem, err = NewRazorClaw()
	case RustedSwordName:
		heldItem, err = NewRustedSword()
	case ShellBellName:
		heldItem, err = NewShellBell()
	case WeaknessPolicyName:
		heldItem, err = NewWeaknessPolicy()
	case AttackWeightName:
		numStacks := lo.FromPtr(numStacksPtr)
		heldItem, err = NewAttackWeight(numStacks)
	case DrainCrownName:
		heldItem, err = NewDrainCrown()
	case FloatStoneName:
		heldItem, err = NewFloatStone()
	case MuscleBandName:
		heldItem, err = NewMuscleBand()
	case RescueHoodName:
		heldItem, err = NewRescueHood()
	case ScopeLensName:
		heldItem, err = NewScopeLens()
	case SlickSpoonName:
		heldItem, err = NewSlickSpoon()
	case WiseGlassesName:
		heldItem, err = NewWiseGlasses()
	}

	return
}
