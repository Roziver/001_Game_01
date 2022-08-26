package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func finishFunc() {

	fmt.Println("\nПопробовать еще раз? Если Да нажмите y если Нет нажмите любую другую клавишу")
	yes := "y"

	var restart string
	fmt.Scanf("%s\n", &restart)

	if restart == (yes) {
		fmt.Println("\nНачнем игру сначала!")
	} else {
		fmt.Print("\nДо следущей встречи!", "\n", "\n")
		os.Exit(0)
	}
}

func main() {

	success := true
	for game := 0; game <= 100; game++ { // цикл повтора всей игры

		//Генерируем случайное число
		second := time.Now().Unix()  // берем значение времени, переводим его в целое число
		rand.Seed(second)            // генерируем слачайное число на основе меняющейся переменной second
		target := rand.Intn(100) + 1 // берем из случайных чисел интервал от 1 до 100

		fmt.Println(target)

		fmt.Println("\nЯ загадал число от 1 до 100.", "\nСможете угадать?")

		for attempts := 0; attempts <= 9; attempts++ {
			fmt.Println("У вас", 10-attempts, "попыток.")

			fmt.Println("\nВведите ваше число => ")

			var input int
			fmt.Scanf("%d\n", &input)

			if input < target {
				fmt.Println("Ваше число меньше того, что я загадал.")
			} else if input > target {
				fmt.Println("Ваше число больше загаданного.")
			} else {
				success = false
				fmt.Println("\nУра! Поздравляем, вы угадали!")
				finishFunc()
				break
			}
		}
		if success {
			fmt.Println("\nК сожалению вы проиграли.", "\nЗагаданное число", target)
			finishFunc()
		}
	}
}
