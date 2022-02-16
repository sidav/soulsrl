package main

type action struct {
	x, y        int
	tickToOccur int
	owner       *mob
}

func (b *battlefield) applyActions() {
	for i := 0; i < len(b.actions); i++ {
		if b.currentTick >= b.actions[i].tickToOccur {
			action := b.actions[i]
			mobAtCoords := b.getMobPresentAt(action.x, action.y)
			if mobAtCoords != nil && mobAtCoords != action.owner &&
				mobAtCoords.wasAlreadyAffectedByActionBy != action.owner {

				mobAtCoords.wasAlreadyAffectedByActionBy = action.owner
				log.AppendMessagef("%s hits %s!", action.owner.name, mobAtCoords.name)
				mobAtCoords.hitpoints--
			}

			// remove actions
			b.actions[i] = b.actions[len(b.actions)-1]
			b.actions = append(b.actions[:len(b.actions)-1])
			i -= 1
			continue
		}
	}
}
