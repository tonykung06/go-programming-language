package main

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type Salutations []Salutation

//salutation is a copy
func (salutation Salutation) Rename(newName string) {
	salutation.Name = newName
}

func Rename2(salutation *Salutation, newName string) {
	salutation.Name = newName
}

func (salutation *Salutation) Rename3(newName string) {
	salutation.Name = newName
}

func (salutations Salutations) Greet(isFormal bool, times int) {
	for _, s := range salutations {
		fmt.Println(s)
	}
}

func main() {
	salutations := Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
	}

	salutations[0].Rename("testing")
	salutations.Greet(true, 4)
	Rename2(&salutations[0], "testing2")
	salutations.Greet(true, 4)
	salutations[0].Rename3("testing3")
	salutations.Greet(true, 4)
}
