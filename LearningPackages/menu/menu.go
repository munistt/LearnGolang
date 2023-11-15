package menu

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

type menu []menuItem

var in = bufio.NewReader(os.Stdin)

func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", len(item.name)))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}

func (m *menu) add() {
	fmt.Println("Please enter the name of item.")
	itemName, _ := in.ReadString('\n')
	*m = append(*m, menuItem{name: itemName, prices: make(map[string]float64)})
}

func Add() {
	data.add()
}

func Print() {
	data.print()
}
