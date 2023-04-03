package inputOutput

import "fmt"

var (
	A uint16
	B uint16
)

func InputOutput39() {

	fmt.Printf("For39. Даны целые положительные числа A и B (A < B). Вывести все целые числа от A до B включительно; " +
		"при этом каждое число должно выводиться столько раз, каково его значение (например, число 3 выводится 3 раза).\n")
	fmt.Printf("Введите целые положительные числа А и В:\n")
	_, err := fmt.Scan(&A, &B)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
