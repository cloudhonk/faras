package game

func rotatePlayers(players []*Juwadey, currentPlayeridx int) []*Juwadey {
	n := len(players)
	rotated := make([]*Juwadey, n)
	for i := range n {
		// Calculate the position based on the current player's perspective
		rotated[i] = players[(currentPlayeridx+i)%n]
	}
	return rotated
}
