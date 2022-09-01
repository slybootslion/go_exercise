package main

import "fmt"

type Animal struct {
	Name   string
	Age    int
	Gender string
}

func (a *Animal) Eat() {
	fmt.Println(a.Name, "feed!")
}

type Bird struct {
	WingColor    string
	CommonAnimal Animal
}

func NewBird(name string, age int, gender string, wingColor string) *Bird {
	return &Bird{
		WingColor: wingColor,
		CommonAnimal: Animal{
			Name:   name,
			Age:    age,
			Gender: gender,
		},
	}
}

func (b *Bird) Fly() {
	fmt.Println("bird fly!")
}

type Dog1 struct {
	Color        string
	CommonAnimal Animal
}

func NewDog1(name string, age int, gender string, color string) *Dog1 {
	return &Dog1{
		Color: color,
		CommonAnimal: Animal{
			Name:   name,
			Age:    age,
			Gender: gender,
		},
	}
}

func (d *Dog1) Bark() {
	fmt.Println("wangwangwang!")
}

func main() {
	bird := *NewBird("xiaoniao", 1, "m", "green")
	fmt.Println(bird)
	bird.Fly()

	dog := *NewDog1("xiaogou", 2, "m", "yellow")
	fmt.Println(dog)
	dog.Bark()

	dog.CommonAnimal.Eat()
}
