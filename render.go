package main

import (
	"fmt"
	"strconv"

	"github.com/Bios-Marcel/2048-terminal/state"
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

			drawRectangle(screen, startX, startY, cellWidth, cellHeight, style)

			runes := []rune(strconv.FormatUint(uint64(cell), 10))
			xOffset := (cellWidth/2 - 1) - (len(runes)-1)/2
			for index, r := range runes {
				screen.SetContent(startX+xOffset+index, startY+(cellHeight-1)/2, r, nil, style)
			}
		}
	}

	if session.GameOver {
		boardWidth := cellWidth*len(session.GameBoard) + len(session.GameBoard) - 1
		boardHeight := cellHeight*len(session.GameBoard) + len(session.GameBoard) - 1
		gameOverBoxHeight := 3
		text := fmt.Sprintf("Game Over; Score: %d", session.Score())
		startX := boardWidth/2 - len(text)/2 + 1
		startY := boardHeight/2 - gameOverBoxHeight/2
		drawRectangle(screen, startX, startY, len(text)+2, gameOverBoxHeight, tcell.StyleDefault)
		startX = startX + 1
		for index, r := range text {
			screen.SetContent(startX+index, startY+(cellHeight-1)/2, r, nil, tcell.StyleDefault)
		}
	}

	screen.Show()
}
