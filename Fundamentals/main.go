package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//	func main() {
//		fmt.Println("Munish Kumar")
//	}
func main() {

	// var a int
	// a = 10

	// fmt.Println(a)

	// var b string = "Munish"
	// fmt.Println(b)

	// c := 3.23
	// d := int(c)
	// fmt.Println(d)

	// r int8 := 54
	// s := r
	// fmt.Println(s)

	// var num int8 = 32
	// var num1 int16 = int16(num)
	// fmt.Println(num1)

	// varA, varB := 10, 20

	// varC := varA + varB
	// fmt.Println((varC))

	// fmt.Println(varA % varB)
	// fmt.Println(float32(varA) / float32(varB))
	// fmt.Println(varA == varB)
	// fmt.Println(varA > varB)

	// const a = 43
	// var i int = a

	// const b int = 32
	// const c int = b

	// const f32 float32 = 3.13
	// const f4 float64 = float64(f32)

	// fmt.Println(f4)
	// fmt.Println(i)
	// fmt.Println(c)

	// const a = iota
	// fmt.Println(a)

	// const (
	// 	b = 2 * 4
	// 	c
	// 	d = iota

	// 	e = 10 * d * iota
	// )
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(e)

	// s := "Hello, Munish"

	// p := &s

	// fmt.Println(p)
	// fmt.Println(*p)

	// *p = "Hello Ajanya"

	// // a := *p
	// fmt.Println(*p)

	// p = new(string)
	// fmt.Println(*p)

	// CLI Application

	// fmt.Println("Type anything you want !")
	// in := bufio.NewReader(os.Stdin)
	// s, _ := in.ReadString('\n')
	// s = strings.TrimSpace(s)
	// s = strings.ToUpper(s)
	// fmt.Println(s + "!")

	// fmt.Println("Enter your name: ")
	// in := bufio.NewReader(os.Stdin)
	// s, _ := in.ReadString('\n')
	// s = strings.ToLower(s)
	// fmt.Println(s)

	// http.HandleFunc("/", Handler)
	// http.ListenAndServe(":3000", nil)

	// Arrays
	// var arr [5]int
	// fmt.Println(arr)

	// arr = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// arr1 := arr
	// fmt.Println(arr[4])

	// fmt.Println(len(arr))

	// arrString := [3]string{"a", "b", "c"}

	// fmt.Println(arrString)

	// fmt.Println(arr == arr1)
	// fmt.Println(reflect.TypeOf(arrString))

	// var s []int

	// fmt.Println(s)

	// s = []int{1, 2, 3, 4, 5}
	// fmt.Println(s)

	// s = append(s, 10, 11, 12)
	// fmt.Println(s)

	// slices.Delete(s, 1, 3)
	// fmt.Println(s)

	// Maps

	// var m map[string]int
	// fmt.Println(m)

	// m = map[string]int{"a": 1, "b": 2}
	// fmt.Println(m)

	// fmt.Println(m["a"])

	// m["a"] = 4
	// fmt.Println(m)

	// delete(m, "a")

	// fmt.Println(m)

	// fmt.Println(m["a"])

	// v, ok := m["a"]
	// fmt.Println(v, ok)

	// m1 := m

	// fmt.Println(m1)

	// fmt.Println(m1 == m)

	// var m map[string][]string
	// fmt.Println(m)

	// m = map[string][]string{
	// 	"cse": {"Munish", "Akshay", "Sahil", "Raman", "Ajanya"},
	// 	"ece": {"Rohit", "Menakshi", "Rahul"},
	// }
	// fmt.Println(m)

	// m["other"] = []string{"Ajju", "Yuhi", "Ajanya"}

	// fmt.Println(m)
	// fmt.Println(m["other"])

	// delete(m, "cse")
	// fmt.Println(m)

	// fmt.Println(m["cse"])

	// v, ok := m["cse"]
	// fmt.Println(v, ok)

	// struct study

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "coffee", prices: map[string]float64{"small": 10.0, "medium": 25.0, "large": 50.0}},
		{name: "Espresso", prices: map[string]float64{"small": 30.0, "medium": 60.0, "large": 100.0}},
	}

	in := bufio.NewReader(os.Stdin)

loop:
	for {
		fmt.Println("Please enter an option: ")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add item to menu")
		fmt.Println("q) Quit! ")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":

			for _, item := range menu {

				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {

					fmt.Printf("\t%10s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Please enter the name of new item.")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}

	// fmt.Println(menu)

	// i := 99

	// for {
	// 	i += 1
	// 	fmt.Println(i)
	// 	break
	// }
	// i := 0
	// for i < 10 {
	// 	i += 1
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// if i := 10; i < 10 {
	// 	fmt.Println("i is less than 10")
	// } else if i < 5 {
	// 	fmt.Println("i is less than 5")
	// } else {
	// 	fmt.Println("i is greater than or equal to 10")
	// }
	// fmt.Println("after the if statement")

	// i := 10
	// switch i {
	// case 1:
	// 	fmt.Println("first case")
	// case 5 + 0, 5*i + 10:
	// 	fmt.Println("second case")
	// case 2:
	// 	fmt.Println("third case")
	// case 10:
	// 	fmt.Println("fourth case")
	// default:
	// 	fmt.Println("default case")
	// }

	// defer function
	// fmt.Println("Hwy1")

	// defer fmt.Println("last1")
	// fmt.Println("hwy2")
	// defer fmt.Println("last2")
	// fmt.Println(choice)

	// panic and recover

	// fmt.Println("main1")
	// fmt.Println("main2")
	// func1()
	// fmt.Println("Hello reached end of line")

	// name, otherName := "Munish", "Muni"
	// fmt.Println(name)
	// fmt.Println(otherName)
	// myFunc(name, &otherName)
	// fmt.Println(name)
	// fmt.Println(otherName)

}

// func func1() {
// 	defer func() {
// 		fmt.Println(recover())
// 	}()
// 	fmt.Println("func1")
// 	panic("program crashed")
// 	// fmt.Println("Program running nicely")
// }

// func Handler(w http.ResponseWriter, r *http.Request) {

// 	f, _ := os.Open("./menu.txt")
// 	io.Copy(w, f)

// }

// func myFunc(name string, otherName *string) {
// 	name = "Ashish"
// 	*otherName = "Shunu"
// }
