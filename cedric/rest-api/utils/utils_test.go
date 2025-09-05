package utils

import "testing"

func TestHelp(t *testing.T) {
	if Help() != "Help!" {
		t.Errorf("Expected 'Help!', got '%s'", Help())
	}
}
