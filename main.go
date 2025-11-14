package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func PaswordGenaretor(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" + "!@#$%^&*()_-+=<>?/{}~|"

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

func main() {

	var length int
	fmt.Print("Enter password length: ")
	fmt.Scan(&length)

	randomp := PaswordGenaretor(length)
	fmt.Println("Random Password is:", randomp)
}

//saikat
