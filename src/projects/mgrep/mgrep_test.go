package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMgrepMain(t *testing.T) {
	os.Args = []string{"mgrep", "Summary", "."}
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = stdout

	if !strings.Contains(string(out), "./REQUIREMENTS [1] //--Summary:") {
		t.Errorf("mgrep did not print expected output")
	}
}
