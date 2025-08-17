package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// captureOutput executes the given function while capturing its standard output.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = old }()

	f()

	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return strings.TrimSpace(buf.String())
}

func TestMain(t *testing.T) {
	// Detailed patterns are tested with sakanizer_test.go
	rootCmd.SetArgs([]string{"--message=" + "おはようございます。ごきげんよろしゅうございますか？"})
	got := captureOutput(Execute)

	want := "おはようございます。ごきげんよろしゅうございますか？"
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
