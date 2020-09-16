package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/gookit/color"
)

func randomCase(l string) string {
	rand := rand.Intn(42)
	flip := rand % 2
	if flip == 0 {
		return strings.ToUpper(l)
	}
	return strings.ToLower(l)
}

func main() {
	if len(os.Args) < 2 {
		color.Error.Println("Error: a string is required.")
		os.Exit(1)
	}
	str := os.Args[1]
	letters := strings.Split(str, "")
	rand.Seed(time.Now().UnixNano())
	output := ""
	for _, l := range letters {
		output = output + randomCase(l)
	}
	clipboard.WriteAll(output)
	fmt.Println(output)
}
