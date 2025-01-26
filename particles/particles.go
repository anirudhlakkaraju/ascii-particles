package particles

import (
	"math"
	"time"
)

// Particle is a single particle in the system
type Particle struct {
	Lifetime int64
	Speed    float64

	// positions
	X float64
	Y float64
}

// ParticleParams is the parameters for the particle system
type ParticleParams struct {
	MaxLife  float64
	MaxSpeed float64

	ParticleCount int64

	// dimensions of system
	X int
	Y int

	nextPosition NextPositionFunc
	ascii        Ascii
	reset        Reset
}

type NextPositionFunc func(particle *Particle, deltaMS int64)
type Ascii func(row, col int, count [][]int) rune
type Reset func(particle *Particle, params *ParticleParams)

// ParticleSystem is the system of particles
type ParticleSystem struct {
	ParticleParams
	particles []*Particle

	// last time of update
	lastTime int64
}

// NewParticleSystem creates a new particle system
func NewParticleSystem(params ParticleParams) ParticleSystem {
	return ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
	}
}

// Start starts the particle system
func (ps *ParticleSystem) Start() {
	for _, p := range ps.particles {
		ps.reset(p, &ps.ParticleParams)
	}
}

// Update updates the particle system
func (ps *ParticleSystem) Update() {
	now := time.Now().UnixMilli()
	delta := now - ps.lastTime
	ps.lastTime = now

	for _, p := range ps.particles {
		ps.nextPosition(p, delta)

		if p.Y >= float64(ps.Y) || p.X >= float64(ps.X) {
			ps.reset(p, &ps.ParticleParams)
		}
	}
}

// Display returns the ascii representation of the particle system
func (ps *ParticleSystem) Display() [][]rune {
	counts := make([][]int, 0)

	// Initialize counts to size of particle system dimensions
	for row := 0; row < ps.Y; row++ {
		count := make([]int, 0)
		for col := 0; col < ps.X; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}

	for _, p := range ps.particles {
		row := int(math.Floor(p.Y))
		col := int(math.Floor(p.X))

		counts[row][col]++
	}

	out := make([][]rune, 0)
	for r, row := range counts {
		outRow := make([]rune, 0)
		for c := range row {
			outRow = append(outRow, ps.ascii(r, c, counts))
		}
		out = append(out, outRow)
	}

	return out
}
