package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	cases := []struct {
		name     string
		attr     string
		username string
		message  string
		expected string
	}{
		{
			name:     "gopher and sakanakun",
			attr:     "gopher",
			username: "sakanakun",
			message:  "ごはんですよ",
			expected: "Go",
		},
		{
			name:     "gopher",
			attr:     "gopher",
			username: "",
			message:  "ごはんですよ",
			expected: "Go",
		},
		{
			name:     "sakanakun",
			attr:     "",
			username: "sakanakun",
			message:  "ごはんですよ",
			expected: "ギョ",
		},
		{
			name:     "no attribute",
			attr:     "",
			username: "",
			message:  "ごはんですよ",
			expected: "ご",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			flag.Set("attribute", tt.attr)
			flag.Set("name", tt.username)
			flag.Set("message", tt.message)

			// redirect stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			Main()

			// restore stdout
			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)

			if !strings.Contains(buf.String(), tt.expected) {
				t.Errorf("expected %s, but got %s", tt.expected, buf.String())
			}
		})
	}
}
