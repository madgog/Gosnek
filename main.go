package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen tcell.Screen
  direction string
  speed time.Duration
	snek   *Snek
	snekX  int
	snekY  int
	foodX  int
	foodY  int
	score  int
}

func (g *Game) drawScore(s tcell.Screen, x1, y1, x2, y2 int, score *int) {
	row := y1
	col := x1
	text := "Score: " + strconv.Itoa(*score)
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.ColorReset).Attributes(tcell.AttrBold)
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (g *Game) drawBox(s tcell.Screen, x1, y1, x2, y2 int, fx, fy *int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// draw food
	s.SetContent(*fx, *fy, tcell.RuneDiamond, nil, tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack))
}

func (g *Game) checkIfAteFood() {
	if g.snek.head.x == g.foodX && g.snek.head.y == g.foodY {
		g.snek.Insert(g.snek.head.x, g.snek.head.y)
		g.foodX = rand.Intn(45)
		g.foodY = rand.Intn(23)
		g.score += 1
	}
}

func move(x, y *int, screen tcell.Screen, snek *Snek, score, fx, fy *int) {

  game := &Game {

  } 
	screen.Clear()
	drawBox(screen, 1, 1, 50, 25, foodX, foodY)
	drawScore(screen, 21, 26, 50, 26, score)
	checkIfAteFood()
	snek.Move(x, y)
	snek.DrawSnek(screen)
}

func (g *Game) run() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	snek := Snek{}
	snek.Insert(3, 1)
	snek.Insert(2, 1)
	snek.Insert(1, 1)

	// Score
	score := 0
	// Snek coodrinates
	x, y := snek.head.x, snek.head.y

	// Food coodrinates
	fx, fy := rand.Intn(45), rand.Intn(23)

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	screen.SetStyle(defStyle)
	screen.Clear()

	g.drawBox(screen, 1, 1, 50, 25, &fx, &fy)

	snek.DrawSnek(screen)

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}

	defer quit()
	// Clear screen
	for {
		// Update screen
		screen.Show()

		// Poll event
		ev := screen.PollEvent()

		// Process event
		switch ev := ev.(type) {

		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				screen.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				screen.Clear()
			}
			if ev.Key() == tcell.KeyRight {
				if x >= 50 {
					x = 0
				}
				x = x + 1
			}
			if ev.Key() == tcell.KeyLeft {
				if x <= 1 {
					x = 51
				}
				x = x - 1
			}
			if ev.Key() == tcell.KeyUp {
				if y <= 1 {
					y = 26
				}
				y = y - 1
			}
			if ev.Key() == tcell.KeyDown {
				if y >= 25 {
					y = 0
				}
				y = y + 1
			}
			g.move(&x, &y, screen, &snek, &score, &fx, &fy)
		}
	}

}

func main() {
	direction := make(chan string)
	go move(direction)
	run()
}
