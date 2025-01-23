# Pokedex CLI

Pokedex CLI is a command-line interface application that allows users to interact with the PokeAPI to catch, inspect, and explore Pokémon and their locations.

## Features

- **Catch Pokémon**: Attempt to catch a Pokémon by name.
- **Inspect Pokémon**: View detailed information about a caught Pokémon.
- **Explore Locations**: Explore locations to see which Pokémon can be encountered there.
- **List Locations**: View a paginated list of locations.
- **Help**: Display a help message with available commands.
- **Exit**: Exit the Pokedex CLI.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/mdelgadonyc/pokedexcli.git
    cd pokedexcli
    ```

2. Build the project:
    ```sh
    go build -o pokedexcli
    ```

## Usage

Run the application:
```sh
./pokedexcli
```

Once the application is running, you can use the following commands:

- `catch <pokemon name>`: Attempt to catch a Pokémon by name.
- `inspect <pokemon name>`: View detailed information about a caught Pokémon.
- `explore <location name>`: Explore a location to see which Pokémon can be encountered there.
- `map`: Get the next page of locations.
- `mapb`: Get the previous page of locations.
- `pokedex`: See all the Pokémon you've caught.
- `help`: Display a help message with available commands.
- `exit`: Exit the Pokedex CLI.

## Example

```sh
Pokedex > catch pikachu
Throwing a Pokeball at Pikachu...
Pikachu was caught!
You may now inspect it with the inspect command.

Pokedex > inspect pikachu
Name: Pikachu
Height: 4
Weight: 60
Stats:
  speed: 90
  special-defense: 50
  special-attack: 50
  defense: 40
  attack: 55
  hp: 35
Types:
  - electric

Pokedex > explore kanto
Exploring kanto...
Pokemon encounters:
 - bulbasaur
 - charmander
 - squirtle

Pokedex > exit
Closing the Pokedex... Goodbye!
```

## Project Structure

- `main.go`: Entry point of the application.
- `repl.go`: Handles the REPL (Read-Eval-Print Loop) for user input.
- `internal/pokeapi`: Contains the PokeAPI client and related functions.
- `internal/pokecache`: Contains the caching mechanism for API responses.
- `command_*.go`: Contains the implementations of various commands.

## Testing

Run the tests:
```sh
go test ./...
```

## License

This project is licensed under the MIT License.