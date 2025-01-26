package particles

type Coffee struct {
	ParticleSystem
}

func ascii(row, col int, counts [][]int) rune {
	return '}'
}

func NewCoffee(width, height int) Coffee {
	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				MaxLife:       7,
				MaxSpeed:      0.5,
				ParticleCount: 100,

				X: width,
				Y: height,

				reset:        reset,
				nextPosition: nextPosition,
				ascii:        ascii,
			}),
	}

}
