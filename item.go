package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type item interface {
	user
	namer
	drawer
}

// an object is an item that is on the map
type object struct {
	x, y int
	item
	level int
}

func newObject(x, y int) object {
	p := potion{func(s *state, p *player) {
		p.hp += 5
		if p.hp > p.maxhp {
			p.hp = p.maxhp
		}
	}, "health potion"}
	return object{x, y, p, 0}
}

type user interface {
	Use(s *state)
}

func pickItem(s *state) item {
	writeStringAt(1, 1, "Choose item: (0 cancels, invalid segfaults)")
	for i, o := range s.Player().items {
		writeStringAt(1, i+2, fmt.Sprintf("%d: %s", i+1, o.Name()))
	}
	termbox.Flush()
	e := termbox.PollEvent()
	for e.Ch < '0' && e.Ch > '9' {
		e = termbox.PollEvent()
	}

	if e.Ch == '0' {
		return nil
	}
	p := s.Player()
	i := int(e.Ch) - '1'
	it := p.items[i]
	p.items = append(p.items[:i], p.items[i+1:]...)
	return it
}

func writeStringAt(x, y int, s string) {
	for xi, c := range s {
		termbox.SetCell(x+xi, y, c, termbox.ColorBlack, termbox.ColorWhite)
	}
}
