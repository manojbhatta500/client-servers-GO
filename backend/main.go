package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var courselist []Course

func main() {
	c1 := Course{
		CourseId:   "1",
		CourseName: "software engineer",
		Price:      100,
		Auth: &Author{
			Fullname: "manoj bhatta",
			Website:  "manoj.co",
		},
	}
	courselist = append(courselist, c1)
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", addOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	http.ListenAndServe(":8000", r)
}

type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Auth       *Author `json:"auth"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func (c *Course) nilChecker() bool {
	if c.CourseId == "" && c.CourseName == "" {
		return true
	} else {
		return false
	}

}

// handler functions

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Write([]byte("<h1> hello welcome to the manojbhatta golang page <h1>"))

}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courselist)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println("the params in get one function is ", params)

	for _, v := range courselist {
		if v.CourseId == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("no course id found")
	return
}

func addOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add one course method executed")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("body is nil")
		return
	}
	var gotter Course
	json.NewDecoder(r.Body).Decode(&gotter)
	res := gotter.nilChecker()
	if res == true {
		json.NewEncoder(w).Encode("please send an actual data ")
		return
	}
	courselist = append(courselist, gotter)
	json.NewEncoder(w).Encode(gotter)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course controller")
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for i, v := range courselist {
		if v.CourseId == param["id"] {
			courselist = append(courselist[:i], courselist[i+1:]...)
		}
	}

	if r.Body == nil {
		json.NewEncoder(w).Encode("sorry body is empty")
		return
	}
	var getter Course
	json.NewDecoder(r.Body).Decode(&getter)
	if getter.nilChecker() == false {
		json.NewEncoder(w).Encode("sorry you need to send the id and course name ")
		return
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("delete one course function called")
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	for i, v := range courselist {
		if v.CourseId == id {
			courselist = append(courselist[:i], courselist[i+1:]...)
			json.NewEncoder(w).Encode("successfully deleted your course")
			return
		}
	}
	json.NewEncoder(w).Encode("sorry can't find your id ")
	return
}
