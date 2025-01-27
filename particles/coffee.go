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
	percent := float64(deltaMS) / 1000.0
	p.Y += p.Speed * percent
}

var dirs = [][]int{
	// bottom row
	{-1, -1},
	{-1, 0},
	{-1, 1},

	// middle row
	{0, -1},
	{0, 1},

	// top row
	{1, 0},
	{1, 1},
	{1, -1},
}

func countParticles(row, col int, counts [][]int) int {
	count := 0
	for _, dir := range dirs {
		r := row + dir[0]
		c := col + dir[1]

		if r < 0 || r >= len(counts) || c < 0 || c >= len(counts[0]) {
			continue
		}
		count += counts[row+dir[0]][col+dir[1]]
	}
	return count
}

func normalize(row, col int, counts [][]int) {

	if countParticles(row, col, counts) > 4 {
		counts[row][col] = 0
	}

}

// NewCoffee creates a new coffee system
func NewCoffee(width, height int, scale float64) Coffee {

	// force odd system width to help with normal distribution
	if width%2 == 0 {
		width++
	}

	startTime := time.Now().UnixMilli()
	ascii := func(row, col int, counts [][]int) string {
		count := counts[row][col]
		if count < 1 {
			return " "
		}

		direction := row +
			int(((time.Now().UnixMilli()-startTime)/2000)%2)

		if countParticles(row, col, counts) > 3 {
			if direction%2 == 0 {
				return "{"
			}
			return "}"
		}
		return "."
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
