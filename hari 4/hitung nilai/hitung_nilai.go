package main

import (
	"fmt"
	"math"
)

// Struct untuk Student
type Student struct {
	Name  string
	Score float64
}

func main() {
	// Membuat slice untuk menyimpan data siswa
	students := []Student{
		{"Siswa A", 85.5},
		{"Siswa B", 92.0},
		{"Siswa C", 78.5},
		{"Siswa D", 88.0},
		{"Siswa E", 95.5},
	}

	// Menghitung skor rata-rata
	var totalScore float64
	minScore := math.MaxFloat64
	maxScore := -math.MaxFloat64

	var studentWithMinScore Student
	var studentWithMaxScore Student

	for _, student := range students {
		totalScore += student.Score

		if student.Score < minScore {
			minScore = student.Score
			studentWithMinScore = student
		}

		if student.Score > maxScore {
			maxScore = student.Score
			studentWithMaxScore = student
		}
	}

	averageScore := totalScore / float64(len(students))

	// Menampilkan hasil
	fmt.Printf("Skor rata-rata: %.2f\n", averageScore)
	fmt.Printf("Siswa dengan skor minimum: %s dengan skor %.2f\n", studentWithMinScore.Name, studentWithMinScore.Score)
	fmt.Printf("Siswa dengan skor maksimum: %s dengan skor %.2f\n", studentWithMaxScore.Name, studentWithMaxScore.Score)
}
