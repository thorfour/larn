package io

import (
	"os"
	"testing"
)

// Ensures NewGame files can be created
func TestSmokeNewGame(t *testing.T) {

	f, err := NewGame()
	if err != nil {
		t.Fatalf("failed to create new game file: %v", err)
	}

	t.Logf("Created save file %s", f)

	if err := os.Remove(f); err != nil {
		t.Fatalf("failed to remove game file: %v", err)
	}
}
