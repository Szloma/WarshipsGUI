package main

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func setSelectionIndidatorState(n int) [20]int {
	var arr [20]int

	for i := 0; i < 20; i++ {
		arr[i] = 1
	}
	for i := 0; i < n; i++ {
		arr[i] = 0
	}
	return arr
}

func createButtonRow() []*widget.Clickable {
	buttons := make([]*widget.Clickable, 20)
	for i := range buttons {
		buttons[i] = new(widget.Clickable)
	}
	return buttons
}

func createEmptyState(rows, cols int) [][]int {
	state := make([][]int, rows)
	for i := range state {
		state[i] = make([]int, cols)
		for j := range state[i] {
			state[i][j] = 0
		}
	}
	return state
}

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

func buttonWidgets(buttons [][]*widget.Clickable, labels [][]string, states [][]int, th *material.Theme) []layout.FlexChild {
	var children []layout.FlexChild
	for i := 0; i < 10; i++ {
		i := i // capture range variable
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, buttonRow(buttons[i], labels[i], states[i], th)...)
		}))
	}
	return children
}

func slimButtonRow(buttons []*widget.Clickable, th *material.Theme, states [20]int) []layout.FlexChild {
	var children []layout.FlexChild
	index := 0
	for i, btn := range buttons {
		btn := btn
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			size := unit.Dp(20)
			if btn.Clicked(gtx) {
				states[i] = (states[i] + 1) % 4
				fmt.Printf("Indicator %d: %d\n", i, states[i])
			}
			btnWidget := material.Button(th, btn, fmt.Sprintf("Ind %d", i+1))
			switch states[i] {
			case Empty:
				btnWidget.Background = color.NRGBA{R: 0, G: 0, B: 255, A: 255}
			case Ship:
				btnWidget.Background = color.NRGBA{R: 100, G: 0, B: 0, A: 255}
			case Hit:
				btnWidget.Background = color.NRGBA{R: 255, G: 0, B: 0, A: 255}

			case Miss:
				btnWidget.Background = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
			}
			index += 1

			btnWidget.Inset = layout.UniformInset(unit.Dp(5))
			return layout.UniformInset(unit.Dp(1)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = int(gtx.Metric.DpToSp(size))
				gtx.Constraints.Max.X = int(gtx.Metric.DpToSp(size))
				gtx.Constraints.Min.Y = int(gtx.Metric.DpToSp(size))
				gtx.Constraints.Max.Y = int(gtx.Metric.DpToSp(size))
				return btnWidget.Layout(gtx)
			})
		}))
	}
	return children
}

func buttonRow(buttons []*widget.Clickable, labels []string, states []int, th *material.Theme) []layout.FlexChild {
	var children []layout.FlexChild
	for j, btn := range buttons {
		j := j
		btn := btn
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			size := unit.Dp(50)
			if btn.Clicked(gtx) {
				states[j] = (states[j] + 1) % 2
				fmt.Printf("%s: %d\n", labels[j], states[j])
			}
			btnWidget := material.Button(th, btn, labels[j])
			switch states[j] {
			case Empty:
				btnWidget.Background = color.NRGBA{R: 0, G: 0, B: 255, A: 255}

			case Ship:
				btnWidget.Background = color.NRGBA{R: 100, G: 0, B: 0, A: 255}
			case Hit:
				btnWidget.Background = color.NRGBA{R: 255, G: 0, B: 0, A: 255}

			case Miss:
				btnWidget.Background = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
			}

			btnWidget.Inset = layout.UniformInset(unit.Dp(5))
			return layout.UniformInset(unit.Dp(1)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = int(gtx.Metric.DpToSp(size))
				gtx.Constraints.Max.X = int(gtx.Metric.DpToSp(size))
				gtx.Constraints.Min.Y = int(gtx.Metric.DpToSp(size))
				gtx.Constraints.Max.Y = int(gtx.Metric.DpToSp(size))
				return btnWidget.Layout(gtx)
			})
		}))
	}
	return children
}
