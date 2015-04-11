package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type player struct {
	hp     int
	maxhp  int
	x, y   int
	items  []item
	weapon weapon
	level  int
}

func (p *player) SendToLevel(l int) {
	p.level = l
}

func (p player) CurrentLevel() int {
	return p.level
}

func newPlayer(x, y int) *player {
	return &player{
		hp:     10,
		maxhp:  10,
		x:      x,
		y:      y,
		weapon: newUnarmed(),
	}
}

func (p player) Name() string {
	return "Ralphie"
}

func (p player) Draw() rune {
	return '@'
}

func (p player) CurPos() (int, int) {
	return p.x, p.y
}

// returns true if player moved, false otherwise
func (p *player) Move(s *state, dx, dy int) bool {
	if dx == 0 && dy == 0 {
		return false
	}
	if p.x+dx > width || p.x+dx < 0 || p.y+dy > height || p.y+dy < 0 {
		return false
	}
	if s.Walkable(p.x+dx, p.y+dy) {
		p.x += dx
		p.y += dy
		return true
	}
	return false
}

func (p *player) Attack() (int, attackType) {
	return p.weapon.attack()
}

func (p *player) Defend(amount int, ty attackType) {
	p.hp -= amount
}

func (p *player) Update(s *state) updateResult {
	var dx, dy int
	moved := false
	for !moved {
		event := termbox.PollEvent()
		switch event.Ch {
		case 'h':
			dx = -1
		case 'n':
			fallthrough
		case 'j':
			dy = 1
		case 'e':
			fallthrough
		case 'k':
			dy = -1
		case 'i':
			fallthrough
		case 'l':
			dx = 1
		case 'q':
			return gameOver
		case 'u':
			moved = true
			i := pickItem(s)
			if i != nil {
				i.Use(s)
			}
		}
		moved = moved || p.Move(s, dx, dy)
		if moved {
			s.m.onEnter(s, p.x, p.y)
			s.m = &s.ms[p.level]
		}
		if !moved {
			e := s.AttackableAt(p.x+dx, p.y+dy)
			if e != nil && !isPlayer(e) {
				e.Defend(p.Attack())
				p.Defend(e.Attack())
				moved = true
				debug("Attacking", e.(*oyster))
			} else if o := s.ObjectAt(p.x+dx, p.y+dy); o != nil {
				p.items = append(p.items, o.item)
				s.messages = append(s.messages,
					statusMessage(
						fmt.Sprintf("Picked up %s", o.item.Name())))
				moved = true
				p.Move(s, dx, dy)
			}
		}

		dx = 0
		dy = 0
	}

	if p.hp < 1 {
		return gameOver
	}
	return none
}
