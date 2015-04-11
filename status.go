package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type statusMessage string

func (s statusMessage) display() {
	for x := len(s); x < width; x++ {
		termbox.SetCell(x, 26, ' ', termbox.ColorBlack, termbox.ColorWhite)
	}
	for x, c := range s {
		termbox.SetCell(x, 26, c, termbox.ColorBlack, termbox.ColorWhite)
	}
	termbox.Flush()

	termbox.PollEvent()
	for x := 0; x < width; x++ {
		termbox.SetCell(x, 26, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Flush()
}

func displayPlayerInfo(p *player) {
	str := fmt.Sprintf("hp: %3d/%d  equipped: %s", p.hp, p.maxhp, p.weapon.name)
	for x := len(str); x < width; x++ {
		termbox.SetCell(x, 25, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	for x, c := range str {
		termbox.SetCell(x, 25, c, termbox.ColorDefault|termbox.AttrBold, termbox.ColorDefault)
	}
	termbox.Flush()
}
