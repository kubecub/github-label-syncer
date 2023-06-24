package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestSyncerE2E(t *testing.T) {
	// Set environment variables
	os.Setenv("INPUT_MANIFEST", "./.github/sync_labels.yml")
	os.Setenv("INPUT_PRUNE", "false")

	// Run syncer command
	cmd := exec.Command("go", "run", "../cmd/syncer/main.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run syncer command: %v", err)
	}

	// Check output
	expectedOutput := "Syncing labels...\nDone!"
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Unexpected output: %s", string(output))
	}
}
