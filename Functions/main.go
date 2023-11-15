package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var menu []menuItem

var in = bufio.NewReader(os.Stdin)

func main() {

	menu = []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 10.0, "medium": 20.0, "large": 40.0}},
	}
loop:
	for {

		fmt.Println("Please enter an option: ")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add new item")
		fmt.Println("q) Quit")

		s, _ := in.ReadString('\n')
		choice := strings.TrimSpace(s)

		// fmt.Println(menu)

		switch choice {
		case "1":
			printMenu()

		case "2":
			addItem()
		case "q":
			break loop
		default:
			fmt.Println("Invalid option")
		}
	}

	// fmt.Println(s)

}

func addItem() {
	fmt.Println("Please enter the name of item.")
	itemName, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: itemName, prices: make(map[string]float64)})
}

func printMenu() {

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", len(item.name)))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}
