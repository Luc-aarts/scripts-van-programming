package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

var description string

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

			a := vals()
			fmt.Println(a)

		}

	}
	cmd.Wait()
}
