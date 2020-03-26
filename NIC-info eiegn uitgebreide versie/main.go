package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const macadreslengte = 17

var ipadres string
var macadres string
var hostname string
var description string
var subnetmask string

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

		if strings.Contains(tekst, "IPv4 Address") {
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

		} else if strings.Contains(tekst, "Physical Address") {
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
			macadres = gekinpteTekst

		} else if strings.Contains(tekst, "Host Name") {
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
			hostname = gekinpteTekst

		} else if strings.Contains(tekst, "Description") {
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

		} else if strings.Contains(tekst, "Subnet Mask") {
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
			subnetmask = gekinpteTekst

		}
	}
	cmd.Wait()

	currentTime := time.Now()
	fmt.Println("tijd van aanvraag")
	fmt.Println(currentTime.Format("2006-01-02 15:04:05 Thursday"))

	fmt.Println("je Host Name is:", hostname)
	fmt.Println("je ip-adres is:", ipadres)
	fmt.Println("je subnetmask:", subnetmask)
	fmt.Println("je mac-adres is:", macadres)
	fmt.Println("je Description:", description)

}
