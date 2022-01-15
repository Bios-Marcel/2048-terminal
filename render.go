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

var (
	cellWidth             = 6
	cellHeight            = cellWidth / 2
	defaultCellBackground = tcell.StyleDefault.Reverse(true)
	styles                = map[uint]tcell.Style{
		2:    tcell.StyleDefault.Background(tcell.Color100),
		4:    tcell.StyleDefault.Background(tcell.Color101),
		8:    tcell.StyleDefault.Background(tcell.Color102),
		16:   tcell.StyleDefault.Background(tcell.Color103),
		32:   tcell.StyleDefault.Background(tcell.Color105),
		64:   tcell.StyleDefault.Background(tcell.Color106),
		128:  tcell.StyleDefault.Background(tcell.Color108),
		256:  tcell.StyleDefault.Background(tcell.Color109),
		1024: tcell.StyleDefault.Background(tcell.Color110),
		2048: tcell.StyleDefault.Background(tcell.Color111),
		4096: tcell.StyleDefault.Background(tcell.Color112),
		8192: tcell.StyleDefault.Background(tcell.Color113),
	}
)

func (renderer *renderer) drawGameBoard(screen tcell.Screen, session *state.GameSession) {
	for rowIndex, row := range session.GameBoard {
		for cellIndex, cell := range row {
			startX := cellIndex*cellWidth + cellIndex*2
			startY := rowIndex*cellHeight + rowIndex

			style, avail := styles[cell]
			if !avail {
				style = defaultCellBackground
			}

			for y := startY; y < startY+cellHeight; y++ {
				for x := startX; x < startX+cellWidth; x++ {
					screen.SetContent(x, y, ' ', nil, style)
				}
			}

			for index, r := range []rune(strconv.FormatUint(uint64(cell), 10)) {
				screen.SetContent(startX+index, startY+(cellHeight-1)/2, r, nil, style)
			}
		}
	}

	screen.Show()
}
