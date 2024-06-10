package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		g := NewGUI()
		if err := loop(w, g); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type GUI struct {
	theme                     *material.Theme
	startButton               *widget.Clickable
	exitButton                *widget.Clickable
	personalizationButton     *widget.Clickable
	acceptButton              *widget.Clickable
	discardButton             *widget.Clickable
	nickname                  *widget.Editor
	profileDescription        *widget.Editor
	acceptShipPositions       *widget.Clickable
	discardShipPositions      *widget.Clickable
	randomShipPositions       *widget.Clickable
	showLeftTable             bool
	showTables                bool
	showPersonalization       bool
	showStartMenu             bool
	selectionIincidatorState  [20]int
	selectionIndicatorButtons []*widget.Clickable
	leftTableButtons          [][]*widget.Clickable
	leftShip                  int
	leftTableLabels           [][]string
	leftTableStates           [][]int
}

func NewGUI() *GUI {
	gui := &GUI{
		theme:                 material.NewTheme(),
		startButton:           new(widget.Clickable),
		exitButton:            new(widget.Clickable),
		personalizationButton: new(widget.Clickable),
		acceptButton:          new(widget.Clickable),
		discardButton:         new(widget.Clickable),
		nickname:              new(widget.Editor),
		profileDescription:    new(widget.Editor),
		acceptShipPositions:   new(widget.Clickable),
		discardShipPositions:  new(widget.Clickable),
		randomShipPositions:   new(widget.Clickable),
		leftShip:              20,
		showLeftTable:         false,
		showTables:            false,
		showPersonalization:   false,
		showStartMenu:         true,
	}
	gui.leftTableButtons, gui.leftTableLabels, gui.leftTableStates = createTable()
	gui.selectionIndicatorButtons = createButtonRow()
	gui.selectionIincidatorState = setSelectionIndidatorState(gui.leftShip)

	return gui
}

const (
	Empty = iota
	Ship
	Hit
	Miss
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func handleTableClicks(gtx layout.Context, g *GUI) {
	for i := range g.leftTableButtons {
		for y, btn := range g.leftTableButtons[i] {
			for btn.Clicked(gtx) {
				fmt.Printf("%s: %d\n", g.leftTableLabels[i][y], g.leftTableStates[i][y])
				//zwaluduj czy nowy stan był dobry
				if g.leftShip > 0 {
					if g.leftTableStates[i][y] == Ship {
						g.leftShip += 1
						g.leftTableStates[i][y] = Empty
					}
					if g.leftTableStates[i][y] == Empty {
						g.leftShip -= 1
						g.leftTableStates[i][y] = Ship
					}
					fmt.Printf("%d: ", g.leftShip)
					g.selectionIincidatorState = setSelectionIndidatorState(g.leftShip)
				}

				//g.leftTableStates[i][y] = Ship

				for s := range g.selectionIincidatorState {
					fmt.Print(" ")
					fmt.Print(g.selectionIincidatorState[s])
				}

			}
		}
	}
}

func loop(w *app.Window, g *GUI) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			if g.startButton.Clicked(gtx) {
				g.showStartMenu = false
				g.showLeftTable = true
				fmt.Println("test")
			}
			if g.personalizationButton.Clicked(gtx) {
				g.showStartMenu = false
				g.showPersonalization = true
			}
			if g.acceptButton.Clicked(gtx) {
				g.showStartMenu = true
				g.showPersonalization = false
				fmt.Printf("Nickname: %s\n", g.nickname.Text())
				fmt.Printf("Description: %s\n", g.profileDescription.Text())
			}
			if g.discardButton.Clicked(gtx) {
				g.showStartMenu = true
				g.showPersonalization = false
				fmt.Printf("Nickname: %s\n", g.nickname.Text())
				fmt.Printf("Description: %s\n", g.profileDescription.Text())
			}
			if g.exitButton.Clicked(gtx) {
				os.Exit(0)
			}

			if g.showLeftTable {
				handleTableClicks(gtx, g)

			}
			if g.acceptShipPositions.Clicked(gtx) {
				fmt.Printf("accepted ship positions")
			}
			if g.discardShipPositions.Clicked(gtx) {
				fmt.Printf("discarded ship positions")
			}
			if g.randomShipPositions.Clicked(gtx) {
				fmt.Printf("random ship positions")
			}

			Layout(gtx, g)

			e.Frame(gtx.Ops)
		}
	}
}

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
func emptyLayoutDebug(gtx layout.Context, g *GUI) layout.Dimensions {
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
						btn := material.Button(g.theme, g.startButton, "Start")
						return btn.Layout(gtx)
					},
				)
			},
		))
}

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

func (gui *GUI) renderLeftTable(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, buttonWidgets(gui.leftTableButtons, gui.leftTableLabels, gui.leftTableStates, gui.theme)...)
		}),
	)
}

func displayBoardSelectMenuSubMenu(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical, Spacing: layout.Spacing(layout.Center)}.Layout(gtx,
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
	)
}

func displayBoardSelectMenuBoardMenu(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceEvenly}.Layout(gtx,
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
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, buttonWidgets(g.leftTableButtons, g.leftTableLabels, g.leftTableStates, g.theme)...)
		}),
	)
}

func displayBoardSelectMenu(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return displayBoardSelectMenuBoardMenu(gtx, g)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return displayBoardSelectMenuSubMenu(gtx, g)
		}),
	)
}

func Layout(gtx layout.Context, g *GUI) layout.Dimensions {
	if g.showStartMenu {
		return startMenu(gtx, g)
	}
	if g.showPersonalization {
		return personalizationMenu(gtx, g)
	}
	if g.showLeftTable {
		return displayBoardSelectMenu(gtx, g)
	}
	return emptyLayoutDebug(gtx, g)
}
