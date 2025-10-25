# Pokedex CLI

An interactive terminal application for exploring the Pokémon world using the public PokeAPI. Browse location areas, explore the Pokémon found there, try to catch them, and inspect your growing Pokédex — all from a simple REPL prompt.

## Features

- Browse location areas with forward/back pagination (`map`, `mapb`).
- Explore an area to list encountered Pokémon (`explore <area>`).
- Attempt to catch Pokémon with a chance based on their base experience (`catch <name>`).
- Inspect stats for any Pokémon you’ve caught (`inspect <name>`).
- View your local Pokédex (`pokedex`).
- Built‑in HTTP response caching (5 minutes) to reduce API calls and speed up repeated requests.

## Quick Start

You can either run the prebuilt binary if present or build from source.

### Run (prebuilt)

```
./pokedexcli
```

### Build from source

Requirements:

- Go toolchain installed (any recent Go version should work; see `go.mod` for the declared version).

Build and run:

```
go build -o pokedexcli
./pokedexcli
```

Or run directly without building:

```
go run .
```

## Usage

When launched, you’ll see a REPL prompt:

```
Pokedex >
```

Type a command and press Enter. Commands are case‑insensitive and accept at most one argument. Use `help` any time to see available commands.

### Command Reference

- `help` — Display command help.
- `map` — List the next page of location areas.
- `mapb` — List the previous page of location areas.
- `explore <area>` — Show Pokémon that can be encountered in the specified location area.
  - Tip: Use names exactly as printed by `map`.
- `catch <pokemon>` — Attempt to catch the specified Pokémon.
- `inspect <pokemon>` — Show stats for a caught Pokémon.
- `pokedex` — List all Pokémon you have caught.
- `exit` — Quit the application.

### Examples

```
Pokedex > help
Welcome to the Pokedex!
Usage:

map: Iterate over locations
mapb: Iterate over locations
explore: Explore selected area
catch: Catch a pokemon
inspect: Inspect a pokemon
pokedex: List all catched pokemon names
help: Displays a help message
exit: Exit the Pokedex

Pokedex > map
canalave-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area
Found Pokemon:
 - starly
 - bidoof
 ...

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
 -speed: 90
 -special-attack: 50
 -special-defense: 40
 -attack: 55
 -defense: 40
 -hp: 35

Pokedex > pokedex
Your Pokedex:
 - pikachu
```

Note: Exact area and stat values come from PokeAPI and may differ based on the Pokémon or area you query.

## How It Works

- REPL and commands live in the root (`main.go`, `repl.go`, `command_*.go`).
- PokeAPI client and types are in `internal/pokeapi`, including simple 5‑minute response caching via `internal/pokecache`.
- Network calls are made directly to `https://pokeapi.co/api/v2/...` with JSON unmarshalling into local structs.

## Development

- Run tests:

```
go test ./...
```

- Key directories:
  - `internal/pokeapi` — HTTP client and data models.
  - `internal/pokecache` — in‑memory cache with a background reaper goroutine.

## Troubleshooting

- “Unknown command”: Run `help` to see valid commands.
- Not finding an area: Use the exact name from `map` output (lowercase, hyphenated per PokeAPI).
- Catch keeps failing: Chance is influenced by the Pokémon’s base experience; try a different Pokémon.

## Acknowledgments

- Data provided by PokeAPI (https://pokeapi.co/). No API key required.

