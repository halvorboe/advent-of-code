package day17

import "testing"

func TestBinary(t *testing.T) {
	rows := merge([]uint{BLOCKED}, PIECES[0], 4)
	if len(rows) != 2 {
		t.Errorf("Expected 2 rows, got %d", len(rows))
	}
	if rows[0] != BLOCKED {
		t.Errorf("Expected %09b, got %09b", BLOCKED, rows[0])
	}
	if rows[1] != 0b100111101 {
		t.Errorf("Expected %09b, got %09b", 0b100111101, rows[1])
	}

	rows = merge(rows, PIECES[1], 6)
	if len(rows) != 5 {
		t.Errorf("Expected 5 rows, got %d", len(rows))
	}
	if rows[0] != BLOCKED {
		t.Errorf("Expected %09b, got %09b", BLOCKED, rows[0])
	}
	if rows[1] != 0b100111101 {
		t.Errorf("Expected %09b, got %09b", 0b100111101, rows[1])
	}
	if rows[2] != 0b100010001 {
		t.Errorf("Expected %09b, got %09b", 0b100010001, rows[2])
	}
	if rows[3] != 0b100111001 {
		t.Errorf("Expected %09b, got %09b", 0b100111001, rows[3])
	}
	if rows[4] != 0b100010001 {
		t.Errorf("Expected %09b, got %09b", 0b100010001, rows[2])
	}

}
