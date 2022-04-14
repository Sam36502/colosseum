package main

import (
	"fmt"

	"github.com/Sam36502/go-board/v2"
	"github.com/mgutz/ansi"
)

// Structs //

type Weapon struct {
	att int
	def int
}

type Consumable struct {
	heal int
	pois int
}

// Item Defs //

var (
	WeaponShortSword = Weapon{
		att: 5,
		def: 3,
	}
	WeaponLongSword = Weapon{
		att: 10,
		def: 4,
	}
	WeaponRoundShield = Weapon{
		att: 1,
		def: 5,
	}
	WeaponTowerShield = Weapon{
		att: 2,
		def: 10,
	}
	ConsHealthPotion = Consumable{
		heal: 10,
		pois: 0,
	}
	ConsPoisonPotion = Consumable{
		heal: 0,
		pois: 10,
	}
)

//// Tile implementations ////
// Weapon Tile
func (w *Weapon) SetColours(fg, bg board.Colour) {}

func (w *Weapon) SetChars([]string) {}

func (w *Weapon) GetColourCode() string {
	return ansi.ColorCode(fmt.Sprint(board.LightBlack, ":", board.Yellow))
}

func (w *Weapon) GetChars() []string {
	return []string{
		"⚔⦈",
	}
}

func (w *Weapon) GetSize() (int, int) {
	return 2, 1
}

// Consumable Tile
func (c *Consumable) SetColours(fg, bg board.Colour) {}

func (c *Consumable) SetChars([]string) {}

func (c *Consumable) GetColourCode() string {
	return ansi.ColorCode(fmt.Sprint(board.LightGreen, ":", board.Yellow))
}

func (c *Consumable) GetChars() []string {
	return []string{
		"ó♙",
	}
}

func (c *Consumable) GetSize() (int, int) {
	return 2, 1
}
