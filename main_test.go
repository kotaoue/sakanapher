package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"testing"
)

const testMessage = "おはようございます。ごきげんよろしゅうございますか？"

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

func TestMain_Fixed(t *testing.T) {
	cases := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "no flags",
			args: []string{"--message=" + testMessage},
			want: "おはようございます。ごきげんよろしゅうございますか？",
		},
		{
			name: "gopher",
			args: []string{"--attribute=gopher", "--message=" + testMessage},
			want: "おはようGoざいます。GoきげんよろしゅうGoざいますか？",
		},
		{
			name: "sakanakun",
			args: []string{"--name=sakanakun", "--message=" + testMessage},
			want: "おはようギョざいます。ギョきげんよろしゅうギョざいますか？",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flag variables to their zero value before each run
			attribute = ""
			name = ""

			rootCmd.SetArgs(tt.args)

			got := captureOutput(Execute)

			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}

func TestMain_Random(t *testing.T) {
	wants := []string{
		"おはようGoざいます。ぎょきげんよろしゅうGoざいますか？",
		"おはようぎょざいます。Goきげんよろしゅうぎょざいますか？",
	}

	args := []string{"--attribute=gopher", "--name=sakanakun", "--message=" + testMessage}

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("attempt_%02d", i+1), func(t *testing.T) {
			// Reset flag variables to their zero value before each run
			attribute = ""
			name = ""

			rootCmd.SetArgs(args)
			got := captureOutput(Execute)

			if !slices.Contains(wants, got) {
				t.Errorf("got %q, want one of %v", got, wants)
			}
		})
	}
}
