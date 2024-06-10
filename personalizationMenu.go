package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func personalizationMenu(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}.Layout(gtx,
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
						return layout.Center.Layout(gtx,
							func(gtx C) D {
								return material.Editor(g.theme, g.nickname, "Nickname").Layout(gtx)
							})
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
						return layout.Center.Layout(gtx,
							func(gtx C) D {
								return material.Editor(g.theme, g.profileDescription, "Profile Description").Layout(gtx)
							})
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
						btn := material.Button(g.theme, g.acceptButton, "Accept")
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
						btn := material.Button(g.theme, g.discardButton, "Discard")
						return btn.Layout(gtx)
					},
				)
			},
		),
	)

}
