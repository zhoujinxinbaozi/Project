package entity

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	var cat, dog Animal
	cat = &Cat{}
	dog = &Dog{}
	fmt.Printf("cat call : %v\n", cat.Call(""))
	fmt.Printf("dog call : %v\n", dog.Call(""))
}
