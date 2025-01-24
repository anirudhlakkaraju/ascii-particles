package particles

type Particle struct {
	lifetime int64
	speed    float64

	x float64
	y float64
}

type ParticleParams struct {
	MaxLife  float64
	MaxSpeed float64

	ParticleCount int64
	
	X int
	Y int
}

type UpdateFunc = func(particle *Particle, deltaMS int64)

type ParticleSystem struct {
	ParticleParams

	lastTime int64
	place UpdateFunc
}

func NewParticleSystem(params ParticleParams, updateFunc UpdateFunc) ParticleSystem {
	particleSystem := ParticleSystem{}
	return particleSystem
}
