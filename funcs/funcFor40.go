package funcs

import (
	"fmt"
	"homeTasksFor39-40/inputOutput"
)

func OutputFor40() {
	inputOutput.InputOut40()
	if inputOutput.A40 > inputOutput.B40 {
		fmt.Println("А превосходит В")
	} else {
		for i := inputOutput.A40; i <= inputOutput.B40; i++ {
			for j := inputOutput.A40; j <= i; j++ {
				fmt.Print(i)
				fmt.Print(", ")
			}
		}
	}

	return
}
