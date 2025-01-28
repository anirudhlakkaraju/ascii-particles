# ASCII Particle System for Coffee

This project simulates an ASCII particle system for a hot cup of coffee.

The particle densities are generated as a normal distribution to create a realistic effect. Project inspiration credit goes to the Primeagen!


![Coffee Particle Effect](https://github.com/user-attachments/assets/ce0603f5-b3b2-47bb-a488-f5fac722d922)

## Features:
- **Effect Selection**: Choose from a variety of pre-defined particle effects, accessible by specifying the `--list` argument.
- **Add Custom Effects**: You can add new effects through the --add option. This will update `particles/effects.json`

## Installation:
1. Clone or download the repository.
2. Make sure you have Go installed.

## Usage:

Run the project with the `--effect` flag to choose between effects (default is `'type1'`).

```bash
go run main.go --effect type1
```

Use `--help` to see all the args provided.

---

Feel free to experiment with the particle system parameters in `main.go`.

### Enjoy your coffee! ☕
