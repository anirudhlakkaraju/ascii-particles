package particles

import (
	"encoding/json"
	"log"
	"os"
	"unicode/utf8"
)

const EffectsFile = "particles/effects.json"

type Effect struct {
	Name  string `json:"name"`
	Asset string `json:"asset"`
}

func NewEffect(name, asset string) *Effect {
	return &Effect{Name: name, Asset: asset}
}

var NoParticle = ' '
var DefaultEffect = NewEffect("type1", ".:{}")
var Effects map[string]*Effect

// ParticleEffect is a particle effect at various densities
type ParticleEffect struct {
	Name   string
	None   rune
	Low    rune
	Medium rune
	High   rune
	Max    rune
}

// NewParticleEffect creates a new particle effect using the provided name and assets
func NewParticleEffect(effect Effect) *ParticleEffect {

	// check if number of runes are 4
	if utf8.RuneCountInString(effect.Asset) != 4 {
		log.Fatalf("Expected number of particle assets: 4. Got: %v", utf8.RuneCountInString(effect.Asset))
	}

	assetRunes := []rune(effect.Asset)

	return &ParticleEffect{
		Name:   effect.Name,
		None:   NoParticle,
		Low:    assetRunes[0],
		Medium: assetRunes[1],
		High:   assetRunes[2],
		Max:    assetRunes[3],
	}
}

// LoadEffects loads particle effects from a JSON file
func LoadEffects() {
	data, err := os.ReadFile(EffectsFile)
	if os.IsNotExist(err) {
		// Create the default file if it doesn't exist
		Effects[DefaultEffect.Name] = DefaultEffect
		SaveEffects()
		return
	} else if err != nil {
		log.Fatalf("Error reading effects file: %v", err)
	}

	if err := json.Unmarshal(data, &Effects); err != nil {
		log.Fatalf("Error parsing effects file: %v", err)
	}
}

// SaveEffects writes the current effects map to the JSON file
func SaveEffects() {
	data, err := json.MarshalIndent(Effects, "", "  ")
	if err != nil {
		log.Fatalf("Error serializing effects: %v", err)
	}

	if err := os.WriteFile(EffectsFile, data, 0644); err != nil {
		log.Fatalf("Error writing effects file: %v", err)
	}
}
