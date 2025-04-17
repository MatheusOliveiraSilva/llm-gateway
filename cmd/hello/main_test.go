package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestCLI_DefaultName(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		t.Fatalf("erro ao executar CLI: %v", err)
	}

	got := strings.TrimSpace(out.String())
	want := "Oi, Assistant! 🚀"

	if got != want {
		t.Errorf("esperava %q, obtive %q", want, got)
	}
}
