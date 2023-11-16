package main

import "fmt"

func main() {

	// testScores := []float64{23.0, 99.9, 87.4, 23.88}

	// cloneOfScore := cloneTestScore(testScores)

	// fmt.Println(&cloneOfScore[0], &testScores[0], cloneOfScore)

	studentScores := map[string]float64{
		"Munish Kumar": 84,
		"Sahil Singh":  85,
		"Roshni":       87,
		"Romesh":       90,
		"Ajanya":       91,
	}

	clonedStudentScores := cloneStudentScores(studentScores)

	// fmt.Println(&studentScores[0], &clonedStudentScores[0], clonedStudentScores)

	fmt.Println(clonedStudentScores)

}

//using any, comparable constraint

func cloneStudentScores[K comparable, V any](score map[K]V) map[K]V {

	result := make(map[K]V, len(score))

	for k, v := range score {
		result[k] = v
	}
	return result

}

// func cloneStudentScores(scores map[string]float64) map[string]float64 {
// 	result := make(map[string]float64, len(scores))

// 	for k, v := range scores {
// 		result[k] = v
// 	}
// 	return result
// }

// any type contraint

// func cloneTestScore[V any](s []V) []V {
// 	result := make([]V, len(s))

// 	for i, v := range s {
// 		result[i] = v
// 	}
// 	return result
// }

// func cloneTestScore(s []float64) []float64 {

// 	result := make([]float64, len(s))

// 	// for i, v := range s {
// 	// 	result[i] = v
// 	// }
// 	copy(result, s)
// 	return result

// }
