package generics

func main() {

	a1 := []int{1, 2, 3, 4}
	a2 := []float64{2.3, 4.5, 6.0}
	s1 := []string{"foo", "bar", "baz"}

}

func add(s []int) int {

	result := 0
	for _, v := range s {
		result += v
	}
	return result
}
