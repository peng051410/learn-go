package struct_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("pet...")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(" ", name)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("tomyli")
}
