# ASCII Particle System for Coffee

This project simulates an ASCII particle system for a hot cup of coffee. The particle effect can be toggled between two options: **steam** and **fire**.

The particle densities are generated as a normal distribution to create a realistic effect. Project inspiration credit goes to the Primeagen!


https://github.com/user-attachments/assets/ce0603f5-b3b2-47bb-a488-f5fac722d922


## Features:
- **Steam Effect**: Simulates steam rising from your hot coffee.
- **Fire Effect**: Simulates a more intense, fiery particle effect.

## Installation:
1. Clone or download the repository.
2. Make sure you have Go installed.

## Usage:
Run the project with the `--effect` flag to choose between the two effects. The default is "steam".

```bash
go run main.go --effect steam
```
To run the fire effect:
```bash
go run main.go --effect fire
```

Feel free to experiment with the particle system parameters in `main.go`.

### Enjoy your coffee! ☕
