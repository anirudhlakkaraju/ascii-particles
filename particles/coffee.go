package particles

import (
	"math"
	"math/rand"
)

type Coffee struct {
	ParticleSystem
}

func ascii(row, col int, counts [][]int) rune {
	return '}'
}

// reset particle's lifetime, speed and position
func reset(p *Particle, params *ParticleParams) {
	p.Lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	p.Speed = math.Floor(params.MaxSpeed * rand.Float64())

	// translate X coordinate on generation
	maxX := math.Floor(float64(params.X) / 2)
	p.X = math.Max(-maxX, math.Min(rand.NormFloat64(), maxX))
	p.Y = 0
}

// nextPosition updates the particle's vertical position
func nextPosition(p *Particle, deltaMS int64) {
	p.Lifetime -= deltaMS
	if p.Lifetime < 0 {
		return
	}

	// rise particle straight up for time elapsed (in seconds)
	p.Y += p.Speed * (float64(deltaMS) / 1000.0)
}

// NewCoffee creates a new coffee system
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
