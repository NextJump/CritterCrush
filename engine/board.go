package engine

import ()

type Board struct {
	Height       int
	Width        int
	Players      []*Player
	Enemies      []*enemy
	Environments []*environment
}

func (b *Board) setup() {
	b.initBoard(10, 10)
	b.initPlayers()
	b.initEnvironment()
	b.initEnemies()
}

func (b *Board) initBoard(height int, width int) {
	b.Height = height
	b.Width = width
}

func (b *Board) initPlayers() {
	b.Players = make([]*Player, 2)
	b.Players[0] = &Player{1, 0, 0, 0, 5, "http://localhost:%d/action?playernum=%d"}
	b.Players[1] = &Player{2, 0, 0, 0, 5, "http://localhost:%d/action?playernum=%d"}
}

func (b *Board) initEnemies() {
	b.Enemies = make([]*enemy, 8)

	b.Enemies[0] = &enemy{1, 2, 3, 5, 5, false}
	b.Enemies[1] = &enemy{1, 8, 6, 5, 5, false}

	b.Enemies[2] = &enemy{2, 4, 1, 10, 10, false}
	b.Enemies[3] = &enemy{2, 5, 9, 10, 10, false}

	b.Enemies[4] = &enemy{3, 1, 6, 10, 10, false}
	b.Enemies[5] = &enemy{3, 3, 6, 10, 10, false}

	b.Enemies[6] = &enemy{3, 8, 1, 10, 10, false}
	b.Enemies[7] = &enemy{3, 9, 6, 10, 10, false}
}

func (b *Board) initEnvironment() {
	b.Environments = make([]*environment, 4)

	b.Environments[0] = &environment{0, 8, 4}
	b.Environments[1] = &environment{0, 9, 4}
	b.Environments[2] = &environment{0, 3, 4}
	b.Environments[3] = &environment{0, 4, 4}
}

func (b *Board) IsEnemyAt(x int, y int) bool {
	for _, e := range b.Enemies {
		if e.X == x && e.Y == y {
			return true
		}
	}
	return false
}

func (b *Board) GetPlayer(player int) *Player {
	p := b.Players[player]
	return p
}
