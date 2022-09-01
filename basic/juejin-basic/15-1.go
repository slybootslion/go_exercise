package main

import "fmt"

type Dog struct {
	Breed  string
	Age    int
	Weight float64
	Gender string
}

func NewDog(breed string, age int, weigth float64, gender string) *Dog {
	return &Dog{
		Breed:  breed,
		Age:    age,
		Weight: weigth,
		Gender: gender,
	}
}

func (d *Dog) Sport() {
	fmt.Println("sport")
	d.Weight -= 0.1
	fmt.Println("减重到了：", d.Weight)
}

func (d *Dog) Eat() {
	fmt.Println("eat")
	d.Weight += 0.1
	fmt.Println("我增重到了", d.Weight)
}

func main() {
	fatShibaInu := NewDog("shiba inu", 2, 12.0, "公")
	weakShiabInu := NewDog("shiba inu", 2, 7.0, "公")
	fatShibaInu.Sport()
	weakShiabInu.Eat()
}
