package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)

	user1 := myuser{name, age}

	fmt.Println("Name: " + user1.name)

	fmt.Fscan(os.Stdin, &name)
}
