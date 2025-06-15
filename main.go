package main

import "fmt"
import "net/http"

var todos = []string{"Go learn go lang", "go do schema changes"}

func printTodos(w http.ResponseWriter, r *http.Request) {

	for _, task := range todos {
		fmt.Fprintln(w, task)
	}
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	todo:= r.URL.Query().Get("task")
	todos = append(todos, todo)
}
func healthCheck(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Server is healthy")
}
func main() {
	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/todos", printTodos)
	http.HandleFunc("/addtodo", addTodo)

	http.ListenAndServe(":8080", nil)
}
