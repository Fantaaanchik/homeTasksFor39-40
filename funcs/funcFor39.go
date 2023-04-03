package funcs

import (
	"fmt"
	"homeTasksFor39-40/inputOutput"
)

func OutputFor39() {
	inputOutput.InputOutput39()

	for i := inputOutput.A; i <= inputOutput.B; i++ {
		for j := inputOutput.A; j <= i; j++ {
			fmt.Print(i)
			fmt.Print(", ")
		}

	}

	return
}
