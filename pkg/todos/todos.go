package todos

import (
	"encoding/json"
	"fmt"
	"os"
)

const FILE_NAME = "todos.json"

type Todo struct {
	ID int `json:"id"`
	Description string `json:"description"`
	CreatedAt string `json:"createdAt"`
	IsComplete bool `json:"isComplete"`
}

func ListTodo(filename string) ([]Todo, error) {
	var todos []Todo

    fileData, err := os.ReadFile(filename)

    if err != nil {
        return todos, fmt.Errorf("failed to read file: %w", err)
    }

    err = json.Unmarshal(fileData, &todos) 

    if err != nil {
        return todos, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    return todos, nil
}

func AddTodo(todo Todo) ([]Todo, error) {
	todos,err := ListTodo(FILE_NAME);
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return todos, fmt.Errorf("failed to read file: %w", err)
	}
	todos = append(todos,todo)

	jsonData, err := json.MarshalIndent(todos, "", "  ")

	err = os.WriteFile(FILE_NAME, jsonData, 0644) 
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return todos , fmt.Errorf("Error writing to file: %w", err)
	}
	return todos,nil
}

func DeleteTodo(id int)(bool,error){
	todos,err := ListTodo(FILE_NAME);
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false, fmt.Errorf("failed to read file: %w", err)
	}
	var newTodos []Todo
	for _ , todo :=range todos{
		if todo.ID != id{
			newTodos = append(newTodos, todo)
		}
	}

	jsonData, err := json.MarshalIndent(newTodos, "", "  ")

	err = os.WriteFile(FILE_NAME, jsonData, 0644) 
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false , fmt.Errorf("Error writing to file: %w", err)
	}
	return true,nil
}

func UpdateTodo(id int,description string,is_complete bool)(bool,error){
	todos,err := ListTodo(FILE_NAME);
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false, fmt.Errorf("failed to read file: %w", err)
	}
	for index , todo :=range todos{
		if todo.ID == id{
			todos[index].Description = description
			todos[index].IsComplete = is_complete
		}
	}

	jsonData, err := json.MarshalIndent(todos, "", "  ")

	err = os.WriteFile(FILE_NAME, jsonData, 0644) 
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false , fmt.Errorf("Error writing to file: %w", err)
	}
	return true,nil
}