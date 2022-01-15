package main

import (
	"strconv"

	"github.com/Bios-Marcel/2048/state"
	"github.com/gdamore/tcell/v2"
)

type renderer struct {
}

func newRenderer() *renderer {
	return nil
}

func (renderer *renderer) drawGameBoard(screen tcell.Screen, session *state.GameSession) {
	cellWidth := 6
	cellHeight := cellWidth / 2
	cellBackground := tcell.StyleDefault.Reverse(true)

	for rowIndex, row := range session.GameBoard {
		for cellIndex, cell := range row {
			startX := cellIndex*cellWidth + cellIndex*2
			startY := rowIndex*cellHeight + rowIndex
			for y := startY; y < startY+cellHeight; y++ {
				for x := startX; x < startX+cellWidth; x++ {
					screen.SetContent(x, y, ' ', nil, cellBackground)
				}
			}
			screen.SetContent(startX, startY+(cellHeight-1)/2, []rune(strconv.FormatUint(uint64(cell), 10))[0], nil, cellBackground)
		}
	}

	screen.Show()
}
