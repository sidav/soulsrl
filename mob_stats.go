package main

type mobStats struct {
	vitality  int // hitpoints
	endurance int // stamina
	dexterity int
	strength  int
}

func (m *mob) getMaxHitpoints() int {
	return m.stats.dexterity
}

func (m *mob) getMaxStamina() int {
	return m.stats.endurance
}
