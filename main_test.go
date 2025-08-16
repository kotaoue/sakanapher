package main

import (
	"bytes"
	"flag"
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
		name     string
		attr     string
		username string
		want     string
	}{
		{
			name:     "no flags",
			attr:     "",
			username: "",
			want:     "おはようございます。ごきげんよろしゅうございますか？",
		},
		{
			name:     "gopher",
			attr:     "gopher",
			username: "",
			want:     "おはようGoざいます。GoきげんよろしゅうGoざいますか？",
		},
		{
			name:     "sakanakun",
			attr:     "",
			username: "sakanakun",
			want:     "おはようギョざいます。ギョきげんよろしゅうギョざいますか？",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			flag.Set("attribute", tt.attr)
			flag.Set("name", tt.username)
			flag.Set("message", testMessage)

			got := captureOutput(func() { Main() })

			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}

func TestMain_Random(t *testing.T) {
	flag.Set("attribute", "gopher")
	flag.Set("name", "sakanakun")
	flag.Set("message", testMessage)

	wants := []string{
		"おはようGoざいます。ぎょきげんよろしゅうGoざいますか？",
		"おはようぎょざいます。Goきげんよろしゅうぎょざいますか？",
	}

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("attempt_%02d", i+1), func(t *testing.T) {
			got := captureOutput(func() { Main() })

			if !slices.Contains(wants, got) {
				t.Errorf("got %q, want one of %v", got, wants)
			}
		})
	}
}
