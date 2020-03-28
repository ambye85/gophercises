package main

import (
	"fmt"
	"github.com/ambye85/gophercises/caesar"
	"os"
)

func main() {
	var length, rotations int
	var unencrypted string
	_, err := fmt.Scanf("%d\n%s\n%d", &length, &unencrypted, &rotations)
	exitIfError(err)

	encrypted := caesar.Encrypt(unencrypted, rotations)
	fmt.Println(encrypted)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Println("Error reading input, exiting...")
		os.Exit(1)
	}
}