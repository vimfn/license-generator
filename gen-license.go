package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type License struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the LICENSE generator!")
	fmt.Println("Please select a license:")
	fmt.Println("1. MIT")
	fmt.Println("2. Apache 2.0")
	fmt.Println("3. GPL 3.0")
	fmt.Print("Enter your choice: ")

	choiceStr, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(choiceStr[:len(choiceStr)-1])
	if choice < 1 || choice > 3 {
		fmt.Println("Invalid choice")
		return
	}

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]

	fmt.Print("Enter the copyright year: ")
	yearStr, _ := reader.ReadString('\n')

	fmt.Printf("Generating LICENSE for %s...\n", name)

	var licenseUrl string
	switch choice {
	case 1:
		licenseUrl = "https://api.github.com/licenses/mit"
	case 2:
		licenseUrl = "https://api.github.com/licenses/apache-2.0"
	case 3:
		licenseUrl = "https://api.github.com/licenses/gpl-3.0"
	}

	resp, err := http.Get(licenseUrl)
	if err != nil {
		fmt.Println("Error fetching license:", err)
		return
	}
	defer resp.Body.Close()

	var license License
	if err := json.NewDecoder(resp.Body).Decode(&license); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	license.Body = strings.ReplaceAll(license.Body, "[year]", yearStr)
	license.Body = strings.ReplaceAll(license.Body, "[yyyy]", yearStr)
	license.Body = strings.ReplaceAll(license.Body, "<year>", yearStr)

	license.Body = strings.ReplaceAll(license.Body, "[fullname]", name)
	license.Body = strings.ReplaceAll(license.Body, "[name of copyright owner]", name)
	license.Body = strings.ReplaceAll(license.Body, "<name of author>", name)

	f, err := os.Create("LICENSE")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(license.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Done!")
}
