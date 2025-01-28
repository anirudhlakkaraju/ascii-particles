// Package particles provides a particle system for ASCII art
package particles

import (
	"math"
	"math/rand"
	"time"
)

// Coffee is a particle system for steamin hot coffee
type Coffee struct {
	ParticleSystem
}

var startTime = time.Now().UnixMilli()

// ascii represents particle density with steam effect
func ascii(row, col int, counts [][]int, asset *ParticleEffect) string {
	count := counts[row][col]
	surroundCount := countParticlesAround(row, col, counts)

	var value interface{}
	switch {
	case count < 1:
		value = asset.None
	case surroundCount > 5:
		// alternate between high and max asset
		direction := row + int(((time.Now().UnixMilli()-startTime)/2000)%2)
		if direction%2 == 0 {
			value = asset.High
		} else {
			value = asset.Max
		}
	case surroundCount > 4:
		value = asset.Medium
	default:
		value = asset.Low
	}

	// handle particle asset types
	switch v := value.(type) {
	case rune:
		return string(v)
	case string:
		return v
	default:
		return string(asset.None)
	}
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

// dirs represents the 8 directions around a particle
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

// countParticlesAround counts the number of particles around a given particle
func countParticlesAround(row, col int, counts [][]int) int {
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

// NewCoffee creates a new coffee system
func NewCoffee(params ParticleParams) Coffee {

	// force odd system width to help with normal distribution
	if params.X%2 == 0 {
		params.X++
	}

	params.reset = reset
	params.nextPosition = nextPosition
	params.Ascii = ascii

	return Coffee{ParticleSystem: NewParticleSystem(params)}

}
