package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	success := true
	for game := 0; game <= 100; game++ { // цикл повтора всей игры

		//Генерируем случайное число
		second := time.Now().Unix()  // берем значение времени, переводим его в целое число
		rand.Seed(second)            // генерируем слачайное число на основе меняющейся переменной second
		target := rand.Intn(100) + 1 // берем из случайных чисел интервал от 1 до 100

		fmt.Println(target)

		fmt.Println("\nI guessed a number from 1 to 100.", "\nCan you guess?")

		for attempts := 0; attempts <= 9; attempts++ {
			fmt.Println("You have", 10-attempts, "attempts.")

			fmt.Println("\nEnter your number => ")

			var input int
			fmt.Scanf("%d\n", &input)

			if input < target {
				fmt.Println("Your number is less than what I guessed.")
			} else if input > target {
				fmt.Println("Your number is higher than expected.")
			} else {
				success = false
				fmt.Println("\nHooray! Congratulations, you guessed it!")
				finishFunc()
				break
			}
		}
		if success {
			fmt.Println("\nUnfortunately you lost.", "\nHidden number =>", target)
			finishFunc()
		}
	}
}

func finishFunc() {

	fmt.Println("\nTo try one more time? If yes click Y, if No press any other key")
	yes := "Y"

	var restart string
	fmt.Scanf("%s\n", &restart)

	if restart == (yes) {
		fmt.Println("\nLet's start the game over!")
	} else {
		fmt.Print("\nSee you next time!", "\n", "\n")
		os.Exit(0)
	}
}
