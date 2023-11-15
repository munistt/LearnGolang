package main

import (
	"bufio"
	"demo/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {

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
			menu.Print()

		case "2":
			menu.Add()
		case "q":
			break loop
		default:
			fmt.Println("Invalid option")
		}
	}

}
