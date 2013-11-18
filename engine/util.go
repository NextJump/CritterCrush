package engine

func isWall(board *Board, x int, y int) bool {
	if x < 0 || y < 0 {
		return true
	}
	if x > (board.Width - 1) {
		return true
	}
	if y > (board.Height - 1) {
		return true
	}

	for _, env := range board.Environments {
		if env.Id != 0 {
			continue
		}
		if env.X == x && env.Y == y {
			return true
		}
	}
	return false
}

func isEnemy(board *Board, x int, y int) bool {
	for _, e := range board.Enemies {
		if e.IsCrushed {
			continue
		}
		if e.X == x && e.Y == y {
			return true
		}
	}
	return false
}

func getEnemy(board *Board, x int, y int) *enemy {
	for _, e := range board.Enemies {
		if e.X == x && e.Y == y {
			return e
		}
	}
	return nil
}
