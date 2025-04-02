package personaldata

import "fmt"

// Ниже создайте структуру Personal
type Personal struct {
	Name   string
	Weight string
	Height string
}

// Ниже создайте метод Print()
func (p Personal) Print() {
	fmt.Printf("Имя: ???\nВес: ???\nРост: ???\n")
}
