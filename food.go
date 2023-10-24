package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type Food struct {
	x int
	y int
}

func (F *Food) SetXY() {
	F.x = rand.Intn(45)
	F.y = rand.Intn(23)
}

func (F *Food) DrawFood(s tcell.Screen) {
	style := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorReset)
	x := F.x
	y := F.y
	s.SetContent(x, y, tcell.RuneBlock, nil, style)
}
