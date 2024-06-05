package main

import (
	"fmt"

	"gioui.org/widget"
)

func createTable() ([][]*widget.Clickable, [][]string, [][]int) {
	buttons := make([][]*widget.Clickable, 10)
	labels := make([][]string, 10)
	states := make([][]int, 10)
	for i := range buttons {
		buttons[i] = make([]*widget.Clickable, 10)
		labels[i] = make([]string, 10)
		states[i] = make([]int, 10)
		for j := range buttons[i] {
			buttons[i][j] = new(widget.Clickable)
			labels[i][j] = fmt.Sprintf("%c%d", 'A'+i, j+1)
			states[i][j] = Empty
		}
	}
	return buttons, labels, states
}
