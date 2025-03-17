package main

import "github.com/fj68/hyper-tux-go/hyper"

const CELL_SIZE = 20 // px

func main() {
	board, err := hyper.NewBoard(&hyper.Size{W: 16, H: 16})
	if err != nil {
		panic(err)
	}
	if err = board.NewGame(); err != nil {
		panic(err)
	}
}
