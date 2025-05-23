package main

import "fmt"

// Animal is the type for our abstract factory
type Animal interface {
	Says()
	LikesWater() bool
}

// Dog is a concrete factory for dogs

type Dog struct {}

func (d *Dog) Says() {
	fmt.Println("Woof!!!")
}

func (d *Dog) LikesWater() bool {
	return true
}

// Cat is a concrete factory for cats

type Cat struct {}

func (c *Cat) Says() {
	fmt.Println("meow!!!")
}

func (c *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct {}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct {}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	// Create one each of a DogFactor and a CatFactory

	dogFactory := DogFactory{}
	catFactory := CatFactory{}

	// Call the new method to create a dog and a cat

	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	cat.Says()

	fmt.Println("A dog likes water: ", dog.LikesWater())
	fmt.Println("A cat likes water: ", cat.LikesWater())
}