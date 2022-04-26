package main

import "fmt"

// Human представляет собой структуру в качестве родительской структуры.
type Human struct {
	food string
}

// HumanInfo возвращает значения полей human.
func (human Human) HumanInfo() string {
	return human.food
}

// Action представляет собой структуру в качестве child структуры.
type Action struct {
	h Human
}

func main() {
	// Создание экземпляра.
	a := Action{
		// child структура может напрямую обращаться к переменным базовой структуры.
		Human{
			food: "meat",
		},
	}

	// Вывод базового метода с исопльзованием child.
	fmt.Println("Food is: ", a.h.HumanInfo())
}
