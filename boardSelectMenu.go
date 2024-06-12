package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func displayBoardSelectMenuBoardMenu(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				title := material.H3(g.theme, "Select your ship positions")
				title.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				return title.Layout(gtx)
			})
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, slimButtonRow(g.selectionIndicatorButtons, g.theme, g.selectionIincidatorState)...)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, buttonWidgets(g.leftTableButtons, g.leftTableLabels, g.leftTableStates, g.theme, &g.lockLeftTable)...)
		}),
	)
}

func boardSelectMenu(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly}.Layout(gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return displayBoardSelectMenuBoardMenu(gtx, g)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return boardSelectMenuButtons(gtx, g)
		}),
	)
}

func boardSelectMenuButtons(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Alignment(layout.E),
		Spacing:   layout.SpaceEvenly,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx C) D {
				margins := layout.Inset{
					Top:    unit.Dp(10),
					Bottom: unit.Dp(10),
					Right:  unit.Dp(10),
					Left:   unit.Dp(10),
				}
				return margins.Layout(gtx,
					func(gtx C) D {
						btn := material.Button(g.theme, g.acceptShipPositions, "Accept")
						return btn.Layout(gtx)
					},
				)
			},
		),
		layout.Rigid(
			func(gtx C) D {
				margins := layout.Inset{
					Top:    unit.Dp(10),
					Bottom: unit.Dp(10),
					Right:  unit.Dp(10),
					Left:   unit.Dp(10),
				}
				return margins.Layout(gtx,
					func(gtx C) D {
						btn := material.Button(g.theme, g.discardShipPositions, "Discard")
						return btn.Layout(gtx)
					},
				)
			},
		),
		layout.Rigid(
			func(gtx C) D {
				margins := layout.Inset{
					Top:    unit.Dp(10),
					Bottom: unit.Dp(10),
					Right:  unit.Dp(10),
					Left:   unit.Dp(10),
				}
				return margins.Layout(gtx,
					func(gtx C) D {
						btn := material.Button(g.theme, g.randomShipPositions, "Random \n positions")
						return btn.Layout(gtx)
					},
				)
			},
		),
		layout.Rigid(
			func(gtx C) D {
				margins := layout.Inset{
					Top:    unit.Dp(10),
					Bottom: unit.Dp(10),
					Right:  unit.Dp(10),
					Left:   unit.Dp(10),
				}
				return margins.Layout(gtx,
					func(gtx C) D {
						btn := material.Button(g.theme, g.backShipPositions, "back")
						return btn.Layout(gtx)
					},
				)
			},
		),
	)
}
