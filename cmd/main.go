package main

import (
	"fmt"
	"strings"

	task1 "github.com/Darari17/minitask-w8d2/internals/task-1"
	task2 "github.com/Darari17/minitask-w8d2/internals/task-2"
	task3 "github.com/Darari17/minitask-w8d2/internals/task-3"
)

func main() {
	fmt.Println(strings.Repeat("=", 100))
	// valid
	if err := task1.ReadFile("internals/data/text.txt"); err != nil {
		fmt.Println(err.Error())
	}

	// ini gagal membuka file, path nya salah
	if err := task1.ReadFile("internals/data/text.tx"); err != nil {
		fmt.Println(err.Error())
	}

	// ini bisa membuka file, tapi yang dibaca directory buka file
	if err := task1.ReadFile("internals/data"); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(strings.Repeat("=", 100))
	person := task2.NewPerson("Farid", "Jakarta", "08123456789")
	person.Print()
	person.Greet()
	fmt.Println(person)

	person.ChangeName("Darari")
	person.Greet()
	fmt.Println(person)

	fmt.Println(strings.Repeat("=", 100))
	pembayaran1 := []int{1000, 2000, 3000}
	pembayaran2 := []int{5000, -1000}
	pembayaran3 := []int{-5000, -1000}

	bank := task3.Bank{Name: "BRI"}
	res, err := bank.Pay(pembayaran1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)

	res, err = bank.Pay(pembayaran2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)

	online := task3.Online{Name: "Dana"}
	res, err = online.Pay(pembayaran1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)

	online.Pay(pembayaran2)
	res, err = online.Pay(pembayaran2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)

	fiktif := &task3.Fiktif{}
	result, err := fiktif.Pay(pembayaran1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	result, err = fiktif.Pay(pembayaran2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

	result, err = fiktif.Pay(pembayaran3)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

	fmt.Println("List Pembayaran Fiktif:", fiktif.Lists)
	for i, v := range fiktif.Lists {
		fmt.Printf("Item %d sebesar %d\n", i+1, v)
	}
}
