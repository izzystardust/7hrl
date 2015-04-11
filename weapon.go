package main

type weapon struct {
	name   string
	attack func() (int, attackType)
	sym    rune
}

func newSword() weapon {
	return weapon{
		name:   "sword",
		attack: func() (int, attackType) { return 8, attackNormal },
		sym:    '/',
	}
}

func (w weapon) Draw() rune {
	return w.sym
}

func (w weapon) Name() string {
	return w.name
}

func (w weapon) Use(s *state) {
	p := s.Player()
	p.items = append(p.items, p.weapon)
	p.weapon = w
}

func newUnarmed() weapon {
	return weapon{
		name:   "unarmed",
		attack: func() (int, attackType) { return 1, attackNormal },
	}
}
