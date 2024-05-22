package main

import (
	"bufio"
	"fmt"
	"os"
)

func readAll(path string) error {
	dat, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	fmt.Print(string(dat))
	return nil
}

func readByLine(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Failed to open %v: %w", path, err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("> %v\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Failed to read %v: %w", path, err)
	}
	return nil
}

func main() {
	err := readAll("/tmp/a.txt")
	//err := readByLine("/tmp/a.txt")
	if err != nil {
		panic(err)
	}
}
