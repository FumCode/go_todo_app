package main

import (
	"fmt"
	"net/http"
)

var firstTask = "First Task"
var secondTask = "Second Task"
var therdTask = "Therd Task"
var taskList = []string{firstTask, secondTask, therdTask}

func main() {

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/tasks", showTasks)

	http.ListenAndServe(":8080", nil)

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskList {
		fmt.Fprintln(writer, task)
	}

}
