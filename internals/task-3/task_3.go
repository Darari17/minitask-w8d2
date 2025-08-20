package task3

import (
	"errors"
	"fmt"
)

type IPayment interface {
	Pay(ammount []int) (string, error)
}

type Bank struct {
	Name string
}

type Online struct {
	Name string
}

type Fiktif struct {
	Lists []int
}

func (b *Bank) Pay(ammount []int) (string, error) {
	total := 0
	for _, v := range ammount {
		total += v
	}
	return fmt.Sprintf("Pembayaran %d berhasil via Bank %s\n", total, b.Name), nil
}

func (o *Online) Pay(ammount []int) (string, error) {
	total := 0
	for _, v := range ammount {
		total += v
	}
	return fmt.Sprintf("Pembayaran %d berhasil via Online %s\n", total, o.Name), nil
}

func (f *Fiktif) Pay(ammount []int) (string, error) {
	total := 0
	for _, v := range ammount {
		total += v
	}
	if total <= 0 {
		return "", errors.New("total pembayaran tidak valid (pembayaran <= 0)")
	}
	f.Lists = append(f.Lists, total)
	return fmt.Sprintf("Pembayaran %d berhasil\n", total), nil
}
