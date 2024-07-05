# Snakes and Ladders Game

Welcome to a simplified version of Snakes and Ladders implemented in Go. <br>
In this version, there are no snakes or ladders (yet), but you can enjoy the game of rolling the dice and moving across the board. <br>
I'm working on this project to practice Go fundamentals I've learned from the book "For the love of Go" by John Arundel. Highly recommend it.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Compiling and Creating an Executable](#compiling-and-creating-an-executable)
- [Project Structure](#project-structure)
- [How to Play](#how-to-play)
- [Running Tests](#running-tests)

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/snakes-and-ladders.git
    cd snakes-and-ladders
    ```

2. **Install the dependencies:**

    ```bash
    go get github.com/eiannone/keyboard
    ```

## Usage

1. **Run the game:**

    ```bash
    go run main.go
    ```

2. **Press Enter to roll the dice for the current player.**

## Compiling and Creating an Executable

1. **Compile the game:**

    ```bash
    go build -o snakes-and-ladders main.go
    ```

2. **Run the executable:**

    - On Windows:

        ```bash
        snakes-and-ladders.exe
        ```

    - On macOS/Linux:

        ```bash
        ./snakes-and-ladders
        ```

## Project Structure

- `main.go`: The main entry point for the game.
- `board/board.go`: Contains the `Board` struct and methods for updating and printing the board.
- `board/board_test.go`: Contains tests for the `Board` struct and its methods.

## How to Play

1. **Start the game:**
    - The game welcomes you to Snakes and Ladders.
    - Currently, there are no snakes or ladders implemented.

2. **Player turns:**
    - Player One and Player Two take turns rolling the dice.
    - Press Enter to roll the dice for the current player.

3. **Rolling the dice:**
    - A random number between 1 and 6 is generated to simulate the dice roll.
    - The player's position is updated based on the dice roll.

4. **Winning the game:**
    - The first player to reach position 100 wins the game.

## Running Tests

1. **Run the tests:**

    ```bash
    go test ./...
    ```
