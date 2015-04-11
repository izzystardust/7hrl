package main

type potion struct {
	effect func(s *state, p *player)
	name   string
}

func (p potion) Name() string {
	return p.name
}

func (p potion) Use(s *state) {
	p.effect(s, s.Player())
}

func (p potion) Draw() rune {
	return '⬓'
}
