package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Student struct {
	Name string
	RollNo int
	marks map[string]int 
}

func promptOpt(st Student){
	reader := bufio.NewReader(os.Stdin)
	option, err := getUserInput("Select (A - Add marks, B - Save MarkSheet) : ", reader)

	if err != nil {
		panic(err)
	}

	switch option {
		case "A" :
			subjectName, _ := getUserInput("Enter Subject Name : ", reader)
			score, _ := getUserInput("Enter marks obtained : ", reader)

			s, err := strconv.Atoi(score)
			
			if err != nil{
				fmt.Println("Score must be a number")
				promptOpt(st)
			}

			st.addMarks(subjectName, s)
			promptOpt(st)
		
		case "B" :
			st.save()
		
		default :
			fmt.Println("Invalid Choice..")
			promptOpt(st)
	}
}
func (st *Student) addMarks(subName string, score int){
	(*st).marks[subName] = score
}
func (st *Student) format() string{
	result := "Marksheet \n"
	totalMarks := 0
	noofSubject := 0
	var percent float64 

	for key, val := range (*st).marks{
		result += fmt.Sprintf("%-25v : %-25v \n", key, val)
		noofSubject += 1
		totalMarks += val
	}

	percent = float64(totalMarks)/float64(noofSubject)
	result += "-----------------------------------------\n"
	result += fmt.Sprintf("Total %v : %v \n", totalMarks, noofSubject*100)
	result += fmt.Sprintf("Percentage : %0.2f \n", percent)

	return result

}
func (st *Student) save() {
	data := []byte(st.format())
	err := os.WriteFile("marksheet/"+ (*st).Name + "_marksheet" +".txt", data, 0644)
	if err != nil{
		panic(err)															
	}
	fmt.Println("Marksheet was saved in file")
}
func getUserInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func newStudent(name string, roll int) Student{
	fmt.Printf("Name: %v, RollNo : %v \n", name, roll)
	s := Student{
		Name : name,
		RollNo : roll,
		marks : map[string]int{},
	}
	return s
}
// func createReport() Student {
// 	reader := bufio.NewReader(os.Stdin)	
// 	name, _ := getUserInput("Enter name : ", reader)
// 	roll, _ := getUserInput("Enter roll no. : ", reader)
// 	r, err := strconv.Atoi(roll)
// 	if err != nil {
// 		panic(err)
// 	}	
// 	mystu := newStudent(name, r)
// 	return mystu
// }

func main() {
	fmt.Println("Structs in Golang")

	// student1 := createReport()
	// promptOpt(student1)

	students := map[string]int{
		"Amit": 19001,
		"Shikhar": 19002,
		"Yash": 19003,
	}
	for name, roll := range students{
		stu := newStudent(name, roll)
		promptOpt(stu)
	}
	
	
}