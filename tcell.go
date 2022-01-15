package main

import "github.com/gdamore/tcell/v2"

// createScreen generates a ready to use screen. The screen has
// no cursor and doesn't support mouse eventing.
func createScreen() (tcell.Screen, error) {
	screen, screenCreationError := tcell.NewScreen()
	if screenCreationError != nil {
		return nil, screenCreationError
	}

	screenInitError := screen.Init()
	if screenInitError != nil {
		return nil, screenInitError
	}

	//Make sure it's disable, even though it should be by default.
	screen.DisableMouse()
	//Make sure cursor is hidden by default.
	screen.HideCursor()

	return screen, nil
}

func eventIsRune(event *tcell.EventKey, r rune) bool {
	return event.Key() == tcell.KeyRune && event.Rune() == r && event.Modifiers() == 0
}

func drawRectangle(screen tcell.Screen, xStart, yStart, width, height int, style tcell.Style) {
	for y := yStart; y < yStart+height; y++ {
		for x := xStart; x < xStart+width; x++ {
			screen.SetContent(x, y, 0, nil, style)
		}
	}
}
