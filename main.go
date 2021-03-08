package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

type Employee struct {
	/* Creating Employee Variables */
	EmpID       string `json:"EmpID"`
	EmpName     string `json:"EmpName"`
	EmpLocation string `json:"EmpLocation"`
}

var Employees []Employee

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Page!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/employees", returnAllEmployees)
	myRouter.HandleFunc("/employee/{id}", returnSingleEmployee)
	myRouter.HandleFunc("/employee", createNewEmployee).Methods("POST")
	myRouter.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")
	myRouter.HandleFunc("/employee/{id}", updateEmployee).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func returnAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllEmployees")
	json.NewEncoder(w).Encode(Employees)
}

func returnSingleEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, employee := range Employees {
		if employee.EmpID == id {
			json.NewEncoder(w).Encode(employee)
		}
	}
}

func createNewEmployee(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee Employee
	json.Unmarshal(reqBody, &employee)
	Employees = append(Employees, employee)
	json.NewEncoder(w).Encode(employee)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, employee := range Employees {
		if employee.EmpID == id {
			Employees = append(Employees[:index], Employees[index+1:]...)
		}
	}
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedEmployee Employee
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEmployee)
	for _, employee := range Employees {
		if employee.EmpID == id {
			employee.EmpID = updatedEmployee.EmpID
			employee.EmpName = updatedEmployee.EmpName
			employee.EmpLocation = updatedEmployee.EmpLocation
			Employees = append(Employees, employee)
			json.NewEncoder(w).Encode(employee)
		}

	}

}
func main() {
	Employees = []Employee{
		{EmpID: "10001", EmpName: "John", EmpLocation: "Bangalore"},
		{EmpID: "10002", EmpName: "Martin", EmpLocation: "Trivandrum"},
		{EmpID: "10003", EmpName: "Kate", EmpLocation: "Hyderabad"},
	}
	handleRequests()
}
