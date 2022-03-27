package main

const (
	ACTIONTYPE_ATTACK = iota
	ACTIONTYPE_MOVE
)

type action struct {
	x, y        int
	tickToOccur int
	actionType  int

	toHitRoll, damageRoll int

	owner *mob
}

func (b *battlefield) applyActions() {
	for i := 0; i < len(b.actions); i++ {
		if b.currentTick >= b.actions[i].tickToOccur {
			action := b.actions[i]
			mobAtCoords := b.getMobPresentAt(action.x, action.y)
			if mobAtCoords != nil && mobAtCoords != action.owner &&
				mobAtCoords.wasAlreadyAffectedByActionBy != action.owner {

				mobAtCoords.wasAlreadyAffectedByActionBy = action.owner
				b.harmMob(action.owner, action.toHitRoll, action.damageRoll, mobAtCoords)
			}
		}
	}
}

func (b *battlefield) cleanupActions() {
	for i := 0; i < len(b.actions); i++ {
		shouldRemove := b.currentTick >= b.actions[i].tickToOccur || b.actions[i].owner.hitpoints <= 0
		if shouldRemove {
			b.actions[i] = b.actions[len(b.actions)-1]
			b.actions = append(b.actions[:len(b.actions)-1])
			i -= 1
			continue
		}
	}
}

func (b *battlefield) harmMob(attacker *mob, toHit, dmg int, target *mob) {
	targetArmorClass := 0
	targetDamageReduction := 0
	if target.body != nil {
		targetArmorClass = target.body.AsArmor.GetData().ArmorClass
		targetDamageReduction = target.body.AsArmor.GetData().DamageReduction
	}
	if toHit > targetArmorClass {
		dmg -= targetDamageReduction
		if dmg < 1 {
			dmg = 1
		}
		target.hitpoints -= dmg
		log.AppendMessagef("%s hits %s (%d dmg)!", attacker.name, target.name, dmg)
	} else {
		log.AppendMessagef("%s misses %s!", attacker.name, target.name)
	}
}
