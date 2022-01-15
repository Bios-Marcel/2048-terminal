package state

import (
	"math/rand"
	"sync"
)

type GameSession struct {
	Mutex                     *sync.Mutex
	renderNotificationChannel chan bool

	GameOver  bool
	GameBoard [4][4]uint
}

// NewGameSession produces a ready-to-use session state.
func NewGameSession(renderNotificationChannel chan bool) *GameSession {
	session := &GameSession{
		Mutex:                     &sync.Mutex{},
		renderNotificationChannel: renderNotificationChannel,

		GameOver:  false,
		GameBoard: [4][4]uint{},
	}

	//We want to start off with one filled cell.
	session.fillCell()

	return session
}

func (session *GameSession) update() {
	// In order to avoid dead-locking the caller.
	go func() {
		session.renderNotificationChannel <- true
	}()
}

func (session *GameSession) fillCell() {
	if session.GameOver {
		return
	}

	var freeIndices [][2]int
	for rowIndex, row := range session.GameBoard {
		for cellIndex, cell := range row {
			if cell == 0 {
				freeIndices = append(freeIndices, [2]int{rowIndex, cellIndex})
			}
		}
	}

	if len(freeIndices) == 0 {
		session.GameOver = true
		return
	}

	indexToFill := freeIndices[rand.Intn(len(freeIndices))]
	session.GameBoard[indexToFill[0]][indexToFill[1]] = 2
}

func (session *GameSession) Down() {
	session.downNoFill()
	session.fillCell()
	session.update()
}

// downNoFill is necessary for proper unit testing without the
// randomness factor.
func (session *GameSession) downNoFill() {
	if session.GameOver {
		return
	}

	for cellIndex := 0; cellIndex < len(session.GameBoard); cellIndex++ {
		//Combination run
		//We combine from top to bottom, since that's how the original game
		//does it. So 2,2,2,0 would become 4,0,2,0
		session.combineVertically(
			len(session.GameBoard)-1,
			func(i int) bool { return i >= 0 },
			func(i int) int { return i - 1 },
			cellIndex)

		//Shifting run
		//The previously combined 4,0,2,0 now becomes 4,2,0,0
		for rowIndex := len(session.GameBoard) - 2; rowIndex >= 0; rowIndex-- {
			cell := session.GameBoard[rowIndex][cellIndex]
			if cell == 0 {
				continue
			}

			moveTo := -1
			for tempRowIndex := rowIndex + 1; tempRowIndex < len(session.GameBoard); tempRowIndex++ {
				if session.GameBoard[tempRowIndex][cellIndex] == 0 {
					moveTo = tempRowIndex
				} else {
					break
				}
			}

			if moveTo != -1 {
				session.GameBoard[moveTo][cellIndex] = cell
				session.GameBoard[rowIndex][cellIndex] = 0
			}
		}
	}
}

func (session *GameSession) Up() {
	session.upNoFill()
	session.fillCell()
	session.update()
}

func (session *GameSession) upNoFill() {
	if session.GameOver {
		return
	}

	for cellIndex := 0; cellIndex < len(session.GameBoard); cellIndex++ {
		//Combination run
		//We combine from top to bottom, since that's how the original game
		//does it. So 2,2,2,0 would become 4,0,2,0
		session.combineVertically(
			0,
			func(i int) bool { return i < len(session.GameBoard) },
			func(i int) int { return i + 1 },
			cellIndex)

		//Shifting run
		//The previously combined 4,0,2,0 now becomes 4,2,0,0
		for rowIndex := 1; rowIndex < len(session.GameBoard); rowIndex++ {
			cell := session.GameBoard[rowIndex][cellIndex]
			if cell == 0 {
				continue
			}

			moveTo := -1
			for tempRowIndex := rowIndex - 1; tempRowIndex >= 0; tempRowIndex-- {
				if session.GameBoard[tempRowIndex][cellIndex] == 0 {
					moveTo = tempRowIndex
				} else {
					break
				}
			}

			if moveTo != -1 {
				session.GameBoard[moveTo][cellIndex] = cell
				session.GameBoard[rowIndex][cellIndex] = 0
			}
		}
	}
}

