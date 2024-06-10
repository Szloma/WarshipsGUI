package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func startMenu(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				title := material.H1(g.theme, "Gigawarships")
				title.Alignment = text.Middle
				title.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				return title.Layout(gtx)
			})
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
						btn := material.Button(g.theme, g.startButton, "Start")
						return btn.Layout(gtx)
					},
				)
			},
		), layout.Rigid(
			func(gtx C) D {
				margins := layout.Inset{
					Top:    unit.Dp(25),
					Bottom: unit.Dp(25),
					Right:  unit.Dp(35),
					Left:   unit.Dp(35),
				}

				return margins.Layout(gtx,
					func(gtx C) D {
						btn := material.Button(g.theme, g.personalizationButton, "Personalisation")
						return btn.Layout(gtx)
					},
				)
			},
		),
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
						btn := material.Button(g.theme, g.exitButton, "exit")
						return btn.Layout(gtx)
					},
				)
			},
		),
	)

}
