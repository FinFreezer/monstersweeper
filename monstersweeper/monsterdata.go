package monstersweeper

func newImp(id int) *Monster {
	m := Monster{
		Name:         "Imp",
		Health:       5,
		Strength:     6,
		Dexterity:    10,
		Intelligence: 8,
		MonsterId:    id,
	}
	return &m
}

func newSkeleton(id int) *Monster {
	m := Monster{
		Name:         "Skeleton",
		Health:       8,
		Strength:     10,
		Dexterity:    5,
		Intelligence: 5,
		MonsterId:    id,
	}
	return &m
}

func newZombie(id int) *Monster {
	m := Monster{
		Name:         "Zombie",
		Health:       12,
		Strength:     8,
		Dexterity:    4,
		Intelligence: 2,
		MonsterId:    id,
	}
	return &m
}

func newWitch(id int) *Monster {
	m := Monster{
		Name:         "Witch",
		Health:       4,
		Strength:     6,
		Dexterity:    10,
		Intelligence: 10,
		MonsterId:    id,
	}
	return &m
}

func newOrc(id int) *Monster {
	m := Monster{
		Name:         "Orc",
		Health:       14,
		Strength:     10,
		Dexterity:    4,
		Intelligence: 2,
		MonsterId:    id,
	}
	return &m
}
