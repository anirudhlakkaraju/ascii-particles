package particles

type Coffee struct {
	ParticleSystem
}

func ascii(row, col int, counts [][]int) rune {
	return '}'
}

func NewCoffee(width, height int) Coffee {
	reset := func(row, col int, counts [][]int) {}

	ascii := func(row, col int, counts [][]int) rune {
	}
}
