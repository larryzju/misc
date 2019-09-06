package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	sleep = 1 * time.Second
)

var offsets = []struct{ x, y int }{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

// Life is a mathematical game which takes place on a two dimensional grid of cells
// Each cell can be either alive or dead.
// Cells interact with their direct neighbors to determine whether they will live or die in the next generation of cells
//
// For each new generation of cells, the follwing rules are applied:
// * Any live cell with fewer than two live neighbors dies of underpopulation
// * Any live cell with more than three live neighbors dies of overcrowding
// * Any live cell with two or three live neighbors lives on to the next generation
// * Any dead cell with exactly three live neighbors becomes a live cell
type Status byte

const (
	Dead Status = iota
	Live
)

func (s Status) String() string {
	if s == Live {
		return "O"
	}
	return " "
}

type Life struct {
	Cells  []Status
	width  int
	height int
}

func NewLife(width, height int) *Life {
	return &Life{
		Cells:  make([]Status, width*height),
		width:  width,
		height: height,
	}
}

func (life *Life) Cap() int {
	return life.width * life.height
}

func (life *Life) NextGen() *Life {
	next := NewLife(life.width, life.height)
	for i := 0; i < life.width; i++ {
		for j := 0; j < life.height; j++ {
			index := life.index(i, j)
			next.Cells[index] = life.nextStatus(i, j)
		}
	}
	return next
}

func (life *Life) index(i, j int) int {
	return j*life.width + i
}

func (life *Life) nextStatus(i, j int) Status {
	count := 0
	for _, off := range offsets {
		x := i + off.x
		y := j + off.y

		if x < 0 || x >= life.width || y < 0 || y >= life.height {
			continue
		}

		if life.Cells[life.index(x, y)] == Live {
			count += 1
		}
	}

	switch status := life.Cells[life.index(i, j)]; {
	case status == Dead && count == 3:
		return Live
	case status == Live && (count < 2 || count > 3):
		return Dead
	case status == Live:
		return Live
	default:
		return status
	}
}

func (life *Life) Dump() {
	fmt.Println(strings.Repeat("-", life.width))
	for i := 0; i < life.width; i++ {
		for j := 0; j < life.height; j++ {
			fmt.Printf("%s", life.Cells[life.index(i, j)])
		}
		fmt.Println()
	}
}

func (life *Life) Over() bool {
	for _, v := range life.Cells {
		if v == Live {
			return false
		}
	}
	return true
}

func (life *Life) Seed(i, j int) {
	life.Cells[life.index(i, j)] = Live
}

func main() {
	life := NewLife(20, 20)
	life.Seed(7, 6)
	life.Seed(6, 6)
	life.Seed(6, 5)
	life.Seed(6, 4)
	life.Seed(5, 6)
	life.Seed(5, 5)
	life.Seed(5, 4)
	life.Seed(4, 4)
	life.Seed(4, 5)
	life.Seed(4, 6)

	for !life.Over() {
		life.Dump()
		life = life.NextGen()
		time.Sleep(sleep)
	}
}
