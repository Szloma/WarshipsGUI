package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func gameWonLabel(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gameOver := material.H2(g.theme, "You Won!")
		gameOver.Color = color.NRGBA{R: 000, G: 255, B: 0, A: 255} // Red color
		return gameOver.Layout(gtx)
	})
}

func gameWon(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return gameWonLabel(gtx, g)
		}),
		layout.Rigid(
			func(gtx C) D {
				margins := layout.Inset{
					Top:    unit.Dp(25),
					Bottom: unit.Dp(25),
					Right:  unit.Dp(35),
					Left:   unit.Dp(35),
				}
				return margins.Layout(gtx,
					func(gtx C) D {
						btn := material.Button(g.theme, g.gameOverButton, "Continue")
						return btn.Layout(gtx)
					},
				)
			},
		))

}
