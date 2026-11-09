# gopokr

A terminal-based five-card poker game written in Go. You play against a dealer that decides whether to redraw based on the chosen difficulty.

## Features

- Five-card draw gameplay in the terminal
- Dealer AI with multiple difficulty settings
- Score tracking for wins, losses, and ties
- Hand evaluation for standard poker categories
- Simple text-based interface

## How to play

1. Start the game:

   ```bash
   go run .
   ```

2. Choose a difficulty:
   - `e` for easy
   - `m` for medium
   - `h` for hard
   - `fuckyou` for the extra-hard mode (if you want to test the dealer's nastiest setting)

3. On each round, your hand is shown with the dealer's hand partially hidden.
4. Enter one of the following commands:
   - `r` or `R` to redraw cards
   - `e` or `E` to end the turn
   - `f` or `F` to fold
   - `q` or `Q` to quit

5. If you redraw, enter the card positions to replace, using numbers `1` through `5`.

## Difficulty modes

- Easy: the dealer keeps strong hands and discards weak high-card hands.
- Medium: balanced draw strategy.
- Hard: stronger strategic decisions and better discard logic.
- fuckYouMode: ruthless behavior with a more aggressive redraw strategy.

## Winning hands

The game evaluates standard poker hands in this order, from strongest to weakest:

1. Straight flush
2. Four of a kind
3. Full house
4. Flush
5. Straight
6. Three of a kind
7. Two pair
8. One pair
9. High card

## Development

Run the tests:

```bash
go test ./...
```

## Notes

This project is a lightweight command-line game and is intentionally designed around a fast, text-based experience rather than a polished GUI.
