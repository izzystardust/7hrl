package main

type oyster struct {
	hp    int
	x, y  int
	level int
}

func newOyster() *oyster {
	return &oyster{10, 0, 0, 0}
}

func (o oyster) Name() string {
	return "Oyster"
}

func (o oyster) Draw() rune {
	return 'o'
}

func (o oyster) CurPos() (int, int) {
	return o.x, o.y
}

func (o oyster) Update(s *state) updateResult {
	if o.hp < 1 {
		debug("oyster should die")
		return shouldDie
	}
	return none
}

func (o *oyster) Attack() (int, attackType) {
	return 1, attackNormal
}

func (o *oyster) Defend(amount int, ty attackType) {
	o.hp -= amount
	debug("oyster took", amount, "to", o.hp)
}

func (o *oyster) SendToLevel(l int) {
	o.level = l
}

func (o *oyster) CurrentLevel() int {
	return o.level
}
