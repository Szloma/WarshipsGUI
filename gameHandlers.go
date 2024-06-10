package main

func (g *GUI) GetShipPositions() []string {
	var shipPositions []string
	for i := range g.leftTableStates {
		for j, state := range g.leftTableStates[i] {
			if state == Ship {
				shipPositions = append(shipPositions, g.leftTableLabels[i][j])
			}
		}
	}
	return shipPositions
}
