package main

import "math/rand"

type mapcell int

const (
	cellearth mapcell = iota
	cellempty
	cellwall
	cellstaird
	cellstairu
)

func (a mapcell) Draw() rune {
	switch a {
	case cellempty:
		return '.'
	case cellearth:
		return '~'
	case cellwall:
		return '#'
	case cellstaird:
		return '↳'
	case cellstairu:
		return '↑'
	}
	return '!'
}

func (a mapcell) Walkable() bool {
	good := false
	if a == cellempty || a == cellstaird || a == cellstairu {
		good = true
	}
	return good
}

type gameMap struct {
	contents           []mapcell
	w, h               int
	enterUpX, enterUpY int //enterUp is entering from beneath
	enterDwX, enterDwY int
}

type room struct {
	x, y, w, h int
}

func newDungeon(w, h int) [10]gameMap {
	var maps [10]gameMap
	for i := 0; i < 10; i++ {
		maps[i] = newMap(w, h)
	}
	maps[0].addStairDown()
	for i := 1; i < 9; i++ {
		maps[i].addStairDown()
		maps[i].addStairUp()
	}
	maps[9].addStairUp()
	return maps
}

func (g *gameMap) onEnter(s *state, x, y int) {
	c := g.At(x, y)
	p := s.Player()
	if c == cellstairu {
		p.SendToLevel(p.CurrentLevel() - 1)
		p.x = s.ms[p.CurrentLevel()].enterUpX
		p.y = s.ms[p.CurrentLevel()].enterUpY
	} else if c == cellstaird {
		p.SendToLevel(p.CurrentLevel() + 1)
		p.x = s.ms[p.CurrentLevel()].enterDwX
		p.y = s.ms[p.CurrentLevel()].enterDwY
	}
}

func newMap(w, h int) gameMap {
	g := gameMap{w: w, h: h, contents: make([]mapcell, h*w)}

	rooms := []room{}

	// mapgen alg from http://www.roguebasin.com/index.php?title=Dungeon-Building_Algorithm
	// but adapted for only having 6 hours
	// step 1: fill with solid earth - done by zero value
	// step 2: dig room out in center of map
	initH := rand.Intn(5) + 4
	initW := rand.Intn(5) + 4
	g.digRoom(w/2-initW/2, h/2-initH/2, initW, initH)
	rooms = append(rooms, room{w/2 - initW/2, h/2 - initH/2, initW, initH})
	for i := 0; i < 50; i++ {
		// step 3: pick a wall of any room
		tx, ty := randomPoint(w, h)
		for !g.adjacentClear(tx, ty) {
			tx, ty = randomPoint(w, h)
		}
		g.addSpace(tx, ty)
	}
	for x := 0; x < w; x++ {
		g.Set(x, 0, cellwall)
		g.Set(x, h-1, cellwall)
	}
	for y := 0; y < h; y++ {
		g.Set(0, y, cellwall)
		g.Set(w-1, y, cellwall)
	}
	return g
}

func (g *gameMap) addStairUp() {
	tx, ty := randomPoint(g.w, g.h)
	for !g.adjacentClear(tx, ty) {
		tx, ty = randomPoint(g.w, g.h)
	}
	g.Set(tx, ty, cellstairu)
	g.enterDwX = tx
	g.enterDwY = ty
}

func (g *gameMap) addStairDown() {
	tx, ty := randomPoint(g.w, g.h)
	for !g.adjacentClear(tx, ty) {
		tx, ty = randomPoint(g.w, g.h)
	}
	g.Set(tx, ty, cellstaird)
	g.enterUpX = tx
	g.enterUpY = ty

}

func randomPoint(w, h int) (int, int) {
	return rand.Intn(w-2) + 1, rand.Intn(h-2) + 1
}

func (g gameMap) At(x, y int) mapcell {
	return g.contents[x+width*y]
}

func (g *gameMap) Set(x, y int, to mapcell) {
	if x+width*y >= len(g.contents) || x < 0 || y < 0 {
		return
	}
	g.contents[x+width*y] = to
}

func (g *gameMap) digRoom(x, y, w, h int) {
	for xi := 0; xi < w; xi++ {
		for yi := 0; yi < h; yi++ {
			g.Set(x+xi, y+yi, cellempty)
		}
	}

}

func (g gameMap) adjacentClear(x, y int) bool {
	e := cellempty
	return g.At(x+1, y) == e || g.At(x-1, y) == e || g.At(x, y+1) == e || g.At(x, y-1) == e
}

func (g gameMap) addSpace(x, y int) {
	isHall := rand.Intn(4) != 0
	w, h := 1, 1
	if isHall {
		if rand.Intn(2) == 0 {
			w = rand.Intn(g.w/4) + 2
		} else {
			h = rand.Intn(g.h/4) + 2
		}
	} else {
		w = rand.Intn(g.w/6) + 2
		h = rand.Intn(g.w/6) + 2
	}

	l := rand.Intn(2) == 0
	u := rand.Intn(2) == 0

	if l {
		x -= w
	}
	if u {
		y -= h
	}
	g.digRoom(x, y, w, h)
}
