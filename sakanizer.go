package main

import (
	"math/rand"
	"strings"
)

// Sakanize applies transformations to a message based on attributes.
func Sakanize(message, attribute, name string) string {
	switch {
	case attribute == "gopher" && name == "sakanakun":
		ss := []string{"Go", "ぎょ"}
		rand.Shuffle(len(ss), func(i, j int) { ss[i], ss[j] = ss[j], ss[i] })

		cnt := strings.Count(message, "ご")
		if cnt > 0 {
			for i := 0; i < cnt; i++ {
				message = strings.Replace(message, "ご", ss[i%len(ss)], 1)
			}
		}
		return message
	case attribute == "gopher":
		return strings.ReplaceAll(message, "ご", "Go")
	case name == "sakanakun":
		return strings.ReplaceAll(message, "ご", "ギョ")
	default:
		return message
	}
}