func (session *GameSession) combineVertically(start int, resume func(int) bool, update func(int) int, cellIndex int) {
	indexLastNonZero := -1
	for rowIndex := start; resume(rowIndex); rowIndex = update(rowIndex) {
		cell := session.GameBoard[rowIndex][cellIndex]
		if cell == 0 {
			continue
		}

		if indexLastNonZero == -1 || cell != session.GameBoard[indexLastNonZero][cellIndex] {
			indexLastNonZero = rowIndex
			continue
		}

		session.GameBoard[indexLastNonZero][cellIndex] = cell * 2
		session.GameBoard[rowIndex][cellIndex] = 0
		indexLastNonZero = -1
	}
}

func (session *GameSession) Left() {
	session.leftNoFill()
	session.fillCell()
	session.update()
}

func (session *GameSession) leftNoFill() {
	if session.GameOver {
		return
	}

	for rowIndex := 0; rowIndex < len(session.GameBoard); rowIndex++ {
		//Combination run
		session.combineHorizontally(0,
			func(i int) bool { return i < len(session.GameBoard) },
			func(i int) int { return i + 1 },
			rowIndex)

		//Shifting run
		//The previously combined 4,0,2,0 now becomes 4,2,0,0
		for cellIndex := 1; cellIndex < len(session.GameBoard); cellIndex++ {
			cell := session.GameBoard[rowIndex][cellIndex]
			if cell == 0 {
				continue
			}

			moveTo := -1
			for tempCellIndex := cellIndex - 1; tempCellIndex >= 0; tempCellIndex-- {
				if session.GameBoard[rowIndex][tempCellIndex] == 0 {
					moveTo = tempCellIndex
				} else {
					break
				}
			}

			if moveTo != -1 {
				session.GameBoard[rowIndex][moveTo] = cell
				session.GameBoard[rowIndex][cellIndex] = 0
			}
		}
	}

}

func (session *GameSession) Right() {
	session.rightNoFill()
	session.fillCell()
	session.update()
}

func (session *GameSession) rightNoFill() {
	if session.GameOver {
		return
	}

	for rowIndex := 0; rowIndex < len(session.GameBoard); rowIndex++ {
		//Combination run
		//We combine from top to bottom, since that's how the original game
		//does it. So 2,2,2,0 would become 4,0,2,0
		session.combineHorizontally(
			len(session.GameBoard)-1,
			func(i int) bool { return i >= 0 },
			func(i int) int { return i - 1 },
			rowIndex)

		//Shifting run
		//The previously combined 4,0,2,0 now becomes 4,2,0,0
		for cellIndex := len(session.GameBoard) - 2; cellIndex >= 0; cellIndex-- {
			cell := session.GameBoard[rowIndex][cellIndex]
			if cell == 0 {
				continue
			}

			moveTo := -1
			for tempCellIndex := cellIndex + 1; tempCellIndex < len(session.GameBoard); tempCellIndex++ {
				if session.GameBoard[rowIndex][tempCellIndex] == 0 {
					moveTo = tempCellIndex
				} else {
					break
				}
			}

			if moveTo != -1 {
				session.GameBoard[rowIndex][moveTo] = cell
				session.GameBoard[rowIndex][cellIndex] = 0
			}
		}
	}
}

func (session *GameSession) combineHorizontally(start int, resume func(int) bool, update func(int) int, rowIndex int) {
	indexLastNonZero := -1
	for cellIndex := start; resume(cellIndex); cellIndex = update(cellIndex) {
		cell := session.GameBoard[rowIndex][cellIndex]
		if cell == 0 {
			continue
		}

		if indexLastNonZero == -1 || cell != session.GameBoard[rowIndex][indexLastNonZero] {
			indexLastNonZero = cellIndex
			continue
		}

		session.GameBoard[rowIndex][indexLastNonZero] = cell * 2
		session.GameBoard[rowIndex][cellIndex] = 0
		indexLastNonZero = -1
	}
}
