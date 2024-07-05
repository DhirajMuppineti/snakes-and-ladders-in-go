package board_test

import (
	"main/board"
	"testing"
)

func TestGetBoard(t *testing.T) {
	t.Parallel()
	_ = board.GetBoard()
}

func TestUpdateBoard(t *testing.T) {
	t.Parallel()
	b := board.GetBoard()
	err := b.UpdateBoard(1, 2)

	if err != nil {
		t.Fatal(err)
	}

	wantPlayerOnePos := 1
	wantPlayerTwoPos := 2

	gotPlayerOnePos, gotPlayerTwoPos := b.PlayerOnePos, b.PlayerTwoPos

	if wantPlayerOnePos != gotPlayerOnePos {
		t.Errorf("want %d, got %d", wantPlayerOnePos, gotPlayerOnePos)
	}

	if wantPlayerTwoPos != gotPlayerTwoPos {
		t.Errorf("want %d, got %d", wantPlayerTwoPos, gotPlayerTwoPos)
	}
}

func TestUpdateBoardInvalidInput(t *testing.T) {
	t.Parallel()
	b := board.GetBoard()
	err := b.UpdateBoard(-1, 101)

	if err == nil {
		t.Error("want error with invalid out of bounds inputs, got nil")
	}
}

func TestPrintBoard(t *testing.T) {
	t.Parallel()
	b := board.GetBoard()
	b.PrintBoard()
}
