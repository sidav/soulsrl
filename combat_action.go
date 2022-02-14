package main

type action struct {
	x, y        int
	tickToOccur int
	owner       *mob
}

func (b *battlefield) applyActions() {
	for i := 0; i < len(b.actions); i++ {
		if b.currentTick >= b.actions[i].tickToOccur {
			mobAtCoords := b.getMobPresentAt(b.actions[i].x, b.actions[i].y)
			if mobAtCoords != nil && mobAtCoords != b.actions[i].owner {
				log.AppendMessagef("%s hits %s!", b.actions[i].owner.name, mobAtCoords.name)
			}
			b.actions[i] = b.actions[len(b.actions)-1]
			b.actions = append(b.actions[:len(b.actions)-1])
			i -= 1
			continue
		}
	}
}
