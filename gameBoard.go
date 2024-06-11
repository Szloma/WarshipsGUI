package main

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func displayPlayerAndEnemyBoardInside(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			timerText := fmt.Sprintf("time\n%d", g.timeLeft)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H4(g.theme, timerText).Layout(gtx)
				})
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			timerText := fmt.Sprintf("accuracy\n%f", g.accuracy)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H4(g.theme, timerText).Layout(gtx)
				})
			})
		}),
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

						btn := material.Button(g.theme, g.abandonButton, "abandon\ngame")
						btn.Background = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
						return btn.Layout(gtx)
					},
				)
			},
		),
	)
}

func displayPlayerAndEnemyBoardWithoutLabels(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.Spacing(layout.Baseline)}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, buttonWidgets(g.leftTableButtons, g.leftTableLabels, g.leftTableStates, g.theme)...)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return displayPlayerAndEnemyBoardInside(gtx, g)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, buttonWidgets(g.rightTableButtons, g.rightTableLabels, g.rightTableStates, g.theme)...)
		}),
	)
}

func displayPlayerAndEnemyBoard(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical, Spacing: layout.Spacing(layout.Baseline)}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return displayPlayerAndEnemyBoardWithoutLabels(gtx, g)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return displayEnemyNameAndDescription(gtx, g)
		}),
	)

}
