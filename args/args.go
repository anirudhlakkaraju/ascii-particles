// args.go
package args

import (
	"flag"
	"fmt"
	"log"
	"particles/particles"
	"strings"
)

// ParseArgs handles flag parsing and returns the appropriate action.
func ParseArgs() (*string, *string, *bool) {
	argType := flag.String("effect", particles.DefaultEffect.Name, "Specify the type of particle effect (--list to see all effects)")
	argAdd := flag.String("add", "", "Add new particle effect assets in the format: {name:assets} (e.g., 'steam:.:{}')")
	argList := flag.Bool("list", false, "List all available particle effects")

	flag.Parse()

	return argType, argAdd, argList
}

// HandleList displays all available particle effects
func HandleList() {
	fmt.Println("Available Particle Effects:")
	for name, effect := range particles.Effects {
		isDefault := ""
		if name == particles.DefaultEffect.Name {
			isDefault = " (default)"
		}
		fmt.Printf("  %s: %s%s\n", name, effect.Asset, isDefault)
	}
}

// HandleAdd adds a new particle effect
func HandleAdd(argAdd string) {
	parts := strings.SplitN(argAdd, ":", 2)
	if len(parts) != 2 {
		log.Fatalf("Invalid format for --add. Use: name:asset (e.g., 'rain:.,,:')")
	}
	name, asset := parts[0], parts[1]
	if _, exists := particles.Effects[name]; exists {
		log.Fatalf("Effect '%s' already exists.", name)
	}

	particles.Effects[name] = particles.NewEffect(name, asset)
	particles.SaveEffects() // Persist the changes to the file
	fmt.Printf("Added new particle effect '%s'.\n", name)
	fmt.Printf("Access through argument '--type %s'\n", name)
}
