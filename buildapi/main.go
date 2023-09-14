package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for Courses - file
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author	`json:"author"`
}
type Author struct{
	FullName string `json:"fullname"`
	Website string 	`json:"website"`
}

// Creating Database (fake)
var courses []Course

// Middleware - file
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

// Controllers - file
// serve home 
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to building API</h1>"))
}


// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")   			// setting header
	json.NewEncoder(w).Encode(courses)								
}


// get one course with courseid from request
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// get id from request
	// id := r.URL.Query().Get("courseid")							// using url.Query
	// params := r.URL.Query()
	params := mux.Vars(r)											// using mux

	// loop through courses, find matching id
	for _ , course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		} 
	}
}


// create one course by generating courseid an data from req.body from request
func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
	} 

	var courseInstance Course
	_ = json.NewDecoder(r.Body).Decode(&courseInstance)
	
	// If course Instance is empty, No JSON Data
	if courseInstance.isEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	// If a duplicate course name is encountered
	for _, existingCourse := range courses {
        if existingCourse.CourseName == courseInstance.CourseName {
            json.NewEncoder(w).Encode("Duplicate courseName")
            return
        }
    }
	
	randomid, _ := rand.Prime(rand.Reader, 128)								// using crypto/rand package
	courseInstance.CourseId = randomid.String()								// generate unique id, string
	courses = append(courses, courseInstance)								// append new course into courses
	json.NewEncoder(w).Encode(courses)
	json.NewEncoder(w).Encode("Course successfully added..")
}


// update one course with courseid from request
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// get id from request
	params := mux.Vars(r)

	// loop through the courses, find id, remove, add with my id
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)						// remove from the slice
			var courseInstance Course													// created because we need to pass it as reference for decoding
			_ = json.NewDecoder(r.Body).Decode(&courseInstance)							// decode the req.body, updating the value directly
			courseInstance.CourseId = params["id"]										// rewriting the id
			courses = append(courses, courseInstance)									// appending the course with updated info to courses
			json.NewEncoder(w).Encode(courses)											// encoding to JSON
			json.NewEncoder(w).Encode("Successfully updated")
		}
	}
	json.NewEncoder(w).Encode("Could'nt get the id...")
}
// delete one course with courseid from request
func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")								

	params := mux.Vars(r)

	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]... )
			json.NewEncoder(w).Encode("Successfully deleted course..")
			break
		}
	}
}


func main() {
	fmt.Println("Building API in Golang")
	r := mux.NewRouter()

	// Seeding data
	courses = append(courses, 
		Course{
			CourseId: "1", 
			CourseName: "Reactjs", 
			CoursePrice: 999, 
			Author: &Author{
				FullName: "Abhijeet",
				Website: "abhijeet@go.dev",
			},
		},
		Course{
			CourseId: "2", 
			CourseName: "Nextjs", 
			CoursePrice: 999, 
			Author: &Author{
				FullName: "Abhi",
				Website: "abhi@go.dev",
			},
		},
		Course{
			CourseId: "3", 
			CourseName: "Nodejs", 
			CoursePrice: 1999, 
			Author: &Author{
				FullName: "Mishra",
				Website: "amishra@go.dev",
			},
		},
	)
	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}