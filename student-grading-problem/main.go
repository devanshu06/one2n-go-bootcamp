package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

const (
	FirstName int = iota
	LastName
	University
	Test1Score
	Test2Score
	Test3Score
	Test4Score
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	_, err = csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	var students []student

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		test1, _ := strconv.Atoi(record[Test1Score])
		test2, _ := strconv.Atoi(record[Test2Score])
		test3, _ := strconv.Atoi(record[Test3Score])
		test4, _ := strconv.Atoi(record[Test4Score])

		students = append(students, student{
			firstName:  record[FirstName],
			lastName:   record[LastName],
			university: record[University],
			test1Score: test1,
			test2Score: test2,
			test3Score: test3,
			test4Score: test4,
		})
	}
	return students
}

func calculateGrade(students []student) []studentStat {
	var studentStats []studentStat

	for _, s := range students {
		finalScore := float32(s.test1Score+s.test2Score+s.test3Score+s.test4Score) / 4
		var grade Grade

		switch {
		case finalScore >= 70:
			grade = A
		case finalScore >= 50 && finalScore < 70:
			grade = B
		case finalScore >= 35 && finalScore < 50:
			grade = C
		default:
			grade = F
		}

		studentStats = append(studentStats, studentStat{
			student:    s,
			finalScore: finalScore,
			grade:      grade,
		})
	}
	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var topper studentStat

	for _, s := range gradedStudents {
		if s.finalScore > topper.finalScore {
			topper = s
		}
	}
	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	topperPerUniversity := make(map[string]studentStat)

	for _, s := range gs {
		if currentTopper, exists := topperPerUniversity[s.university]; !exists || s.finalScore > currentTopper.finalScore {
			topperPerUniversity[s.university] = s
		}
	}

	return topperPerUniversity
}

func main() {
	students := parseCSV("grades.csv")
	gradedStudents := calculateGrade(students)

	fmt.Println("Graded Students:")
	for _, s := range gradedStudents {
		fmt.Printf("%+v\n", s)
	}

	overallTopper := findOverallTopper(gradedStudents)
	fmt.Printf("\nOverall Topper: %+v\n", overallTopper)

	topperPerUniversity := findTopperPerUniversity(gradedStudents)
	fmt.Println("\nTopper Per University:")
	for university, topper := range topperPerUniversity {
		fmt.Printf("University: %s, Topper: %+v\n", university, topper)
	}
}
