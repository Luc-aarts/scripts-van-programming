package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var description string
var macadres string
var ipadres string

func vals() string {
	return description
}

func main() {

	cmd := exec.Command("ipconfig", "/all")
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	cmd.Start()

	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		tekst := scanner.Text()

		if strings.Contains(tekst, "Description") {
			dubbelePuntIndex := strings.Index(tekst, ":")
			if dubbelePuntIndex == -1 {
				continue
			}
			gekinpteTekst := tekst[dubbelePuntIndex+1:]
			gekinpteTekst = strings.Trim(gekinpteTekst, "\t\n")
			haakjeIndex := strings.Index(gekinpteTekst, "(")
			if haakjeIndex != -1 {
				gekinpteTekst = gekinpteTekst[:haakjeIndex]
			}
			description = gekinpteTekst

		} else if strings.Contains(tekst, "Physical Address") {
			dubbelePuntIndex := strings.Index(tekst, "")
			if dubbelePuntIndex == -1 {
				continue
			}
			gekinpteTekst := tekst[dubbelePuntIndex+1:]
			gekinpteTekst = strings.Trim(gekinpteTekst, "\t\n")
			haakjeIndex := strings.Index(gekinpteTekst, "(")
			if haakjeIndex != -1 {
				gekinpteTekst = gekinpteTekst[:haakjeIndex]
			}
			macadres = gekinpteTekst

		} else if strings.Contains(tekst, "IPv4 Address") {
			dubbelePuntIndex := strings.Index(tekst, ":")
			if dubbelePuntIndex == -1 {
				continue
			}
			gekinpteTekst := tekst[dubbelePuntIndex+1:]
			gekinpteTekst = strings.Trim(gekinpteTekst, "\t\n")
			haakjeIndex := strings.Index(gekinpteTekst, "(")
			if haakjeIndex != -1 {
				gekinpteTekst = gekinpteTekst[:haakjeIndex]
			}
			ipadres = gekinpteTekst

			a := vals()
			fmt.Println(a)
		}
	}

	fmt.Println("kies een netwerk kaart")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	tekst := scanner1.Text()
	if tekst != "intel" {
		return
	}

	cmd.Wait()

	fmt.Println("hardware:", description)
	fmt.Println("mac-adres:", macadres)
	fmt.Println("Ip-adres:", ipadres)
}
