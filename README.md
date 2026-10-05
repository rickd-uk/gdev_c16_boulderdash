# Boulder Dash C16 Remake

A remake of the Commodore 16 version of **Boulder Dash**, written in Go using [Ebitengine](https://ebitengine.org/).

This project is being built from scratch as a learning exercise, with a focus on understanding how a classic tile-based game works internally.

## Goals

- Recreate the core gameplay of Boulder Dash
- Learn tile-based game programming in Go
- Build a deterministic cave simulation
- Study the behaviour of the original C16 version
- Reproduce classic mechanics such as:
  - digging through dirt
  - falling boulders
  - rolling boulders
  - diamonds
  - enemies
  - explosions
  - exits
  - timers
- Eventually recreate the visual and audio style of the C16 version

## Technology

- Go
- Ebitengine

## Current Progress

- [x] Basic Ebitengine window
- [x] Tile-based cave grid
- [x] Dirt, walls, and empty tiles
- [ ] Player movement
- [ ] Digging
- [ ] Boulder physics
- [ ] Diamonds
- [ ] Enemies
- [ ] Cave completion
- [ ] Original-style graphics
- [ ] Sound

## Development

Run the game with:

```bash
go run .
```

## Project Status

Early development.

The project is intentionally being built step by step rather than starting from a completed engine. The aim is to understand each part of the game simulation as it is implemented.

## Disclaimer

Boulder Dash is an existing commercial game and trademark of its respective rights holders.

This project is an educational remake and is not intended for commercial distribution.
