package utils

import (
	"fmt"
)

func Students() {

	studentsGrades := map[string][]int{

		"Donald":    {67, 78, 98, 85, -99},
		"Goofy":     {45, -89, 97, 88, 78},
		"Balto":     {89, -56, 89, 45, 34},
		"Smith":     {98, 34, 67, 53, 56},
		"Wonderful": {-78, 71, 99, 89, 98},
		"a":         {},
	}

	for key, studentGrades := range studentsGrades {
		if len(studentGrades) == 0 {
			fmt.Println("This student doesn't have grades")
		} else {
			fmt.Printf("%s: %d\n", key, Max(studentGrades))
		}
	}

}

func Max(studentGrades []int) int {
	maxGrade := studentGrades[0]
	for _, grade := range studentGrades {
		if grade > maxGrade {
			maxGrade = grade
		}
	}
	return maxGrade
}
