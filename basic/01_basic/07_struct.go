package main

import "fmt"

func main() {
	bird := *NewBird("小鸟", 1, "公", "绿色")
	fmt.Println(bird)
	bird.Fly()
	bird.CommonAnimal.Eat()

	dog := *NewDog2("小狗", 2, "公", "黄色")
	fmt.Println(dog)
	dog.Bark()
	dog.CommonAnimal.Eat()
}

// Go语言中继承，是通过结构体的嵌套来实现的。

type Animal struct {
	Name   string
	Age    int
	Gender string
}

func (a *Animal) Eat() {
	fmt.Println(a.Name, "我要吃到饱！")
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
	fmt.Println("我起飞了！")
}

type Dog02 struct {
	Color        string
	CommonAnimal Animal
}

func NewDog2(name string, age int, gender string, color string) *Dog02 {
	return &Dog02{
		Color: color,
		CommonAnimal: Animal{
			Name:   name,
			Age:    age,
			Gender: gender,
		},
	}
}

func (d *Dog02) Bark() {
	fmt.Println("狗叫声！")
}
