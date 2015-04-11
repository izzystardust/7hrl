package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	width  = 80
	height = 24
)

type drawer interface {
	Draw() rune
}

func newGame() state {

	s := state{
		entities: []entity{newPlayer(width/2, height/2)},
	}
	s.ms = newDungeon(80, 24)
	s.m = &s.ms[0]
	s.generateEnemies()
	s.generateObjects()
	return s
}

func debug(rest ...interface{}) {
	fmt.Fprintln(os.Stderr, rest...)
}

func main() {
	termbox.Init()
	defer termbox.Close()
	rand.Seed(time.Now().Unix())
	s := newGame()
	s.Draw()
	for {
		displayPlayerInfo(s.entities[0].(*player)) // the exact point I said 'fuck it' to type safety
		s = s.Update()
		s.Draw()
		for _, m := range s.messages {
			m.display()
		}

		s.messages = nil
		if s.e != nil {
			break
		}
		displayPlayerInfo(s.Player())
	}
}
