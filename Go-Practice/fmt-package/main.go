package main

import (
	"fmt"
	"math"
)

func main() {

	integer := 42
	fmt.Print(integer)
	// fmt.Println()
	fmt.Println(integer)

	// name := "Munish"

	// fmt.Printf("Name : %s\n", name)

	// %v -> the value in a default format
	fmt.Printf("%v", integer)
	fmt.Println()

	truth := true
	// %t is the format for boolean value true or false
	fmt.Printf("%v %t\n", truth, truth)

	// %T is for type of the value
	// %#v is go-syntax representation of value (prints same representation as go-syntax)
	// %% a literal percent sign, it consumes no value
	fmt.Printf("%T\n", integer)  // op: int
	fmt.Printf("%#v\n", integer) //op: 42
	fmt.Printf("%%\n")           // returns %

	/*
		%d for decimal values
		%x for hexadecimal values
		%o for octal
		%b for binary
	*/
	fmt.Printf("%v %d %x %o %b\n", integer, integer, integer, integer, integer)
	// Result: 42 42 2a 52 101010

	/*
		%v and %g prints the compact representation
		%f prints the decimal point
		%e uses exponential representation
		%6.2f here shows the width and precision to control the apprearance of floating point value
		6 -> total 6 spaces width and .2 -> is the decimal after point
	*/
	p := math.Pi

	fmt.Printf("%v %g %f %.2f (%6.2f) %e\n", p, p, p, p, p, p)
	// Result : 3.141592653589793 3.141592653589793 3.141593 3.14 (  3.14) 3.141593e+00

	// complex numbers are formated as paranthesized pairs of floats with imaginary parts

	point := 11.7 + 90.7i
	fmt.Printf("%.1f %v\n", point, point)
	//  Result:(11.7+90.7i) (11.7+90.7i)

	/*
		rune type in go is alias type of int32 bcz go uses UTF-* encoding
		characters can be defined using runes
		-------------------------------------------
		https://golangdocs.com/rune-in-golang
		-------------------------------------------
		%c shows the character with the unicode value
		%q shows them as quoted characters
		%U shows them as hex Unicode point
		%#U shows them in both code point and quoted form it rune is printable

	*/

	character := 'a'
	fmt.Printf("%v %c %q %U %#U\n", character, character, character, character, character)

	// strings are formatted with %v and %s, with %q as quoted strings
	// %#q as backquoted strings

	name := `Munish "Kumar"`
	fmt.Printf("%v %s %q %#q\n", name, name, name, name)

}
