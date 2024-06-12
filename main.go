package main

import (
	"fmt"
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
	theme                      *material.Theme
	startButton                *widget.Clickable
	exitButton                 *widget.Clickable
	personalizationButton      *widget.Clickable
	acceptButton               *widget.Clickable
	discardButton              *widget.Clickable
	nickname                   *widget.Editor
	profileDescription         *widget.Editor
	acceptShipPositions        *widget.Clickable
	backShipPositions          *widget.Clickable
	discardShipPositions       *widget.Clickable
	randomShipPositions        *widget.Clickable
	abandonButton              *widget.Clickable
	gameOverButton             *widget.Clickable
	showStatsButton            *widget.Clickable
	backFromStatsButton        *widget.Clickable
	youLoseScreen              bool
	displayPlayerAndEnemyBoard bool
	showShipSetUpMenu          bool
	showLeftTable              bool
	showTables                 bool
	showPersonalization        bool
	showStartMenu              bool
	inGame                     bool
	showLoadingMenu            bool
	showGameOver               bool
	showGameWon                bool
	showLeaderBoards           bool
	showStats                  bool
	selectionIincidatorState   [20]int
	leftShip                   int
	selectionIndicatorButtons  []*widget.Clickable
	leftTableButtons           [][]*widget.Clickable
	leftTableLabels            [][]string
	leftTableStates            [][]int
	rightTableButtons          [][]*widget.Clickable
	rightTableLabels           [][]string
	rightTableStates           [][]int
	accuracy                   float64
	timeLeft                   int
	enemyName                  string
	enemyDescription           string
}

func NewGUI() *GUI {
	gui := &GUI{
		theme:                      material.NewTheme(),
		startButton:                new(widget.Clickable),
		exitButton:                 new(widget.Clickable),
		personalizationButton:      new(widget.Clickable),
		acceptButton:               new(widget.Clickable),
		discardButton:              new(widget.Clickable),
		nickname:                   new(widget.Editor),
		profileDescription:         new(widget.Editor),
		acceptShipPositions:        new(widget.Clickable),
		discardShipPositions:       new(widget.Clickable),
		randomShipPositions:        new(widget.Clickable),
		backShipPositions:          new(widget.Clickable),
		abandonButton:              new(widget.Clickable),
		gameOverButton:             new(widget.Clickable),
		showStatsButton:            new(widget.Clickable),
		backFromStatsButton:        new(widget.Clickable),
		leftShip:                   20,
		showStats:                  false,
		showLeaderBoards:           false,
		showGameWon:                false,
		showGameOver:               false,
		showLoadingMenu:            false,
		youLoseScreen:              false,
		inGame:                     false,
		displayPlayerAndEnemyBoard: false,
		showShipSetUpMenu:          false,
		showLeftTable:              false,
		showTables:                 false,
		showPersonalization:        false,
		showStartMenu:              true,
		timeLeft:                   60,
		accuracy:                   0.0,
		enemyName:                  "Janusz",
		enemyDescription:           "aaa",
	}
	gui.leftTableButtons, gui.leftTableLabels, gui.leftTableStates = createTable()
	gui.rightTableButtons, gui.rightTableLabels, gui.rightTableStates = createTable()
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
				g.showShipSetUpMenu = true
				fmt.Println("test")
				InitGame()
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
			if g.backFromStatsButton.Clicked(gtx) {
				g.showStartMenu = true
				g.showStats = false
			}
			if g.showStatsButton.Clicked(gtx) {
				g.showStartMenu = false
				g.showStats = true
			}
			if g.exitButton.Clicked(gtx) {
				os.Exit(0)
			}

			if g.showLeftTable {
				handleTableClicks(gtx, g)

			}
			if g.acceptShipPositions.Clicked(gtx) {
				g.showShipSetUpMenu = false
				g.displayPlayerAndEnemyBoard = true
				g.inGame = true
				fmt.Printf("accepted ship positions")
			}
			if g.discardShipPositions.Clicked(gtx) {
				g.leftTableStates = createEmptyState(10, 10)
				fmt.Printf("discarded ship positions")
			}
			if g.randomShipPositions.Clicked(gtx) {
				g.showShipSetUpMenu = false
				//tmp for testing
				g.showLoadingMenu = true
				fmt.Printf("random ship positions")
			}
			if g.backShipPositions.Clicked(gtx) {
				g.showStartMenu = true
				g.displayPlayerAndEnemyBoard = false
			}

			if g.abandonButton.Clicked(gtx) {
				g.displayPlayerAndEnemyBoard = false
				g.showStartMenu = true
			}

			if g.inGame {
				//time.Sleep(time.Second)
				// to bddzie z requestu, nie ma co robic timera
				print("game in progress")
				//g.timeLeft -= 1
			}

			Layout(gtx, g)

			e.Frame(gtx.Ops)
		}
	}
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

func displayEnemyNameAndDescription(gtx layout.Context, g *GUI) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			timerText := fmt.Sprintf("Enemy Name: %s", g.enemyName)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H4(g.theme, timerText).Layout(gtx)
				})
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			timerText := fmt.Sprintf("Enemy Description: %s", g.enemyDescription)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H6(g.theme, timerText).Layout(gtx)
				})
			})
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
	if g.showShipSetUpMenu {
		return boardSelectMenu(gtx, g)
	}
	if g.displayPlayerAndEnemyBoard {
		return displayPlayerAndEnemyBoard(gtx, g)
	}
	if g.showLoadingMenu {
		return loadingMenu(gtx, g)
	}
	if g.showGameOver {
		return gameOver(gtx, g)
	}
	if g.showGameWon {
		return gameWon(gtx, g)
	}
	if g.showStats {
		return showStatsMenu(gtx, g)
	}
	return emptyLayoutDebug(gtx, g)
}
