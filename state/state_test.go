package state

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestGameSession_Down(t *testing.T) {
	tests := []shiftTest{
		{
			name: "combine twice in one column and shift both",
			board: [4][4]uint{
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.downNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{4, 0, 0, 0},
				{4, 0, 0, 0},
			},
		},
		{
			name: "combine once in one column and shift one cell (1)",
			board: [4][4]uint{
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.downNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{4, 0, 0, 0},
			},
		},
		{
			name: "combine once in one column and shift one cell (2)",
			board: [4][4]uint{
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.downNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{4, 0, 0, 0},
			},
		},
		{
			name: "combine once in one column and shift one cell (3)",
			board: [4][4]uint{
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.downNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{4, 0, 0, 0},
			},
		},
		{
			name: "shift one cell",
			board: [4][4]uint{
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.downNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
			},
		},
		{
			name: "combine twice and shift both, but last column",
			board: [4][4]uint{
				{0, 0, 0, 2},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
			},
			move: func(session *GameSession) func() bool { return session.downNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 4},
				{0, 0, 0, 4},
			},
		},
		{
			name: "do nothing",
			board: [4][4]uint{
				{0, 0, 0, 2},
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{0, 0, 0, 16},
			},
			move: func(session *GameSession) func() bool { return session.downNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 2},
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{0, 0, 0, 16},
			},
		},
	}

	runShiftTests(t, tests)
}

func TestGameSession_Up(t *testing.T) {
	tests := []shiftTest{
		{
			name: "combine twice in one column and shift both",
			board: [4][4]uint{
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.upNoFill },
			expectedBoard: [4][4]uint{
				{4, 0, 0, 0},
				{4, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "combine once in one column and shift one cell (1)",
			board: [4][4]uint{
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.upNoFill },
			expectedBoard: [4][4]uint{
				{4, 0, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "combine once in one column and shift one cell (2)",
			board: [4][4]uint{
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.upNoFill },
			expectedBoard: [4][4]uint{
				{4, 0, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "shift one cell",
			board: [4][4]uint{
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.upNoFill },
			expectedBoard: [4][4]uint{
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "shift one cell all the way",
			board: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.upNoFill },
			expectedBoard: [4][4]uint{
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "combine twice and shift both, but last column",
			board: [4][4]uint{
				{0, 0, 0, 2},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
			},
			move: func(session *GameSession) func() bool { return session.upNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 4},
				{0, 0, 0, 4},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "do nothing",
			board: [4][4]uint{
				{0, 0, 0, 2},
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{0, 0, 0, 16},
			},
			move: func(session *GameSession) func() bool { return session.upNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 2},
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{0, 0, 0, 16},
			},
		},
	}

	runShiftTests(t, tests)
}

func TestGameSession_Left(t *testing.T) {
	tests := []shiftTest{
		{
			name: "combine twice in one column and shift both",
			board: [4][4]uint{
				{2, 2, 2, 2},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.leftNoFill },
			expectedBoard: [4][4]uint{
				{4, 4, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
		},
		{
			name: "combine once in one column and shift one cell (1)",
			board: [4][4]uint{
				{0, 2, 2, 2},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.leftNoFill },
			expectedBoard: [4][4]uint{
				{4, 2, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
		},
		{
			name: "combine once in one column and shift one cell (2)",
			board: [4][4]uint{
				{2, 0, 2, 2},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.leftNoFill },
			expectedBoard: [4][4]uint{
				{4, 2, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
		},
		{
			name: "shift one cell",
			board: [4][4]uint{
				{0, 2, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.leftNoFill },
			expectedBoard: [4][4]uint{
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "shift one cell all the way",
			board: [4][4]uint{
				{0, 0, 0, 2},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.leftNoFill },
			expectedBoard: [4][4]uint{
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
			},
		},
		{
			name: "combine twice and shift both, but last column",
			board: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 2, 2, 2},
			},
			move: func(session *GameSession) func() bool { return session.leftNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{4, 4, 0, 0},
			},
		},
		{
			name: "do nothing",
			board: [4][4]uint{
				{2, 4, 8, 16},
				{4, 0, 0, 0},
				{8, 0, 0, 0},
				{16, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.leftNoFill },
			expectedBoard: [4][4]uint{
				{2, 4, 8, 16},
				{4, 0, 0, 0},
				{8, 0, 0, 0},
				{16, 0, 0, 0},
			},
		},
	}

	runShiftTests(t, tests)
}

func TestGameSession_Right(t *testing.T) {
	tests := []shiftTest{
		{
			name: "combine twice in one column and shift both",
			board: [4][4]uint{
				{2, 2, 2, 2},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.rightNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 4, 4},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
			},
		},
		{
			name: "combine once in one column and shift one cell (1)",
			board: [4][4]uint{
				{0, 2, 2, 2},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.rightNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 2, 4},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
			},
		},
		{
			name: "combine once in one column and shift one cell (2)",
			board: [4][4]uint{
				{2, 0, 2, 2},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.rightNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 2, 4},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
			},
		},
		{
			name: "shift one cell",
			board: [4][4]uint{
				{0, 2, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.rightNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 2},
				{0, 0, 0, 2},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "shift one cell all the way",
			board: [4][4]uint{
				{2, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 0, 0, 0},
			},
			move: func(session *GameSession) func() bool { return session.rightNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 2},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
			},
		},
		{
			name: "combine twice and shift both, but last column",
			board: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 2, 2, 2},
			},
			move: func(session *GameSession) func() bool { return session.rightNoFill },
			expectedBoard: [4][4]uint{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 4, 4},
			},
		},
		{
			name: "do nothing",
			board: [4][4]uint{
				{2, 4, 8, 16},
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{0, 0, 0, 16},
			},
			move: func(session *GameSession) func() bool { return session.rightNoFill },
			expectedBoard: [4][4]uint{
				{2, 4, 8, 16},
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{0, 0, 0, 16},
			},
		},
	}

	runShiftTests(t, tests)
}

type shiftTest struct {
	name          string
	board         [4][4]uint
	move          func(*GameSession) func() bool
	expectedBoard [4][4]uint
}

func runShiftTests(t *testing.T, tests []shiftTest) {
	for _, test := range tests {
		session := &GameSession{
			GameOver:  false,
			GameBoard: test.board,
		}
		t.Run(test.name, func(t *testing.T) {
			test.move(session)()
			if !reflect.DeepEqual(session.GameBoard, test.expectedBoard) {
				t.Fatalf("Incorrect board:\nExpected:\n%s\nActual:  \n%s",
					formatBoard(test.expectedBoard), formatBoard(session.GameBoard))
			}
		})
	}
}

func formatBoard(board [4][4]uint) string {
	var buffer strings.Builder
	for rowIndex, row := range board {
		if rowIndex != 0 {
			buffer.WriteRune('\n')
		}

		for cellIndex, cell := range row {
			if cellIndex != 0 {
				buffer.WriteRune(',')
			}

			buffer.WriteString(strconv.FormatUint(uint64(cell), 10))
		}
	}

	return buffer.String()
}
