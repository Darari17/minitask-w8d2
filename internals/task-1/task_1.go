package task1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("gagal membuka file")
	}
	defer file.Close()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		fmt.Println(value)
	}
	if err := scanner.Err(); err != nil {
		panic("continue")
	}
	return nil
}
