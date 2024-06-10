package main

import (
	"image"
	"image/color"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
)

func loadingMenu(gtx layout.Context, g *GUI) layout.Dimensions {
	var angle float32

	layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				//defer op.Save(ops).Load()
				//op.Offset(f32.Pt(float32(gtx.Constraints.Min.X)/2, float32(gtx.Constraints.Min.Y)/2)).Add(ops)
				//	r := float32(gtx.Px(unit.Dp(24)))
				//	c := float32(gtx.Px(unit.Dp(8)))

				//	for i := 0; i < 12; i++ {
				//startAngle := angle + float32(i)*2*3.14/12
				//	endAngle := startAngle + 3.14/6
				//paintPath(gtx, startAngle, endAngle, r, c, color.NRGBA{A: 255 - uint8(i*255/12)})
				//}

				angle += 0.05
				if angle > 2*3.14 {
					angle -= 2 * 3.14
				}

				//

				return layout.Dimensions{Size: gtx.Constraints.Min}
			})
		}),
		layout.Rigid(
			func(gtx C) D {
				var circlePath clip.Path
				op.Offset(image.Pt(gtx.Dp(200), gtx.Dp(150))).Add(gtx.Ops)
				circlePath.Begin(gtx.Ops)
				for deg := 0.0; deg <= 360; deg++ {

					rad := deg / 360 * 2 * math.Pi
					cosT := math.Cos(rad)
					sinT := math.Sin(rad)
					a := 110.0
					x := a * cosT
					y := a * sinT
					p := f32.Pt(float32(x), float32(y))
					circlePath.LineTo(p)
				}
				circlePath.Close()

				eggArea := clip.Outline{Path: circlePath.End()}.Op()

				color := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				paint.FillShape(gtx.Ops, color, eggArea)

				d := image.Point{Y: 335}
				return layout.Dimensions{Size: d}
			},
		),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				lbl := material.H3(g.theme, "Loading")
				return lbl.Layout(gtx)
			})
		}),
	)

	return layout.Dimensions{Size: gtx.Constraints.Min}
}
