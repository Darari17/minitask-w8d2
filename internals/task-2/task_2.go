package task2

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

func (p *Person) Print() {
	fmt.Printf("Nama: %s\n", p.Name)
	fmt.Printf("Address: %s\n", p.Address)
	fmt.Printf("Phone: %s\n", p.Phone)
}

func (p *Person) Greet() {
	fmt.Printf("Selamat pagi %s\n", p.Name)
}

func (p *Person) ChangeName(name string) {
	p.Name = name
}

func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}

}
