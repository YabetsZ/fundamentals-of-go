package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		studentName string
		numberOfSubjects int
	)

	
	for {
		fmt.Print("Enter your name: ")
		fmt.Scanln(&studentName)
		studentName = strings.TrimSpace(studentName)
		if studentName != "" {
			break
		}
		fmt.Println("Invalid input: give a name which is not empty word.")
	}

	for {
		fmt.Print("Enter the number of subjects you are taking: ")
		var input string
		fmt.Scanln(&input)

		num, err := strconv.Atoi(input)
		if err == nil && num > 0 {
			numberOfSubjects = num
			break
		}
		// Error/warning for the invalid input received
		fmt.Println("Invalid input: Please enter a positive integer for the number of subjects.")
	}

	subjectGrades := make(map[string]float64)

	for i := 0; i < numberOfSubjects; i++ {
		fmt.Printf("\n--- Subject %d ---\n", i+1)
		var subjectName string
		for {
			fmt.Print("Enter the name of the subject: ")
			fmt.Scanln(&subjectName)
			subjectName = strings.TrimSpace(subjectName)
			_, ok := subjectGrades[subjectName]
			if subjectName != "" && !ok {
				break
			}
			fmt.Println("Invalid input: give a subject which is not already registered or empty word.")
		}
		
		
		for {
			fmt.Print("Enter the score in this, ", subjectName, ", subject: ")
			var input string
			fmt.Scanln(&input)

			score, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
			if err == nil && score >= 0 && score <= 100 {
				subjectGrades[subjectName] = score
				break
			}
			// Error/warning for the invalid input received
			fmt.Println("Invalid input: Please enter a grade score between 0 and 100")
		}

	}

	averageScore := calculateAvarage(subjectGrades)

	fmt.Println("\n--- Student Report ---")
	fmt.Printf("Student name: %s\n", studentName)
	fmt.Println("Individual grades:")
	for subject, score := range subjectGrades {
		fmt.Printf("%20s: %-6.2f\n", subject, score)
	}
	fmt.Printf("%20s: %-6.2f","✔️ Average Grade", averageScore)

}

// Function to calculate avarage score of the given subject_name:grade map
func calculateAvarage(grades map[string]float64) (average float64) {
	var sumScore float64
	for _, score := range grades {
		sumScore += score
	}

	average = sumScore/float64(len(grades))

	return
}