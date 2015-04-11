package main

type entity interface {
	updater
	drawer
	attacker
	namer
	leveler
}

type leveler interface {
	CurrentLevel() int
	SendToLevel(int)
}

type updater interface {
	Update(s *state) updateResult
	CurPos() (x, y int)
}

type namer interface {
	Name() string
}

type attackType byte

const (
	attackNormal attackType = 1 << iota
	attackMagic
)

type attacker interface {
	// returns the amount of damage dealt and the type - the attacked may counterattack
	Attack() (int, attackType)
	Defend(amount int, ty attackType) // takes the amount of damage dealt and type, modifies itself thusly.
}

func isPlayer(a attacker) bool {
	switch a.(type) {
	case *player:
		return true
	default:
		return false
	}
}

type updateResult byte

const (
	gameOver updateResult = iota
	shouldDie
	none
)
