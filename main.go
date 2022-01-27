package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	attr    = flag.String("attribute", "", "your attribute [gopher|other]")
	name    = flag.String("name", "", "your name")
	message = flag.String("message", "おはようございます", "message of wish to convert")
)

func init() {
	flag.Parse()
}

func main() {
	_, err := Main()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() (int, error) {
	switch {
	case *attr == "gopher" && *name == "sakanakun":
		ss := []string{"Go", "ぎょ"}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(ss), func(i, j int) { ss[i], ss[j] = ss[j], ss[i] })

		cnt := strings.Count(*message, "ご")
		if cnt > 0 {
			for i := 0; i < cnt; i++ {
				*message = strings.Replace(*message, "ご", ss[i%len(ss)], 1)
			}
		}
		return fmt.Println(*message)
	case *attr == "gopher":
		return fmt.Println(strings.ReplaceAll(*message, "ご", "Go"))
	case *name == "sakanakun":
		return fmt.Println(strings.ReplaceAll(*message, "ご", "ギョ"))
	}
	return fmt.Println(*message)
}
