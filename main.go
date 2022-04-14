package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Sam36502/go-board/v2"
)

const (
	ARENA_WIDTH  = 25
	ARENA_HEIGHT = 25
	NUM_TEAMS    = 3
	NUM_PER_TEAM = 3
)

var (
	POSSIBLE_MOVEMENTS = board.DIRECTIONS_ALL
)

// Tile Types
var (
	tSand *board.Pixel = &board.Pixel{
		Colour: board.Colour{
			Foreground: board.Yellow,
			Background: board.LightYellow,
		},
		Chars: []string{
			"',",
		},
	}
	tBody *board.Pixel = &board.Pixel{
		Colour: board.Colour{
			Foreground: board.Black,
			Background: board.Red,
		},
		Chars: []string{
			"xx",
		},
	}
	arena *board.Board
)

func main() {

	arena = board.NewBoard(ARENA_WIDTH, ARENA_HEIGHT, board.ASCIIBorder, tSand)

	rand.Seed(time.Now().UnixNano())
	rand.Seed(rand.Int63())

	for t := 0; t < NUM_TEAMS; t++ {
		team := Team(
			board.Colour{
				Foreground: board.Black,
				Background: GLADIATOR_COLOURS[rand.Intn(len(GLADIATOR_COLOURS))],
			},
		)
		for f := 0; f < NUM_PER_TEAM; f++ {
			glad := GenerateGladiator(team)
			pos := board.RandomPos(arena.GetWidth(), arena.GetHeight())
			arena.SetPiece(pos, glad)
		}
	}

	for {
		board.ClearScreen()
		fmt.Print(arena.RenderString())

		fmt.Println(" Current Gladiators:")
		fmt.Println("---------------------")

		for _, pos := range arena.GetPieceCoords() {
			piece, exist := arena.GetPiece(pos)
			if !exist {
				return
			}
			glad, ok := piece.(*Gladiator)
			if !ok {
				return
			}
			nameOffset := ""
			for i := 0; i < GLADIATOR_NAME_LEN-len(glad.name); i++ {
				nameOffset += " "
			}
			fmt.Printf(" %s%s%s - %s%s (%d/%d) %s\n",
				glad.GetANSIString(),
				"oo",
				board.ANSI_RESET_COLOUR,
				glad.name,
				nameOffset,
				glad.speed,
				glad.strength,
				glad.GetHealthbar(),
			)
			MoveGladiator(pos, arena)
		}

		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
