// l1.1
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
// Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.

package main

import "fmt"

// Родительская структура
type Human struct {
	Name string
	Age  int
}

// Методы Human
func (h Human) SayName() {
	fmt.Printf("Привет! Меня зовут %s.\n", h.Name)
}

func (h Human) SayAge() {
	fmt.Printf("Мне %d лет.\n", h.Age)
}

// Структура Action встраивает Human (embedded struct)
type Action struct {
	Human
	Activity string
}

func main() {
	// Создаём экземпляр Action
	person := Action{
		Human: Human{
			Name: "Макс",
			Age:  18},
		Activity: "бежит",
	}
	// Вызываем методы Human напрямую через Action
	person.SayName()
	person.SayAge()
	// Поле Action
	fmt.Println(person.Activity)
}
