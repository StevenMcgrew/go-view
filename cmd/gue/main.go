package main

import (
	"fmt"
	"go-view/internal/gue"
)

func main() {

	tz := gue.NewTokenizer(`let something = 55`)
	for {
		t := tz.GetNextToken()
		if t.Type == gue.Eof {
			fmt.Println("ALL DONE!")
			return
		}
		fmt.Println(t)
	}
}
