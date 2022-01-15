package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/Bios-Marcel/2048/state"
	"github.com/gdamore/tcell/v2"
)

func main() {
	//Otherwise we'll always start on the same tile.
	rand.Seed(time.Now().Unix())

	screen, screenCreationError := createScreen()
	if screenCreationError != nil {
		panic(screenCreationError)
	}

	//Cleans up the terminal buffer and returns it to the shell.
	defer screen.Fini()

	//renderer used for drawing the board and the menu.
	renderer := newRenderer()

	renderNotificationChannel := make(chan bool)
	gameSession := state.NewGameSession(renderNotificationChannel)

	//Gameloop; We draw whenever there's a frame-change. This means we
	//don't have any specific frame-rates and it could technically happen
	//that we don't draw for a while. The first frame is drawn without
	//waiting for a change, so that the screen doesn't stay empty.

	//Listen for key input on the gameboard.
	go func() {
		for {
			switch event := screen.PollEvent().(type) {
			case *tcell.EventKey:
				if event.Key() == tcell.KeyCtrlC {
					screen.Fini()
					os.Exit(0)
				} else if event.Key() == tcell.KeyCtrlR {
					//RESTART!
					//Remove previous game over message and such and create
					//a fresh state, as we needn't save any information for
					//the next session.
					oldGameSession := gameSession
					oldGameSession.Mutex.Lock()

					//Make sure the state knows it's supposed to be dead.
					oldGameSession.GameOver = true
					screen.Clear()
					gameSession = state.NewGameSession(renderNotificationChannel)
					gameSession.Mutex.Lock()

					oldGameSession.Mutex.Unlock()
					gameSession.Mutex.Unlock()
					renderNotificationChannel <- true
				} else if event.Key() == tcell.KeyDown || eventIsRune(event, 's') {
					gameSession.Mutex.Lock()
					gameSession.Down()
					gameSession.Mutex.Unlock()
				} else if event.Key() == tcell.KeyUp || eventIsRune(event, 'w') {
					gameSession.Mutex.Lock()
					gameSession.Up()
					gameSession.Mutex.Unlock()
				} else if event.Key() == tcell.KeyLeft || eventIsRune(event, 'a') {
					gameSession.Mutex.Lock()
					gameSession.Left()
					gameSession.Mutex.Unlock()
				} else if event.Key() == tcell.KeyRight || eventIsRune(event, 'd') {
					gameSession.Mutex.Lock()
					gameSession.Right()
					gameSession.Mutex.Unlock()
				}
			case *tcell.EventResize:
				gameSession.Mutex.Lock()
				screen.Clear()
				gameSession.Mutex.Unlock()
				renderNotificationChannel <- true
				//TODO Handle resize; Validate session;
			default:
				//Unsupported or irrelevant event
			}
		}
	}()

	for {
		//We start lock before draw in order to avoid drawing crap.
		gameSession.Mutex.Lock()
		renderer.drawGameBoard(screen, gameSession)
		gameSession.Mutex.Unlock()

		<-renderNotificationChannel
	}
}
