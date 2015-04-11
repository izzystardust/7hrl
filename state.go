package main

import (
	"errors"

	"github.com/nsf/termbox-go"
)

type state struct {
	ms       [10]gameMap
	m        *gameMap
	entities []entity
	items    []object
	e        error
	messages []statusMessage
}

func (s state) Draw() {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			termbox.SetCell(x, y, s.m.At(x, y).Draw(), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	l := s.Player().level
	for _, e := range s.entities {
		if e.CurrentLevel() != l {
			continue
		}
		x, y := e.CurPos()
		termbox.SetCell(x, y, e.Draw(), termbox.ColorDefault|termbox.AttrBold, termbox.ColorDefault)
	}

	for _, o := range s.items {
		if o.level != l {
			continue
		}
		termbox.SetCell(o.x, o.y, o.item.Draw(), termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Flush()
}

func (s state) Update() state {
	remaining := make([]entity, 0, len(s.entities))
	for _, e := range s.entities {
		result := e.Update(&s)
		if result == gameOver {
			s.e = errors.New("Game over")
			s.messages = append(s.messages, "Game over")
			break
		}
		if result != shouldDie {
			remaining = append(remaining, e)
		} else {
			s.messages = append(s.messages, statusMessage(e.Name()+" died"))
		}
	}
	s.entities = remaining

	return s
}

func (s state) Walkable(x, y int) bool {
	if !s.m.At(x, y).Walkable() {
		return false
	}
	for _, e := range s.entities {
		ex, ey := e.CurPos()
		if ex == x && ey == y {
			return false
		}
	}

	for _, o := range s.items {
		if o.x == x && o.y == y {
			return false
		}
	}
	return true
}

func (s state) AttackableAt(x, y int) attacker {
	for _, e := range s.entities {
		ex, ey := e.CurPos()
		if ex == x && ey == y {
			return e
		}
	}
	return nil
}

// removes the item!
func (s *state) ObjectAt(x, y int) *object {
	for i, o := range s.items {
		if x == o.x && y == o.y {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return &o
		}
	}
	return nil
}

func (s *state) generateEnemies() {
	for i := 0; i < 40; i++ {
		e := newOyster()
		tx, ty := randomPoint(width, height)
		if s.Walkable(tx, ty) {
			e.x = tx
			e.y = ty
			s.entities = append(s.entities, e)
		}
	}
}

func (s *state) generateObjects() {
	for i := 0; i < 10; i++ {
		tx, ty := randomPoint(width, height)
		debug("here")
		if s.Walkable(tx, ty) {
			debug("and here")
			s.items = append(s.items, newObject(tx, ty))
		}
	}
	tx, ty := randomPoint(width, height)
	for !s.Walkable(tx, ty) {
		tx, ty = randomPoint(width, height)
	}
	s.items = append(s.items, object{tx, ty, newSword(), 0})
}

func (s *state) Player() *player {
	return s.entities[0].(*player)
}
