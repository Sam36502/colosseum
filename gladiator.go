/*
 *
 *	Handles Fighters
 *
 */
package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/Sam36502/go-board/v2"
)

const (
	MAX_SPEED          = 5
	MAX_STRENGTH       = 10
	MAX_HEALTH         = 20
	DMG_ROLL           = 3 // Maximum random damage added to an attack
	HEALTHBAR_LEN      = 10
	GLADIATOR_NAME_LEN = 15
	DEATH_MSG          = "DEAD!"
)

var GLADIATOR_NAMES []string = []string{
	"Achilles",
	"Agamemnon",
	"Ajax",
	"Hector",
	"Aeneas",
	"Odysseus",
	"Diomedes",
	"Patroclus",
	"Menelaus",
	"Antilochus",
	"Thrasymedes",
}

var GLADIATOR_COLOURS []string = []string{
	board.LightBlue,
	board.LightCyan,
	board.LightGreen,
	board.LightMagenta,
	board.LightRed,
	board.LightWhite,
	board.LightYellow,
}

type Team board.Colour

type Gladiator struct {
	name     string
	shape    board.Pixel
	team     Team
	speed    int
	strength int
	health   int
}

// Ensure Gladiator implements Tile
var _ board.Tile = (*Gladiator)(nil)

func GenerateGladiator(team Team) *Gladiator {
	return &Gladiator{
		name: GLADIATOR_NAMES[rand.Intn(len(GLADIATOR_NAMES))],
		shape: board.Pixel{
			Colour: board.Colour(team),
			Chars: []string{
				"oo",
			},
		},
		team:     team,
		speed:    rand.Intn(MAX_SPEED),
		strength: rand.Intn(MAX_STRENGTH),
		health:   rand.Intn(MAX_HEALTH),
	}
}

func MoveGladiator(pos board.Coord, arena *board.Board) {
	piece, exists := arena.GetPiece(pos)
	if !exists {
		return
	}
	glad, ok := piece.(*Gladiator)
	if !ok {
		fmt.Println("Piece was not a gladiator.")
		return
	}

	arena.MovePiece(pos, board.RandomDirection(board.DIRECTIONS_ALL).Scale(glad.speed))

	// Find all soldiers within range
	for _, dir := range board.DIRECTIONS_ALL {
		enemypos := pos.Add(dir)
		if tile, exists := arena.GetPiece(enemypos); exists {
			enemy, ok := tile.(*Gladiator)
			if !ok {
				fmt.Println("(inner) Piece was not a gladiator.")
				return
			}

			if enemy.team != glad.team {
				enemy.health -= rand.Intn(DMG_ROLL) + glad.strength
				if enemy.health <= 0 {
					arena.DeletePiece(enemypos)
					arena.SetTile(enemypos, tBody)
				}
			}
		}
	}
}

func (g *Gladiator) GetHealthbar() string {
	healthColour := board.Colour{
		Foreground: board.LightRed,
		Background: board.Black,
	}
	bar := healthColour.GetANSIString() + "["

	if g.health > 0 {
		hearts := int(math.Ceil(float64(g.health) / (float64(MAX_HEALTH) / float64(HEALTHBAR_LEN))))
		for i := 0; i < hearts; i++ {
			bar += "#"
		}
		for i := 0; i < HEALTHBAR_LEN-hearts; i++ {
			bar += " "
		}
	} else {
		buf := int(math.Floor(float64(HEALTHBAR_LEN-len(DEATH_MSG)) / float64(2)))
		i := -1
		if HEALTHBAR_LEN-len(DEATH_MSG)%2 == 0 {
			i = 0
		}

		for ; i < buf; i++ {
			bar += " "
		}
		bar += DEATH_MSG
		for i = 0; i < buf; i++ {
			bar += " "
		}
	}
	return bar + "]" + board.ANSI_RESET_COLOUR
}

func (g *Gladiator) SetColour(c board.Colour) { g.shape.SetColour(c) }
func (g *Gladiator) SetChars(c []string)      { g.shape.SetChars(c) }
func (g *Gladiator) GetANSIString() string    { return g.shape.GetANSIString() }
func (g *Gladiator) GetChars() []string       { return g.shape.GetChars() }
func (g *Gladiator) GetWidth() int            { return g.shape.GetWidth() }
func (g *Gladiator) GetHeight() int           { return g.shape.GetHeight() }
