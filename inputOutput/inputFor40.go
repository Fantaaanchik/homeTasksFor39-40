package inputOutput

import "fmt"

var (
	A40 int16
	B40 int16
)

func InputOut40() {

	fmt.Printf("For40. Даны целые числа A и B (A < B). Вывести все целые числа от A до B включительно; при этом " +
		"число A должно выводиться 1 раз, число A + 1 должно выводиться 2 раза и т. д. \n")
	_, err := fmt.Scan(&A40, &B40)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
