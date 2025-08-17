package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSakanize_Fixed(t *testing.T) {
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
			got := Sakanize("おはようございます。ごきげんよろしゅうございますか？", tt.attr, tt.username)
			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}

func TestSakanize_Random(t *testing.T) {
	wants := []string{
		"おはようGoざいます。ぎょきげんよろしゅうGoざいますか？",
		"おはようぎょざいます。Goきげんよろしゅうぎょざいますか？",
	}

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("attempt_%02d", i+1), func(t *testing.T) {
			got := Sakanize("おはようございます。ごきげんよろしゅうございますか？", "gopher", "sakanakun")
			if !slices.Contains(wants, got) {
				t.Errorf("got %q, want one of %v", got, wants)
			}
		})
	}
}
