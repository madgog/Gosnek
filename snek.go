package main

import (
	"github.com/gdamore/tcell/v2"
)

type Node struct {
	prev *Node
	next *Node
	x    int
	y    int
}

type Snek struct {
	head *Node
	tail *Node
}

func (s *Snek) Insert(x, y int) {
	list := &Node{
		next: nil,
		x:    x,
		y:    y,
	}
	if s.head == nil {
		s.head = list
		s.tail = list
	} else {
		prev := s.tail
		s.tail.next = list
		s.tail = list
		s.tail.prev = prev
	}
}

func (S *Snek) Move(x, y *int) {
	tail := S.tail
	for tail.prev != nil {
		tail.x = tail.prev.x
		tail.y = tail.prev.y
		tail = tail.prev
	}
	S.head.x = *x
	S.head.y = *y
}

func (S *Snek) DrawSnek(screen tcell.Screen) {
	style := tcell.StyleDefault.Foreground(tcell.ColorGreenYellow)
	head := S.head
	for head != nil {
		screen.SetContent(head.x, head.y, tcell.RuneBlock, nil, style)
		head = head.next
	}

}
