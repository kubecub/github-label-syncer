package tests

import (
	"os"
	"testing"
)

func TestGithooks(t *testing.T) {
	_, err := os.Stat("./_output/tools/go-gitlint")
	if os.IsNotExist(err) {
		t.Fatalf("go-gitlint tool not found")
	}

	// Add code here to run the githooks and check their output
}
