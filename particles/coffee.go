// Package particles provides a particle system for ASCII art
package particles

import (
	"math"
	"math/rand"
)

// Coffee is a particle system for steamin hot coffee
type Coffee struct {
	ParticleSystem
}

// ascii returns the ASCII representation of the particle
func ascii(row, col int, counts [][]int) string {
	count := counts[row][col]
	if count < 3 {
		return " "
	}
	if count < 6 {
		return "."
	}
	if count < 9 {
		return ":"
	}
	if count < 12 {
		return "{"
	}
	return "}"
}

// reset particle's lifetime, speed and position
func reset(p *Particle, params *ParticleParams) {
	p.Lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	p.Speed = params.MaxSpeed * rand.Float64()

	// translate X coordinate on generation
	maxX := math.Floor(float64(params.X) / 2)
	x := math.Max(-maxX, math.Min(rand.NormFloat64()*params.XStDeviation, maxX))
	p.X = x + maxX
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

	// force odd system width to help with normal distribution
	if width%2 == 0 {
		width++
	}

	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				MaxLife:       7,
				MaxSpeed:      0.5,
				ParticleCount: 100,

				XStDeviation: scale,
				X:            width,
				Y:            height,

				reset:        reset,
				nextPosition: nextPosition,
				ascii:        ascii,
			}),
	}

}
