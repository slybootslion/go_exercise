package main

import "fmt"

func createPlayer(name string, career string, gender string) func() (string, string, string, int, int) {
	var hp = 0
	var mp = 0
	if career == "zs" {
		hp = 150
		mp = 100
	} else if career == "fs" {
		hp = 100
		mp = 200
	}

	return func() (string, string, string, int, int) {
		return name, career, gender, hp, mp
	}
}

func main() {
	playerA := createPlayer("kztx", "zs", "m")
	nameA, creerA, genderA, hpA, mpA := playerA()
	fmt.Println(genderA, creerA, nameA, hpA, mpA)
	playerB := createPlayer("wyyl", "fs", "f")
	nameB, creerB, genderB, hpB, mpB := playerB()
	fmt.Println(genderB, creerB, nameB, hpB, mpB)

	mpB -= 15
	hpA -= 20
	fmt.Println(nameB, hpB, mpB)
	fmt.Println(nameA, hpA, mpA)
}
