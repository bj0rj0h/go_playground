package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCorrectInput(t *testing.T) {

	subproc := exec.Command("go_playground")
	input := "1\n"
	reader := strings.NewReader(input)
	subproc.Stdin = reader
	output, _ := subproc.Output()

	s := string(output)
	if "one\n" != s {
		t.Errorf("Wanted: %s, Got: %s", "one", s)
	}

	subproc.Process.Kill()
}
