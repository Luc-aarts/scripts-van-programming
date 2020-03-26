package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

func main() {
	fmt.Println("type wachtwoord voor het laten zien van je wachtwoord")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tekst := scanner.Text()
	if tekst != "wachtwoord" {
		return
	}

	length := flag.Int("l", 15, "Password length")
	quantity := flag.Int("q", 15, "quantity")
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*-"

	flag.Parse()
	for i := 0; i < *quantity; i++ {
		random, err := generatePassword(*length, characters)
		if err != nil {
			panic(err)
		}
		wachtwoord := random
		fmt.Println(random)
		data := []byte(wachtwoord)
		ioutil.WriteFile("./output.txt", data, os.ModePerm)
	}
}
func generatePassword(length int, symbols string) (string, error) {
	result := ""
	for {
		if len(result) >= length {
			return result, nil
		}
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		s := string(num.Int64())
		if strings.Contains(symbols, s) {
			result += s
		}
	}
}
