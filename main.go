package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Tasks struct{
	ID string `json:"id"`
	TaskName string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date string `json:"date"`
}

var tasks []Tasks

func allTasks(){
	task:=Tasks{
		ID:"1",
		TaskName: "New projects",
		TaskDetail: "You must lead the project and finish it",
		Date: "2020-01-01",
	}
	tasks=append(tasks, task)
	task1:=Tasks{
		ID:"2",
		TaskName: "projects",
		TaskDetail: "lead the project and finish it",
		Date: "2020-01-09",
	}
	tasks=append(tasks, task1)
	fmt.Println("Your tasks are ",tasks)
}

func homePage(w http.ResponseWriter, r *http.Request ){
	fmt.Println("Welcome to the HomePage!")
}
func getTasks(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(tasks)
}
func getTask(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type","application/json")
	taskId:= mux.Vars(r);
	flag:=false
	for i := 0; i < len(tasks); i++ {
		if taskId["id"]==tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag=true
			break
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status":"Error"})
	}
}
func createTask(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type", "application/json")

    var task Tasks
    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    task.ID = strconv.Itoa(len(tasks) + 1)
    tasks = append(tasks, task)

    err = json.NewEncoder(w).Encode(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
func deleteTask(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    taskID := params["id"]
    for index, task := range tasks {
        if task.ID == taskID {
            tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Println(tasks)
            break
        }
    }
    json.NewEncoder(w).Encode(tasks)
}
func updateTask(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    taskID := params["id"]
    var updatedTask Tasks
    _ = json.NewDecoder(r.Body).Decode(&updatedTask)
    for index, task := range tasks {
        if task.ID == taskID {
            tasks[index] = updatedTask
            break
        }
    }
    json.NewEncoder(w).Encode(tasks)
}
func handleRouts(){
	router:= mux.NewRouter()
	router.HandleFunc("/",homePage).Methods("GET")
	router.HandleFunc("/gettasks",getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}",getTask).Methods("GET")
	router.HandleFunc("/createTask",createTask).Methods("POST")
	router.HandleFunc("/delete/{id}",deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}",updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8082",router))
}

func main(){
	allTasks()
	fmt.Println("hello")
	handleRouts()
}